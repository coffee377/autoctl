package process

import (
	"encoding/json"
	"fmt"
	"time"
)

// MillisecondTime 自定义时间类型，用于解析毫秒级时间戳为time.Time
type MillisecondTime time.Time

func (t *MillisecondTime) MarshalJSON() ([]byte, error) {
	ms := time.Time(*t).UnixMilli() // 获取毫秒数
	return json.Marshal(ms)
}

// UnmarshalJSON 实现json.Unmarshaler接口，解析毫秒时间戳
func (t *MillisecondTime) UnmarshalJSON(data []byte) error {
	// 1. 先将JSON数据解析为int64（毫秒数）
	var ms int64
	if err := json.Unmarshal(data, &ms); err != nil {
		return fmt.Errorf("时间戳解析失败: %w", err)
	}

	*t = MillisecondTime(time.UnixMilli(ms))
	return nil
}

// 可选：实现String()方法，方便打印时格式化
func (t MillisecondTime) String() string {
	return time.Time(t).Local().Format(time.DateTime)
}

// InstanceMessage https://open.dingtalk.com/document/development/event-bpms-instance-change
type InstanceMessage struct {
	InstanceId    string          `json:"processInstanceId"` // 审批实例id
	FinishTime    MillisecondTime `json:"finishTime"`        // 结束审批实例时间。时间戳，单位毫秒
	CreateTime    MillisecondTime `json:"createTime"`        // 创建审批实例时间。时间戳，单位毫秒
	ProcessCode   string          `json:"processCode"`       // 审批模板的唯一码
	BizCategoryId string          `json:"bizCategoryId"`     // 业务分类标识
	BusinessId    string          `json:"businessId"`        // 流程实例业务标识
	Type          string          `json:"type"`              // 实例状态变更类型：start：审批实例开始 finish：审批正常结束（同意或拒绝） terminate：审批终止（发起人撤销审批单） delete：审批实例删除
	Title         string          `json:"title"`             // 审批实例标题
	BusinessType  string          `json:"businessType"`      // 业务身份
	Url           string          `json:"url"`               // 审批实例详情页URL
	StaffId       string          `json:"staffId"`           // 发起审批实例的员工 userId
	Result        string          `json:"result"`            // 审批结果(审批终止时无此参数) agree： 同意  refuse：拒绝
}

func (i *InstanceMessage) UnmarshalJSON(bytes []byte) error {
	// 定义临时结构体（避免直接解析时递归调用UnmarshalJSON）
	type Temp InstanceMessage
	var temp Temp

	// 先使用标准库解析JSON到临时结构体
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("解析JSON失败: %w", err)
	}

	// 将临时结构体的值赋给原结构体
	*i = InstanceMessage(temp)
	return nil
}

// TaskMessage https://open.dingtalk.com/document/development/event-bpms-task-change
type TaskMessage struct {
	ProcessInstanceId string          `json:"processInstanceId"` // 审批实例id
	FinishTime        MillisecondTime `json:"finishTime"`        // 结束任务的时间。时间戳，单位毫秒
	CreateTime        MillisecondTime `json:"createTime"`        // 创建任务的时间。时间戳，单位毫秒
	ProcessCode       string          `json:"processCode"`       // 审批模板的唯一码
	BizCategoryId     string          `json:"bizCategoryId"`     // 业务分类标识
	BusinessId        string          `json:"businessId"`        // 流程实例业务标识
	Remark            string          `json:"remark"`            // 操作时写的评论内容
	Type              string          `json:"type"`              // 任务类型：start：审批任务开始 finish：审批任务正常结束（完成或转交）,cancel：说明当前节点有多个审批人并且是或签，其中一个人执行了审批，其他审批人会推送cancel类型事件 comment：审批任务评论。
	Title             string          `json:"title"`             // 审批实例标题
	TaskId            string          `json:"taskId"`            // 任务id
	StaffId           string          `json:"staffId"`           // 用户userId
	Result            string          `json:"result"`            // 审批结果 agree,refuse,redirect,audit
}

func (t *TaskMessage) UnmarshalJSON(bytes []byte) error {
	// 定义临时结构体（避免直接解析时递归调用UnmarshalJSON）
	type Temp TaskMessage
	var temp Temp

	// 先使用标准库解析JSON到临时结构体
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("解析JSON失败: %w", err)
	}

	// 将临时结构体的值赋给原结构体
	*t = TaskMessage(temp)
	return nil
}
