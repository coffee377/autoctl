package data

import (
	"cds/dingtalk/app"
	"regexp"
	"time"

	dingtalkworkflow10 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
)

// DingTalkWorkflowData 审批流程通用数据
type DingTalkWorkflowData struct {
	InstanceId  string     // 审批实例 ID
	BusinessId  string     // 审批编号
	CreateBy    string     // 创建人工号
	CreatorName string     // 创建人的名字（从 title 提取）
	CreateAt    *time.Time // 创建时间
	UpdateAt    *time.Time // 更新时间

	ApprovalStatus string // 审批状态

	res *dingtalkworkflow10.GetProcessInstanceResponseBodyResult

	*hooks
}

type KV struct {
	Code string // 编码
	Name string // 名称
}

type DictHook func(codeOrName string) (*KV, bool)

type hooks struct {
	creatorHook    app.UserHook // 获取创建人信息
	departmentHook DictHook     // 获取部门信息
}

type WorkflowOption func(opts *hooks)

func NewWorkflowData(instId string, res *dingtalkworkflow10.GetProcessInstanceResponseBodyResult, opts ...WorkflowOption) *DingTalkWorkflowData {
	data := &DingTalkWorkflowData{
		InstanceId: instId,
		res:        res,
		hooks:      &hooks{},
	}
	for _, opt := range opts {
		opt(data.hooks)
	}
	data.Extract(instId, res)
	return data
}

func (receiver *DingTalkWorkflowData) GetData() *dingtalkworkflow10.GetProcessInstanceResponseBodyResult {
	return receiver.res
}

// Extract 从审批实例响应中提取通用数据
func (receiver *DingTalkWorkflowData) Extract(instId string, res *dingtalkworkflow10.GetProcessInstanceResponseBodyResult) {
	receiver.InstanceId = instId
	receiver.BusinessId = *res.BusinessId
	// 审批发起人
	if receiver.creatorHook != nil {
		// 审批发起人钉钉 UserId
		u, ok := receiver.creatorHook(*res.OriginatorUserId)
		if ok {
			receiver.CreateBy = u.JobNumber
			receiver.CreatorName = u.RealName
		}
	} else {
		// 从标题中提取发起人姓名,获取不到，使用钉钉 UserId 兜底
		name, ok := receiver.extractCreatorName(*res.Title)
		if ok {
			receiver.CreateBy = name
			receiver.CreatorName = name
		} else {
			receiver.CreatorName = *res.OriginatorUserId
		}
	}

	receiver.CreateAt = receiver.convertTime(res.CreateTime)
	receiver.UpdateAt = receiver.convertTime(res.FinishTime)

	switch *res.Status {
	case "RUNNING":
		receiver.ApprovalStatus = "审批中"
		break
	case "TERMINATED":
		receiver.ApprovalStatus = "已撤销"
	case "COMPLETED":
		if *res.Result == "agree" {
			receiver.ApprovalStatus = "审批通过"
		} else if *res.Result == "refuse" {
			receiver.ApprovalStatus = "审批拒绝"
		}
	}
}

func (receiver *DingTalkWorkflowData) convertTime(timeStr *string) *time.Time {
	if timeStr == nil {
		return nil
	}
	// 2025-10-14T08:46Z
	t, _ := time.ParseInLocation("2006-01-02T15:04Z", *timeStr, time.Local)
	return &t
}

// extractPrefix 从字符串中提取"提交"前面的内容
func (receiver *DingTalkWorkflowData) extractCreatorName(s string) (string, bool) {
	// 正则表达式：
	// ^ 表示从字符串开头匹配
	// (.*?) 非贪婪匹配任意字符（除换行符），捕获到第一个分组
	// 提交 固定匹配"提交"关键词
	re := regexp.MustCompile(`^(.*?)提交`)

	// 查找匹配的子串
	match := re.FindStringSubmatch(s)
	if len(match) < 2 {
		// 没有匹配到（如字符串不含"提交"）
		return "", false
	}

	// 第一个分组（索引1）即为"提交"前面的内容
	return match[1], true
}

func WithUserHook(hook app.UserHook) WorkflowOption {
	return func(opts *hooks) {
		if hook != nil {
			opts.creatorHook = hook
		}
	}
}

func WithDepartmentHook(hook DictHook) WorkflowOption {
	return func(opts *hooks) {
		if hook != nil {
			opts.departmentHook = hook
		}
	}
}
