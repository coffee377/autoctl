package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// BidExpense holds the schema definition for the BidExpense entity.
type BidExpense struct {
	ent.Schema
}

// Fields of the BidExpense.
func (BidExpense) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("投标支出 ID").MaxLen(32),
		field.String("business_id").Comment("审批编号").MaxLen(32),
		field.String("instance_id").Comment("审批实例 ID").MaxLen(64),
		field.String("bill_no").Comment("单据编码").MaxRuneLen(16),

		field.String("project_id").Comment("项目 ID").MaxRuneLen(32).Optional().Nillable(),
		// todo 临时字段，拉取数据，后期删除
		field.String("project_name").Comment("项目名称").MaxRuneLen(64),
		field.String("project_code").Comment("项目编码").MaxRuneLen(64),
		field.String("biz_rep_name").Comment("商务代表").MaxRuneLen(16),
		field.String("purchaser").Comment("采购人名称").MaxRuneLen(64).Optional().Nillable(),

		field.String("fee_type_v1").Comment("费用类型").Optional().Nillable(),
		field.Enum("fee_type").Comment("费用类型 RF:报名费 DF:标书工本费 CA:CA费用 EF:专家费 BB:投标保证金 BS:中标服务费 PB:履约保证金 PG:预付保函 OE:其他费用 EB:入库保证金").
			Values("RF", "DF", "CA", "EF", "BB", "BS", "PB", "PG", "OE", "EB").Default("OE"),
		field.String("pay_reason").Comment("付款事由").MaxRuneLen(64).Optional().Nillable(),

		field.Bool("refunded").Comment("是否（保证金）退还金额").Default(false),
		field.String("payee_bank").Comment("收款方开户银行").MaxRuneLen(64),
		field.String("payee_name").Comment("收款方账户名称").MaxRuneLen(64),
		field.String("payee_account").Comment("收款方账号").MaxRuneLen(64),
		field.Float("pay_ratio").Comment("付款比例").Default(100).SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(5,2)",
			},
		),
		field.Float("pay_amount").Comment("付款金额（元）").Default(0).SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(16,2)",
			},
		),
		// 付款主体
		// 付款账户 与付款主体联动
		// 核算项目
		field.Text("pay_remark").Comment("付款备注").Optional().Nillable(),

		field.String("pay_method").Comment("付款方式").MaxRuneLen(64).Optional().Nillable(),
		field.Time("plan_pay_time").Comment("预计转账时间").Optional().Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime", // 适配MySQL datetime类型
			}),
		field.String("approval_status").Comment("费用审批状态"),
		field.Bool("done").Comment("审批流程是否已结束").Default(false),
	}
}

// Edges of the BidExpense.
func (BidExpense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", BidProject.Type).Ref("expense").Unique().
			Field("project_id"),
	}
}

// Indexes of the BidExpense.
func (BidExpense) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("bill_no").StorageKey("idx_bill_no"),
		index.Fields("fee_type").StorageKey("idx_fee_type"),
		index.Fields("approval_status").StorageKey("idx_approval_status"),
	}
}

// Annotations of the BidExpense.
func (BidExpense) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Table("bid_expense"),
		schema.Comment("投标费用支出"),
	}
}

func (BidExpense) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

type RecipientAccount struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	CardNo         string `json:"cardNo,omitempty"`
	InstBranchCode string `json:"instBranchCode,omitempty"`
	IdentityType   string `json:"identityType,omitempty"`
	InstProvince   string `json:"instProvince,omitempty"`
	InstCity       string `json:"instCity,omitempty"`
	InstCode       string `json:"instCode,omitempty"`
	InstName       string `json:"instName,omitempty"`
	InstBranchName string `json:"instBranchName,omitempty"`
	Source         string `json:"source,omitempty"`
}
