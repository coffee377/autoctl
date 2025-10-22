package data

import (
	"cds/bid/ds"
	"cds/bid/ent"
	"cds/bid/ent/bidapply"
	"cds/bid/ent/bidproject"
	"cds/dingtalk/oa"
	"context"
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

	Purchaser     *string // v2 采购人名称
	BidType       *string // v2 招标类型
	AgencyName    *string // v2 招标代理机构名称
	AgencyContact *string // v2 招标代理机构联系人及电话

	DepartmentName string     // v1 项目所属部门
	DepartmentCode *string    // v1 项目所属部门编码
	OpeningDate    *time.Time // v1 开标时间
	NoticeUrl      *string    // v1 招标网址
	BudgetAmount   float64    // v1 预算金额（元）
	Remark         *string    // v1 事项说明

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

	// 处理部门编码（根据名称进行转换）
	if apply.DepartmentCode == nil && apply.DepartmentName != "" && apply.hooks.departmentHook != nil {
		if kv, ok := apply.hooks.departmentHook(apply.DepartmentName); ok {
			apply.DepartmentCode = &kv.Code
		}
	}
	return apply, nil
}

func (af *BidApplyForm) Save(ctx context.Context, client *ent.Client) error {
	return ds.WithEntTx(ctx, client, func(tx *ent.Tx) error {
		// 项目信息
		project, err1 := af.saveProject(ctx, tx)
		if err1 != nil {
			return err1
		}
		// 申请信息
		_, err2 := af.saveApply(ctx, tx, project)
		if err2 != nil {
			return err2
		}
		return nil
	})
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

func (af *BidApplyForm) saveProject(ctx context.Context, tx *ent.Tx) (*ent.BidProject, error) {
	count, _ := tx.BidProject.Query().Where(bidproject.IDEQ(af.ProjectID)).Count(ctx)

	// 新增
	if count == 0 {
		projectCreate := tx.BidProject.Create()
		projectCreate.SetID(af.ProjectID)

		if af.ProjectCode != nil {
			projectCreate.SetCode(*af.ProjectCode)
		}
		projectCreate.SetName(af.ProjectName)

		if af.DepartmentCode != nil {
			projectCreate.SetDepartmentCode(*af.DepartmentCode)
		}
		projectCreate.SetDepartmentName(af.DepartmentName)

		projectCreate.SetBizRepNo(af.CreateBy)
		projectCreate.SetBizRepName(af.CreatorName)

		projectCreate.SetCreateAt(*af.CreateAt)
		if af.UpdateAt != nil {
			projectCreate.SetUpdateAt(*af.UpdateAt)
		}

		project, err := projectCreate.Save(ctx)
		if err != nil {
			return nil, err
		}
		return project, nil
	}

	// 更新
	projectUpdate := tx.BidProject.UpdateOneID(af.ProjectID)

	if af.ProjectCode != nil {
		projectUpdate.SetCode(*af.ProjectCode)
	}
	projectUpdate.SetName(af.ProjectName)

	if af.DepartmentCode != nil {
		projectUpdate.SetDepartmentCode(*af.DepartmentCode)
	}
	projectUpdate.SetDepartmentName(af.DepartmentName)

	projectUpdate.SetBizRepNo(af.CreateBy)
	projectUpdate.SetBizRepName(af.CreatorName)

	projectUpdate.SetCreateAt(*af.CreateAt)
	if af.UpdateAt != nil {
		projectUpdate.SetUpdateAt(*af.UpdateAt)
	}

	project, err := projectUpdate.Save(ctx)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (af *BidApplyForm) saveApply(ctx context.Context, tx *ent.Tx, project *ent.BidProject) (*ent.BidApply, error) {
	count, err := tx.BidApply.Query().Where(bidapply.ID(af.ID)).Count(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		create := tx.BidApply.Create()
		create.SetID(af.ID)
		create.SetBusinessID(af.BusinessId)
		create.SetInstanceID(af.InstanceId)
		create.SetProjectID(project.ID)

		// v2 投标申请新增字段
		create.SetNillablePurchaser(af.Purchaser)
		if af.BidType != nil {
			bidType := bidapply.BidType(*af.BidType)
			create.SetBidType(bidType)
		}
		create.SetNillableAgencyName(af.AgencyName)
		create.SetNillableAgencyContact(af.AgencyContact)

		// v1
		create.SetNillableOpeningDate(af.OpeningDate)
		create.SetNillableNoticeURL(af.NoticeUrl)
		create.SetBudgetAmount(af.BudgetAmount)
		create.SetNillableRemark(af.Remark)

		create.SetApprovalStatus(af.ApprovalStatus)

		create.SetCreateAt(*af.CreateAt)
		if af.UpdateAt != nil {
			create.SetUpdateAt(*af.UpdateAt)
		}

		apply, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return apply, nil
	}

	update := tx.BidApply.UpdateOneID(af.ID)
	update.SetBusinessID(af.BusinessId)
	update.SetInstanceID(af.InstanceId)
	update.SetProjectID(project.ID)

	// v2 投标新增字段
	update.SetNillablePurchaser(af.Purchaser)
	if af.BidType != nil {
		bidType := bidapply.BidType(*af.BidType)
		update.SetBidType(bidType)
	}
	update.SetNillableAgencyName(af.AgencyName)
	update.SetNillableAgencyContact(af.AgencyContact)

	// v1
	update.SetNillableOpeningDate(af.OpeningDate)
	update.SetNillableNoticeURL(af.NoticeUrl)
	update.SetBudgetAmount(af.BudgetAmount)
	update.SetNillableRemark(af.Remark)

	update.SetApprovalStatus(af.ApprovalStatus)

	update.SetCreateAt(*af.CreateAt)
	if af.UpdateAt != nil {
		update.SetUpdateAt(*af.UpdateAt)
	}

	apply, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return apply, nil
}
