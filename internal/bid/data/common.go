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
	CreatorName string     // 创建人的名字
	CreateAt    *time.Time // 创建时间
	UpdateAt    *time.Time // 更新时间

	ApprovalStatus string // 审批状态
	Done           bool   // 审批是否结束
	InvalidData    bool   // 无效数据

	res *dingtalkworkflow10.GetProcessInstanceResponseBodyResult

	*hooks
}

type KV struct {
	Code string // 编码
	Name string // 名称
}

type DictHook func(codeOrName string) (*KV, bool)

var (
	kvs = []*KV{
		{Code: "00100", Name: "综合管理部"},
		{Code: "00200", Name: "项目实施中心"},
		{Code: "00300", Name: "研发中心"},
		{Code: "00301", Name: "卫生产品部"},
		{Code: "00302", Name: "社会管理产品部"},
		{Code: "00303", Name: "社会救助产品部"},
		{Code: "00304", Name: "技术支持部"},
		{Code: "00305", Name: "医疗保险产品部"},
		{Code: "00400", Name: "财务部"},
		{Code: "00500", Name: "营销中心"},
		{Code: "00501", Name: "市场一部"},
		{Code: "00502", Name: "市场二部"},
		{Code: "00503", Name: "市场三部"},
		{Code: "00600", Name: "行政部"},
		{Code: "00700", Name: "商务部"},
		{Code: "00900", Name: "售后服务中心"},
		{Code: "01000", Name: "贵州分公司"},
		{Code: "01100", Name: "人口健康平台事业部"},
		{Code: "01200", Name: "基层卫生事业部"},
		{Code: "01300", Name: "民政综合平台事业部"},
		{Code: "01400", Name: "社会救助事业部"},
		{Code: "01500", Name: "技术研究中心"},
		{Code: "01600", Name: "培训中心"},
		{Code: "01700", Name: "系统集成事业部"},
		{Code: "01800", Name: "智慧健康养老事业部"},
		{Code: "01900", Name: "人力资源部"},
		{Code: "02100", Name: "民政事业部"},
		{Code: "02200", Name: "卫生事业部"},
		{Code: "02300", Name: "医疗保障事业部"},
		{Code: "02400", Name: "智慧医疗子公司"},
		{Code: "02500", Name: "总裁办"},
		{Code: "02600", Name: "证券部"},
		{Code: "02700", Name: "知识产权部"},
		{Code: "02800", Name: "贵州梵晶公司"},
		{Code: "02900", Name: "成都分公司"},
		{Code: "03000", Name: "亨源合义"},
		{Code: "03100", Name: "北京晶奇和一子公司"},
		{Code: "03200", Name: "成都晶高孙公司"},
		{Code: "03300", Name: "晶讯健康子公司"},
		{Code: "03400", Name: "哈尔滨分公司"},
		{Code: "03500", Name: "青海分公司"},
		{Code: "03600", Name: "创新管理委员会"},
		{Code: "03601", Name: "委员会办公室"},
		{Code: "03602", Name: "创新业务发展部"},
		{Code: "03603", Name: "创新项目组"},
		{Code: "03700", Name: "西藏晶喜子公司"},
		{Code: "03800", Name: "四川晶奇泽盛子公司"},
		{Code: "03900", Name: "智慧医疗事业部"},
		{Code: "04000", Name: "智慧医院事业部"},
		{Code: "04100", Name: "数字政企事业部"},
		{Code: "04200", Name: "数字能源事业部"},
		{Code: "04300", Name: "安全及技术保障部"},
	}
	DepartmentHook DictHook = func(key string) (*KV, bool) {
		for i, kv := range kvs {
			if kv.Code == key || kv.Name == key {
				return kvs[i], true
			}
		}
		return nil, false
	}
)

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
	if res == nil {
		return
	}
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
		receiver.Done = true
		receiver.ApprovalStatus = "已撤销"
		receiver.InvalidData = true
	case "COMPLETED":
		receiver.Done = true
		if *res.Result == "agree" {
			receiver.ApprovalStatus = "审批通过"
		} else if *res.Result == "refuse" {
			receiver.ApprovalStatus = "审批拒绝"
			receiver.InvalidData = true
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

// 正则表达式,匹配任意类似“xxx(yyy)”的字符串，`\(([^)]+)\)`
var codeReg = regexp.MustCompile(`\(([^)]+)\)`)

func (receiver *DingTalkWorkflowData) ExtraDictCode(text *string) (string, bool) {
	if text == nil {
		return "", false
	}
	match := codeReg.FindStringSubmatch(*text)
	if len(match) < 2 {
		return *text, false
	}
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
