package es

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// MillisecondTime 自定义时间类型，用于解析毫秒级时间戳为time.Time
type MillisecondTime time.Time

func (t *MillisecondTime) MarshalJSON() ([]byte, error) {
	// 处理 nil 情况（若变量未初始化）
	if t == nil {
		return json.Marshal(nil)
	}
	ms := time.Time(*t).UnixMilli() // 获取毫秒数
	return json.Marshal(ms)
}

// UnmarshalJSON 实现json.Unmarshaler接口，解析毫秒时间戳
func (t *MillisecondTime) UnmarshalJSON(data []byte) error {
	// 1. 处理空值（JSON 中的 null）
	if string(data) == "null" {
		*t = MillisecondTime(time.Time{}) // 设为零值时间
		return nil
	}

	// 2. 先尝试解析为数字型时间戳
	var msNum int64
	if err := json.Unmarshal(data, &msNum); err == nil {
		return t.setFromMillisecond(msNum)
	}

	// 3. 解析失败则尝试字符串型时间戳
	var msStr string
	if err := json.Unmarshal(data, &msStr); err != nil {
		return fmt.Errorf("时间戳解析失败：既非数字也非字符串，原始数据：%s，错误：%w", string(data), err)
	}

	// 4. 字符串转 int64
	msNum, err := strconv.ParseInt(msStr, 10, 64)
	if err != nil {
		return fmt.Errorf("字符串时间戳转数字失败：%s，错误：%w", msStr, err)
	}

	// 5. 从毫秒数设置时间
	return t.setFromMillisecond(msNum)
}

// setFromMillisecond 从毫秒数设置时间，并校验有效性
func (t *MillisecondTime) setFromMillisecond(ms int64) error {
	// 校验毫秒数的合理性（可选：根据业务调整范围，比如不早于1970年，不晚于未来10年）
	minMs := int64(0) // 1970-01-01 00:00:00 UTC
	maxMs := time.Now().AddDate(10, 0, 0).UnixMilli()
	if ms < minMs || ms > maxMs {
		return fmt.Errorf("毫秒时间戳超出合理范围：%d（有效范围：%d~%d）", ms, minMs, maxMs)
	}

	// 转换为 time.Time 并赋值给自定义类型
	*t = MillisecondTime(time.UnixMilli(ms))
	return nil
}

// 可选：实现String()方法，方便打印时格式化
func (t *MillisecondTime) String() string {
	if t == nil {
		return ""
	}
	t2 := time.Time(*t)
	if t2.IsZero() {
		return ""
	}
	return t2.Local().Format(time.DateTime)
}

type ProcessBase struct {
	ProcessCode string          `json:"processCode"`       // 审批模板的唯一码
	InstanceId  string          `json:"processInstanceId"` // 审批实例 id
	BusinessId  string          `json:"businessId"`        // 流程实例业务标识
	CreateTime  MillisecondTime `json:"createTime"`        // 创建(审批实例/任务)时间。时间戳，单位毫秒
	FinishTime  MillisecondTime `json:"finishTime"`        // 结束(审批实例/任务)时间。时间戳，单位毫秒
	Type        string          `json:"type"`              // 实例状态变更类型或任务变更状态类型
}

// InstanceMessage https://open.dingtalk.com/document/development/event-bpms-instance-change
type InstanceMessage struct {
	InstanceId    string          `json:"processInstanceId"` // 审批实例 id
	FinishTime    MillisecondTime `json:"finishTime"`        // 结束审批实例时间。时间戳，单位毫秒
	CreateTime    MillisecondTime `json:"createTime"`        // 创建审批实例时间。时间戳，单位毫秒
	ProcessCode   string          `json:"processCode"`       // 审批模板的唯一码
	BizCategoryId string          `json:"bizCategoryId"`     // 业务分类标识
	BusinessId    string          `json:"businessId"`        // 流程实例业务标识
	Type          string          `json:"type"`              // 实例状态变更类型：start：审批实例开始 finish：审批正常结束（同意或拒绝） terminate：审批终止（发起人撤销审批单） delete：审批实例删除
	Title         string          `json:"title"`             // 审批实例标题
	BusinessType  string          `json:"businessType"`      // 业务身份
	Url           string          `json:"url"`               // 审批实例详情页 URL
	StaffId       string          `json:"staffId"`           // 发起审批实例的员工 userId
	Result        string          `json:"result"`            // 审批结果(审批终止时无此参数) agree： 同意  refuse：拒绝
}

func (i *InstanceMessage) UnmarshalJSON(bytes []byte) error {
	// 定义临时结构体（避免直接解析时递归调用UnmarshalJSON）
	type Temp InstanceMessage
	var temp Temp

	// 先使用标准库解析 JSON 到临时结构体
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("解析JSON失败: %w", err)
	}

	// 将临时结构体的值赋给原结构体
	*i = InstanceMessage(temp)
	return nil
}

// TaskMessage https://open.dingtalk.com/document/development/event-bpms-task-change
type TaskMessage struct {
	ProcessInstanceId string          `json:"processInstanceId"` // 审批实例 id
	FinishTime        MillisecondTime `json:"finishTime"`        // 结束任务的时间。时间戳，单位毫秒
	CreateTime        MillisecondTime `json:"createTime"`        // 创建任务的时间。时间戳，单位毫秒
	ProcessCode       string          `json:"processCode"`       // 审批模板的唯一码
	BizCategoryId     string          `json:"bizCategoryId"`     // 业务分类标识
	BusinessId        string          `json:"businessId"`        // 流程实例业务标识
	Remark            string          `json:"remark"`            // 操作时写的评论内容
	Type              string          `json:"type"`              // 任务类型：start：审批任务开始 finish：审批任务正常结束（完成或转交）,cancel：说明当前节点有多个审批人并且是或签，其中一个人执行了审批，其他审批人会推送cancel类型事件 comment：审批任务评论。
	Title             string          `json:"title"`             // 审批实例标题
	TaskId            int64           `json:"taskId"`            // 任务 id
	StaffId           string          `json:"staffId"`           // 用户 userId
	Result            string          `json:"result"`            // 审批结果 agree,refuse,redirect,audit
}

func (t *TaskMessage) UnmarshalJSON(bytes []byte) error {
	// 定义临时结构体（避免直接解析时递归调用UnmarshalJSON）
	type Temp TaskMessage
	var temp Temp

	// 先使用标准库解析 JSON 到临时结构体
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("解析JSON失败: %w", err)
	}

	// 将临时结构体的值赋给原结构体
	*t = TaskMessage(temp)
	return nil
}
