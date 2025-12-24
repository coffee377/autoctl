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

// BidInfo holds the schema definition for the BidInfo entity.
type BidInfo struct {
	ent.Schema
}

// Fields of the BidInfo.
func (BidInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("投标信息 ID").MaxLen(32),
		field.String("project_id").Comment("项目 ID").MaxLen(32),
		field.String("group_leader").Comment("投标组长工号").MaxRuneLen(8).Default(""),
		field.String("group_leader_name").Comment("投标组长").MaxRuneLen(8).Default(""),
		field.String("bid_subject_code").Comment("投标主体编码").MaxRuneLen(32).Optional().Nillable(),
		field.String("bid_subject_name").Comment("投标主体名称").MaxRuneLen(32).Optional().Nillable(),
		field.Float("bid_amount").Comment("投标金额").Default(0).SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(16,2)",
			},
		),
		field.Enum("bid_status").Comment("投标状态 RP:待报名 RO:报名中 RS:报名成功 RF:报名失败 DP:标书编制中 B:投标中 W:已中标 L:未中标 F:流标 A:弃标 0:-").
			Values("RP", "RO", "RS", "RF", "DP", "B", "W", "L", "F", "A", "0").Default("0"),
		field.Time("bid_date").Comment("中标时间").Optional().Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Float("software_amount").Comment("中标软件金额").Default(0).SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(16,2)",
			},
		),
		field.Float("hardware_amount").Comment("中标硬件金额").Default(0).SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(16,2)",
			},
		),
		field.Float("operation_amount").Comment("中标运维金额").Default(0).SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(16,2)",
			},
		),
		field.Text("result_url").Comment("中标结果网址").Optional().Nillable(),

		field.Bool("contract_signed").Comment("销售合同是否签署").Default(false),
		field.String("contract_no").Comment("销售合同号").MaxRuneLen(64).Optional().Nillable(),
		field.Time("contract_sign_date").Comment("销售合同签署日期").Optional().Nillable().SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Text("remark").Comment("备注信息").Optional().Nillable(),
	}
}

// Edges of the BidInfo.
func (BidInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", BidProject.Type).Ref("info").
			Field("project_id").Unique().Required(),
	}
}

// Indexes of the BidInfo.
func (BidInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("bid_status").StorageKey("idx_bid_status"),
	}
}

// Annotations of the BidInfo.
func (BidInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Table("bid_info"),
		schema.Comment("投标信息"),
	}
}

func (BidInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
