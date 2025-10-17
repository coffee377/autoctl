package oa

import (
	"testing"
	"time"

	"github.com/coffee377/autoctl/internal/dingtalk/app"
	"github.com/coffee377/autoctl/pkg/log"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

const (
	BidApplyProcessCode   = "PROC-958C3100-85BF-45D3-8583-6645DA922756" // 投标申请审批表单编码
	BidExpenseProcessCode = "PROC-D8453B77-B313-4BEB-BE42-C71EE81DA61A" // 投标项目转款表单编码
)

var (
	approval *Approval
	err      error
)

func init() {
	a := app.New("a57e9681-79cb-4242-96df-952be2dc3af7",
		app.WithRedis(redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "redis!@@&",
			DB:       0,
		}),
	)
	approval, err = New(a)
	if err != nil {
		log.Error(err.Error())
	}
}

func TestGetProcessInstanceIds(t *testing.T) {
	ids, err2 := approval.GetProcessInstanceIdsByMonth(BidApplyProcessCode, 2025, 1, nil)
	assert.Nil(t, err2)
	assert.NotNil(t, ids)
}

func TestGetProcessInstance(t *testing.T) {
	instId := "Rp9D_t0WQrqgpfxvxUZ_EQ07201760318742"
	res, err := approval.GetProcessInstance(instId)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	status := res.Status
	assert.Equal(t, "COMPLETED", *status)
	result := res.Result
	assert.Equal(t, "agree", *result)

	var apply Apply
	apply.ID = *res.BusinessId
	apply.ApprovalNumber = *res.BusinessId
	apply.InstanceId = instId

	// 调用通用映射函数填充表单数据
	mappers := getApplyMappers()
	err = MapFormToEntity(res, mappers, &apply)
	assert.Nil(t, err, "表单映射失败")

	t.Logf("映射后的Apply: %+v", apply)

}

type Apply struct {
	ID             string
	ApprovalNumber string
	ProjectName    string     // 项目名称
	DepartmentName string     // 项目所属部门
	NoticeUrl      string     // 招标网址
	OpeningDate    *time.Time // 开标时间
	BudgetAmount   float64    // 预算金额（元）
	Remark         string     // 事项说明
	Handler        string     // 办理人
	Attachment     string     // 附件
	InstanceId     string
}

// 定义Apply的表单映射规则
func getApplyMappers() []FieldMapper {
	return []FieldMapper{
		{
			ComponentId: "TextField_1FNYLBKS38XS0", // 项目名称组件ID
			FieldName:   "ProjectName",
			Converter:   StringConverter, // 字符串转换
		},
		{
			ComponentId: "DDSelectField_5CJS7PFW1CG0", // 所属部门组件ID
			FieldName:   "DepartmentName",
			Converter:   StringConverter,
		},
		{
			ComponentId: "TextField_1P4X7NQK70W00", // 招标网址组件ID
			FieldName:   "NoticeUrl",
			Converter:   StringConverter,
		},
		{
			ComponentId: "DDDateField_1GR3BL6HLUWW", // 开标时间组件ID
			FieldName:   "OpeningDate",
			// 时间转换（格式"2006-01-02"，本地时区）
			Converter: DateConverter(time.DateOnly, time.Local),
		},
		{
			ComponentId: "MoneyField_7FQ2FMK1KQC0", // 预算金额组件ID
			FieldName:   "BudgetAmount",
			Converter:   Float64Converter, // 浮点数转换
		},
		{
			ComponentId: "TextareaField_1EP1SOW22D1C0", // 事项说明组件ID
			FieldName:   "Remark",
			Converter:   StringConverter,
		},
		{
			ComponentId: "InnerContactField_1ER3G5MU4HR40", // 办理人组件ID
			FieldName:   "Handler",
			Converter:   StringConverter,
		},
		{
			ComponentId: "DDAttachment_I8PPSWWCCDC0", // 附件组件ID
			FieldName:   "Attachment",
			Converter:   StringConverter,
		},
	}
}
