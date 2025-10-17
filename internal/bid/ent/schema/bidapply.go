package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BidApply holds the schema definition for the BidApply entity.
type BidApply struct {
	ent.Schema
}

// Fields of the BidApply.
func (BidApply) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("name").MaxLen(32),
		field.String("business_id").Comment("审批编号").MaxLen(32),
		field.String("instance_id").Comment("审批实例 ID").MaxLen(32),
		field.String("project_id").Comment("项目 ID").MaxLen(32),

		field.String("purchaser").Comment("采购人名称").Optional().Nillable().MaxLen(64),
		field.String("bid_type").Comment("招标类型").Optional().Nillable().MaxLen(11),
		field.String("agency_name").Comment("招标代理机构名称").Optional().Nillable().MaxLen(64),
		field.String("agency_contact").Comment("招标代理机构联系人及电话").Optional().Nillable().MaxLen(128),

		field.Time("opening_date").Comment("开标时间").Optional().Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime", // 适配MySQL datetime类型
			}),
		field.Text("notice_url").Comment("招标公告网址").Optional().Nillable(),
		field.Int64("budget_amount").Comment("预算金额（元）").Default(0),
		field.Text("remark").Comment("备注说明;如资质要求、技术难点、事项说明等").Optional().Nillable(),
		field.JSON("attachments", []map[string]any{}).Comment("附件"),

		field.String("status").Comment("审批状态"),
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
	return []ent.Index{}
}

// Annotations of the BidApply.
func (BidApply) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "bid_apply"},
		schema.Comment("投标申请"),
	}
}
