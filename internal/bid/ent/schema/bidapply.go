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

// BidApply holds the schema definition for the BidApply entity.
type BidApply struct {
	ent.Schema
}

// Fields of the BidApply.
func (BidApply) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("投标申请 ID").MaxLen(32),
		field.String("business_id").Comment("审批编号").MaxLen(32),
		field.String("instance_id").Comment("审批实例 ID").MaxLen(64),
		field.String("project_id").Comment("项目 ID").MaxLen(32),

		field.String("purchaser_name").Comment("采购人名称").Optional().Nillable().MaxLen(64),
		field.Enum("bid_type").Comment("招标类型 UT:未知类型 OT:公开招标 IT:邀请招标 CN:竞争性谈判 IP:询价采购 SSP:单一来源采购 CC:竞争性磋商 SCT:自行招标 CIP:询比采购 HIP:医院自主采购 PC:比价 DP:直接采购").
			Values("UT", "OT", "IT", "CN", "IP", "SSP", "CC", "SCT", "CIP", "HIP", "PC", "DP").Default("UT"),
		field.String("agency_name").Comment("招标代理机构名称").Optional().Nillable().MaxLen(64),
		field.String("agency_contact").Comment("招标代理机构联系人及电话").Optional().Nillable().MaxLen(128),

		field.Time("opening_date").Comment("开标时间").Optional().Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime", // 适配MySQL datetime类型
			}),
		field.Text("notice_url").Comment("招标公告网址").Optional().Nillable(),
		field.Float("budget_amount").Comment("预算金额（元）").Default(0).SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(16,2)", // Override MySQL.
			},
		),

		field.Text("remark").Comment("备注说明;如资质要求、技术难点、事项说明等").Optional().Nillable(),
		field.JSON("attachments", []map[string]any{}).Comment("投标报名相关附件").Optional(),

		field.String("approval_status").Comment("审批状态"),
		field.Bool("done").Comment("审批流程是否已结束").Default(false),
	}

}

// Edges of the BidApply.
func (BidApply) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", BidProject.Type).Ref("apply").
			Field("project_id").Unique().Required(),
	}
}

// Indexes of the BidApply.
func (BidApply) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("approval_status").StorageKey("idx_approval_status"),
	}
}

// Annotations of the BidApply.
func (BidApply) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("bid_apply"),
		entsql.WithComments(true),
		schema.Comment("投标申请"),
	}
}

func (BidApply) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
