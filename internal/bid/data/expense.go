package data

import (
	"cds/bid/ds"
	"cds/bid/ent"
	"cds/bid/ent/bidexpense"
	"cds/dingtalk/oa"
	"context"
	"crypto/md5"
	"encoding/hex"
	"time"

	dingtalkworkflow10 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
)

type BidExpenseForm struct {
	*DingTalkWorkflowData

	ID     string // 支出 ID
	BillNo string // 单据编码

	ProjectName string  // 项目名称
	ProjectCode string  // 项目编号
	BizRepName  string  // 商务代表
	Purchaser   *string // 采购人名称

	FeeTypeV1 *string // 费用类型,原始字符串
	FeeType   string  // 费用类型 RF:报名费 DF:标书工本费 CA:CA费用 EF:专家费 BB:投标保证金 BS:中标服务费 PB:履约保证金 OE:其他费用
	PayReason *string // 付款事由, 其他费用类型时备注

	PayeeBank    string // 收款方开户银行
	PayeeName    string // 收款方账户名称
	PayeeAccount string // 收款方账号

	PayRatio    *float64   // 付款比例
	PayAmount   float64    // 付款金额（元）
	PayRemark   *string    // 付款备注
	PayMethod   *string    // 付款方式
	PlanPayTime *time.Time // 预计转账时间
}

func NewBidExpense(instId string, res *dingtalkworkflow10.GetProcessInstanceResponseBodyResult, opts ...WorkflowOption) (*BidExpenseForm, error) {
	expense := &BidExpenseForm{
		DingTalkWorkflowData: NewWorkflowData(instId, res, opts...),
	}
	// 生成支出 ID
	expense.generateID()
	// 调用通用映射函数填充表单数据
	mappers := expense.getApplyMappers()
	err := oa.MapFormToEntity(res, mappers, expense)
	if err != nil {
		return nil, err
	}

	// 费用类型转换, v2 版本支持
	code, ok := expense.ExtraDictCode(&expense.FeeType)
	if ok {
		expense.FeeTypeV1 = nil
		expense.FeeType = code
	} else {
		expense.FeeTypeV1 = &code
		expense.FeeType = ""
	}

	return expense, nil
}

func (ef *BidExpenseForm) Save(ctx context.Context, client *ent.Client) error {
	return ds.WithEntTx(ctx, client, func(tx *ent.Tx) error {
		// 申请信息
		_, err := ef.saveExpense(ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})
}

func (ef *BidExpenseForm) generateID() {
	// 1. 组合两个字段作为哈希源（使用特殊分隔符避免字段值拼接冲突）
	wd := ef.DingTalkWorkflowData
	e := wd.BusinessId

	// 2. 计算SHA-1哈希（160位，20字节）
	eHash := md5.Sum([]byte(e))

	// 3. 转换为16进制字符串（32个字符，因为MD5是16字节=32 hex字符）
	ef.ID = hex.EncodeToString(eHash[:])
}

// {ComponentId: "SeqNumberField_1U1O5KHHSHUDC", FieldName: "单据编号", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "TextField_22K742TU7DTS0", FieldName: "单据编号（废弃）", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "FormRelateField_7YS5XZOZRUDC", FieldName: "项目投标", Converter: oa.StringConverter},
// {ComponentId: "TextField_17HCU4VR5UQK0", FieldName: "项目名称", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "TextField_82D5H6HY60O0", FieldName: "项目编号", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "TextField_8L38ISBPCSG0", FieldName: "商务代表", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "TextField_3Y2MPH8USDG0", FieldName: "采购人", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "MoneyField_1SA588VBGU3K0", FieldName: "中标价/预算价（元）", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "DDSelectField_RE7H7DQVRPS0", FieldName: "费用类型", Converter: oa.StringConverter},
// {ComponentId: "MoneyField_QP9DYIENY4G0", FieldName: "付款金额", Converter: oa.StringConverter},
// {ComponentId: "TextField_1SRLTMRI2F4W0", FieldName: "付款事由", Converter: oa.StringConverter},
// {ComponentId: "TextField_1YIHQLKCJLUO0", FieldName: "付款备注", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "RecipientAccountField_1FXC118HK3MDC", FieldName: "收款账户", Converter: oa.StringConverter},
// {ComponentId: "TextField_1HHDR2YRQGPS0", FieldName: "开户银行（废弃）", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "TextField_138818FKKQC00", FieldName: "账户名称（废弃）", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "TextField_WKU63CA5OGG0", FieldName: "账号（废弃）", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "TextField_B9JO23R61XK0", FieldName: "付款方式", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "DDDateField_10RG38623CV40", FieldName: "预计转账时间", Converter: oa.StringConverter, Pointer: true},
// {ComponentId: "SignatureField_1KKQBJ6VLAWW0", FieldName: "审批人", Converter: oa.StringConverter},
// {ComponentId: "SignatureField_338G622D99Q0", FieldName: "商务部经理", Converter: oa.StringConverter},
// {ComponentId: "SignatureField_J1RW9LZQB480", FieldName: "分管领导", Converter: oa.StringConverter},

func (ef *BidExpenseForm) getApplyMappers() []oa.FieldMapper {
	return []oa.FieldMapper{
		{ComponentId: "TextField_22K742TU7DTS0", FieldName: "BillNo", Converter: oa.StringConverter, Version: "v1"}, // v1 手动编号 2025.11.01 之前
		{ComponentId: "SeqNumberField_1U1O5KHHSHUDC", FieldName: "BillNo", Converter: oa.StringConverter},           // v2 自动编号

		{ComponentId: "TextField_17HCU4VR5UQK0", FieldName: "ProjectName", Converter: oa.StringConverter},
		{ComponentId: "TextField_82D5H6HY60O0", FieldName: "ProjectCode", Converter: oa.StringConverter},
		{ComponentId: "TextField_8L38ISBPCSG0", FieldName: "BizRepName", Converter: oa.StringConverter},
		{ComponentId: "DDSelectField_RE7H7DQVRPS0", FieldName: "FeeType", Converter: oa.StringConverter},
		{ComponentId: "TextField_1SRLTMRI2F4W0", FieldName: "PayReason", Converter: oa.StringConverter, Pointer: true},

		{ComponentId: "TextField_3Y2MPH8USDG0", FieldName: "Purchaser", Converter: oa.StringConverter, Pointer: true},

		// 收款方信息 v1
		{ComponentId: "TextField_138818FKKQC00", FieldName: "PayeeName", Converter: oa.StringConverter},
		{ComponentId: "TextField_1HHDR2YRQGPS0", FieldName: "PayeeBank", Converter: oa.StringConverter},
		{ComponentId: "TextField_WKU63CA5OGG0", FieldName: "PayeeAccount", Converter: oa.StringConverter},

		{ComponentId: "TextField_1YIHQLKCJLUO0", FieldName: "PayRemark", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "MoneyField_QP9DYIENY4G0", FieldName: "PayAmount", Converter: oa.Float64Converter},
		{ComponentId: "TextField_B9JO23R61XK0", FieldName: "PayMethod", Converter: oa.StringConverter, Pointer: true},
		{ComponentId: "DDDateField_10RG38623CV40", FieldName: "PlanPayTime", Converter: oa.DateConverter(time.DateOnly, time.Local), Pointer: true},
	}
}

func (ef *BidExpenseForm) saveExpense(ctx context.Context, tx *ent.Tx) (*ent.BidExpense, error) {
	res, err := tx.BidExpense.Query().Where(bidexpense.ID(ef.ID)).Only(ctx)
	if ent.IsNotFound(err) {
		// 创建记录
		return ef.create(ctx, tx)
	} else if err != nil {
		return nil, err
	}

	// 流程已结束，直接返回
	if res.Done {
		return res, nil
	}

	// 未结束，则根据具体情况进行更新
	return ef.update(ctx, tx)
}

func (ef *BidExpenseForm) create(ctx context.Context, tx *ent.Tx) (*ent.BidExpense, error) {
	// 创建新支出记录
	expense := tx.BidExpense.Create()

	expense.SetID(ef.ID)
	expense.SetBusinessID(ef.BusinessId)
	expense.SetInstanceID(ef.InstanceId)
	expense.SetBillNo(ef.BillNo)

	expense.SetProjectName(ef.ProjectName)
	expense.SetProjectCode(ef.ProjectCode)
	expense.SetBizRepName(ef.BizRepName)
	if ef.Purchaser != nil {
		expense.SetPurchaser(*ef.Purchaser)
	}

	expense.SetNillableFeeTypeV1(ef.FeeTypeV1)
	if ef.FeeType != "" {
		expense.SetFeeType(bidexpense.FeeType(ef.FeeType))
	}

	if ef.PayReason != nil {
		expense.SetPayReason(*ef.PayReason)
	}

	expense.SetPayeeBank(ef.PayeeBank)
	expense.SetPayeeName(ef.PayeeName)
	expense.SetPayeeAccount(ef.PayeeAccount)
	if ef.PayRatio != nil {
		expense.SetPayRatio(*ef.PayRatio)
	}

	expense.SetPayAmount(ef.PayAmount)
	if ef.PayRemark != nil {
		expense.SetPayRemark(*ef.PayRemark)
	}
	if ef.PayMethod != nil {
		expense.SetPayMethod(*ef.PayMethod)
	}
	if ef.PlanPayTime != nil {
		expense.SetPlanPayTime(*ef.PlanPayTime)
	}

	expense.SetApprovalStatus(ef.ApprovalStatus)
	expense.SetDone(ef.Done)
	expense.SetCreateAt(*ef.CreateAt)
	expense.SetCreateBy(ef.CreateBy)
	if ef.UpdateAt == nil {
		expense.SetUpdateAt(*ef.CreateAt)
	} else {
		expense.SetUpdateAt(*ef.UpdateAt)
	}

	bidExpense, err := expense.Save(ctx)
	if err != nil {
		return nil, err
	}
	return bidExpense, nil
}

func (ef *BidExpenseForm) update(ctx context.Context, tx *ent.Tx) (*ent.BidExpense, error) {
	update := tx.BidExpense.UpdateOneID(ef.ID)

	// 更新审批状态及是否审批结束
	update.SetApprovalStatus(ef.ApprovalStatus)
	update.SetDone(ef.Done)

	// 更新修改时间戳
	if ef.UpdateAt != nil {
		update.SetUpdateAt(*ef.UpdateAt)
	}

	expense, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return expense, err
}
