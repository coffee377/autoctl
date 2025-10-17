package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// BidProject holds the schema definition for the BidProject entity.
type BidProject struct {
	ent.Schema
}

// Fields of the BidProject.
func (BidProject) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Comment("项目 ID").MaxLen(32),
		field.String("name").Comment("项目名称").MaxLen(64),
		field.String("code").Comment("项目编码").MaxRuneLen(64),
		field.Enum("type").Comment("项目类型 S:软件 H:硬件 I:软硬集成 O:运维").
			Values("S", "H", "I", "O").Optional().Nillable(), // todo 必填
		field.String("department_code").Comment("所属部门编码").MaxRuneLen(64),
		field.String("department_name").Comment("所属部门名称").MaxRuneLen(64),
		field.String("biz_rep_no").Comment("商务代表工号").MaxRuneLen(8).Default(""),
		field.String("biz_rep_name").Comment("商务代表").MaxRuneLen(16),
	}
}

// Edges of the BidProject.
func (BidProject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("apply", BidApply.Type).Unique().Annotations(
			entsql.OnDelete(entsql.Cascade),
		).StorageKey(edge.Symbol("fk_project_id")),
	}
}

// Indexes of the BidProject.
func (BidProject) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code", "name").Unique().StorageKey("uk_code_name"),
		index.Fields("name").StorageKey("idx_name"),
		index.Fields("department_code").StorageKey("idx_department_code"),
	}
}

// Annotations of the BidProject.
func (BidProject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "bid_project"},
		schema.Comment("投标项目"),
	}
}
