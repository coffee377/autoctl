package data

import (
	"cds/dingtalk/oa"
	"crypto/md5"
	"encoding/hex"
	"time"

	dingtalkworkflow10 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
)

type BidApplyForm struct {
	*DingTalkWorkflowData

	ID        string // 申请 ID
	ProjectID string // 项目ID

	ProjectName string  // v1 项目名称
	ProjectCode *string // v2 项目编号
	ProjectType *string // v2 项目类型

	DepartmentName string     // v1 项目所属部门
	DepartmentCode *string    // v1 项目所属部门编码
	OpeningDate    *time.Time // v1 开标时间
	NoticeUrl      *string    // v1 招标网址
	BudgetAmount   float64    // v1 预算金额（元）
	Remark         *string    // v2 事项说明

	Handler    string // v1 办理人
	Attachment string // v1 附件

}

func NewBidApply(instId string, res *dingtalkworkflow10.GetProcessInstanceResponseBodyResult, opts ...WorkflowOption) (*BidApplyForm, error) {
	apply := &BidApplyForm{
		DingTalkWorkflowData: NewWorkflowData(instId, res, opts...),
	}
	// 生成项目 ID 和申请 ID
	apply.generateID()
	// 调用通用映射函数填充表单数据
	mappers := apply.getApplyMappers()
	err := oa.MapFormToEntity(res, mappers, apply)
	if err != nil {
		return nil, err
	}
	return apply, nil
}

// GenerateID 根据ApprovalNumber生成唯一Id
func (af *BidApplyForm) generateID() {
	// 1. 组合两个字段作为哈希源（使用特殊分隔符避免字段值拼接冲突）
	wd := af.DingTalkWorkflowData
	p := wd.BusinessId + "|project"
	a := wd.BusinessId + "|apply"
	//b := af.BusinessId + "|bid"

	// 2. 计算SHA-1哈希（160位，20字节）
	pHash := md5.Sum([]byte(p))
	aHash := md5.Sum([]byte(a))
	//bHash := md5.Sum([]byte(b))

	// 3. 转换为16进制字符串（32个字符，因为MD5是16字节=32 hex字符）
	af.ProjectID = hex.EncodeToString(pHash[:])
	af.ID = hex.EncodeToString(aHash[:])
	//af.BidId = hex.EncodeToString(bHash[:])
}

func (af *BidApplyForm) getApplyMappers() []oa.FieldMapper {
	return []oa.FieldMapper{
		{
			ComponentId: "TextField_1FNYLBKS38XS0", // 项目名称组件ID
			FieldName:   "ProjectName",
			Converter:   oa.StringConverter, // 字符串转换
		},
		{
			ComponentId: "DDSelectField_5CJS7PFW1CG0", // 所属部门组件ID
			FieldName:   "DepartmentName",
			Converter:   oa.StringConverter,
		},
		{
			ComponentId: "TextField_1P4X7NQK70W00", // 招标网址组件ID
			FieldName:   "NoticeUrl",
			Converter:   oa.StringConverter,
			Pointer:     true,
		},
		{
			ComponentId: "DDDateField_1GR3BL6HLUWW", // 开标时间组件ID
			FieldName:   "OpeningDate",
			// 时间转换（格式"2006-01-02"，本地时区）
			Converter: oa.DateConverter(time.DateOnly, time.Local),
			Pointer:   true,
		},
		{
			ComponentId: "MoneyField_7FQ2FMK1KQC0", // 预算金额组件ID
			FieldName:   "BudgetAmount",
			Converter:   oa.Float64Converter, // 浮点数转换
		},
		{
			ComponentId: "TextareaField_1EP1SOW22D1C0", // 事项说明组件ID
			FieldName:   "Remark",
			Converter:   oa.StringConverter,
			Pointer:     true,
		},
		{
			ComponentId: "InnerContactField_1ER3G5MU4HR40", // 办理人组件ID
			FieldName:   "Handler",
			Converter:   oa.StringConverter,
		},
		{
			ComponentId: "DDAttachment_I8PPSWWCCDC0", // 附件组件ID
			FieldName:   "Attachment",
			Converter:   oa.StringConverter,
		},
	}
}
