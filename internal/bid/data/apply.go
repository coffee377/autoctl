package data

import (
	"cds/bid/ds"
	"cds/bid/ent"
	"cds/bid/ent/bidapply"
	"cds/bid/ent/bidproject"
	"cds/bid/ent/schema"
	"cds/dingtalk/oa"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strings"
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

	PurchaserName            *string // v2 采购人名称
	BidType                  *string // v2 招标类型(采购方式)
	AgencyName               *string // v2 招标代理机构名称
	AgencyContactPerson      *string // v2 招标代理机构联系
	AgencyContactInformation *string // v2 招标代理机构联系方式

	DepartmentName string     // v1 项目所属部门
	DepartmentCode *string    // v1 项目所属部门编码
	OpeningDate    *time.Time // v1 开标时间
	NoticeUrl      *string    // v1 招标网址
	BudgetAmount   float64    // v1 预算金额（元）
	Remark         *string    // v1 备注说明

	BidRegistrationHandler string // v1 投标报名办理人
	Attachments            string // v1 投标报名附件

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

	// 处理项目类型编码
	if code, ok := apply.ExtraDictCode(apply.ProjectType); ok {
		if code == "I" {
			code = "SH"
		}
		apply.ProjectType = &code
	}

	// 处理招标类型编码
	if code, ok := apply.ExtraDictCode(apply.BidType); ok {
		apply.BidType = &code
	}

	return apply, nil
}

func (af *BidApplyForm) Save(ctx context.Context, client *ent.Client, projectUpdatable bool) error {
	return ds.WithEntTx(ctx, client, func(tx *ent.Tx) error {
		count, _ := tx.BidApply.Query().Where(bidapply.ID(af.ID)).Count(ctx)
		projectDone := false
		if count > 0 {
			projectDone = true
		}
		// 项目信息
		_, err1 := af.saveProject(ctx, tx, projectDone, projectUpdatable)
		if err1 != nil {
			return err1
		}

		// 申请信息
		_, err2 := af.saveApply(ctx, tx)
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
		{ComponentId: "TextField_1FNYLBKS38XS0", FieldName: "ProjectName", Converter: oa.StringConverter},
		{ComponentId: "DDSelectField_12X96ZZZPIIO0", FieldName: "ProjectType", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "TextField_R6ZC4JEPWHZ4", FieldName: "ProjectCode", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "DDSelectField_5CJS7PFW1CG0", FieldName: "DepartmentName", Converter: oa.StringConverter},
		{ComponentId: "TextField_1WA1WVNMXKV7K", FieldName: "PurchaserName", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "DDSelectField_DYE2Z0A8VY80", FieldName: "BidType", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "TextField_LJX2NLWH06WW", FieldName: "AgencyName", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "TextField_JJY6USO8AA68", FieldName: "AgencyContactPerson", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "TextField_F62UJ59YNTHC", FieldName: "AgencyContactInformation", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "MoneyField_7FQ2FMK1KQC0", FieldName: "BudgetAmount", Converter: oa.Float64Converter},
		{ComponentId: "DDDateField_1GR3BL6HLUWW", FieldName: "OpeningDate", Converter: oa.DateConverters([]string{"2006-01-02 15:04", time.DateOnly}, time.Local), Pointer: true},
		{ComponentId: "TextField_1P4X7NQK70W00", FieldName: "NoticeUrl", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "TextareaField_1EP1SOW22D1C0", FieldName: "Remark", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "InnerContactField_1ER3G5MU4HR40", FieldName: "BidRegistrationHandler", Converter: oa.StringConverter},
		{ComponentId: "DDAttachment_I8PPSWWCCDC0", FieldName: "Attachments", Converter: oa.StringConverter},
	}
}

func (af *BidApplyForm) saveProject(ctx context.Context, tx *ent.Tx, done, updatable bool) (*ent.BidProject, error) {
	res, err := tx.BidProject.Query().Where(bidproject.IDEQ(af.ProjectID)).Only(ctx)
	if ent.IsNotFound(err) {
		return af.createProject(ctx, tx)
	} else if err != nil {
		return nil, err
	}
	if done {
		return res, nil
	}
	if updatable {
		return af.updateProject(ctx, tx)
	}
	return res, nil
}

func (af *BidApplyForm) createProject(ctx context.Context, tx *ent.Tx) (*ent.BidProject, error) {
	create := tx.BidProject.Create()
	create.SetID(af.ProjectID)

	create.SetNillableCode(af.ProjectCode)
	create.SetName(af.ProjectName)
	if af.ProjectType != nil {
		create.SetType(bidproject.Type(*af.ProjectType))
	}
	create.SetSource(bidproject.SourceBA)

	create.SetNillableDepartmentCode(af.DepartmentCode)
	create.SetDepartmentName(af.DepartmentName)

	create.SetBizRepNo(af.CreateBy)
	create.SetBizRepName(af.CreatorName)

	create.SetCreatedAt(*af.CreateAt)
	create.SetCreatedBy(af.CreateBy)
	if af.UpdateAt != nil {
		create.SetUpdatedAt(*af.UpdateAt)
	} else {
		create.SetUpdatedAt(*af.CreateAt)
	}

	project, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (af *BidApplyForm) updateProject(ctx context.Context, tx *ent.Tx) (*ent.BidProject, error) {
	update := tx.BidProject.UpdateOneID(af.ProjectID)

	update.SetNillableCode(af.ProjectCode)
	update.SetName(af.ProjectName)
	if af.ProjectType != nil {
		update.SetType(bidproject.Type(*af.ProjectType))
	}

	update.SetNillableDepartmentCode(af.DepartmentCode)
	update.SetDepartmentName(af.DepartmentName)

	update.SetBizRepNo(af.CreateBy)
	update.SetBizRepName(af.CreatorName)

	if af.UpdateAt != nil {
		update.SetUpdatedAt(*af.UpdateAt)
	}

	project, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (af *BidApplyForm) saveApply(ctx context.Context, tx *ent.Tx) (*ent.BidApply, error) {
	res, err := tx.BidApply.Query().Where(bidapply.ID(af.ID)).Only(ctx)
	if ent.IsNotFound(err) {
		return af.createApply(ctx, tx)
	} else if err != nil {
		return nil, err
	}

	if res.Done {
		return res, nil
	}

	return af.updateApply(ctx, tx)
}

func (af *BidApplyForm) createApply(ctx context.Context, tx *ent.Tx) (*ent.BidApply, error) {
	create := tx.BidApply.Create()

	create.SetID(af.ID)
	create.SetBusinessID(af.BusinessId)
	create.SetInstanceID(af.InstanceId)
	create.SetProjectID(af.ProjectID)

	// v2 投标申请新增字段
	create.SetNillablePurchaserName(af.PurchaserName)
	if af.BidType != nil {
		bidType := bidapply.BidType(*af.BidType)
		create.SetBidType(bidType)
	}
	create.SetNillableAgencyName(af.AgencyName)
	if af.AgencyContactPerson != nil && af.AgencyContactInformation != nil {
		builder := strings.Builder{}
		builder.WriteString(*af.AgencyContactPerson)
		builder.WriteString(" ")
		builder.WriteString(*af.AgencyContactInformation)
		if builder.String() != "" {
			create.SetAgencyContact(builder.String())
		}
	}

	// v1
	create.SetNillableOpeningDate(af.OpeningDate)
	create.SetNillableNoticeURL(af.NoticeUrl)
	create.SetBudgetAmount(af.BudgetAmount)
	create.SetNillableRemark(af.Remark)

	var attachments []schema.Attachment
	if af.Attachments != "" {
		_ = json.Unmarshal([]byte(af.Attachments), &attachments)
	}
	if attachments != nil {
		create.SetAttachments(attachments)
	}

	create.SetApprovalStatus(af.ApprovalStatus)
	create.SetDone(af.Done)
	create.SetCreatedAt(*af.CreateAt)
	create.SetCreatedBy(af.CreateBy)
	if af.UpdateAt != nil {
		create.SetUpdatedAt(*af.UpdateAt)
	} else {
		create.SetUpdatedAt(*af.CreateAt)
	}

	apply, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func (af *BidApplyForm) updateApply(ctx context.Context, tx *ent.Tx) (*ent.BidApply, error) {
	update := tx.BidApply.UpdateOneID(af.ID)

	update.SetApprovalStatus(af.ApprovalStatus)
	update.SetDone(af.Done)

	if af.UpdateAt != nil {
		update.SetUpdatedAt(*af.UpdateAt)
	}

	apply, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return apply, nil
}
