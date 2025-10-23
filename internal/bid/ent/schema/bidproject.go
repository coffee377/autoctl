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
		field.String("code").Comment("项目编码").MaxRuneLen(64).Default(""),
		field.String("name").Comment("项目名称").MaxRuneLen(64),
		field.Enum("type").Comment("项目类型 UP:未知 S:软件 H:硬件 SHI:软硬集成 OM:运维").
			Values("UP", "S", "H", "SHI", "OM").Default("UP"),
		field.String("department_code").Comment("所属部门编码").MaxRuneLen(64).Default(""),
		field.String("department_name").Comment("所属部门名称").MaxRuneLen(64),
		field.String("biz_rep_no").Comment("商务代表工号").MaxRuneLen(8).Default(""),
		field.String("biz_rep_name").Comment("商务代表").MaxRuneLen(16).Default(""),
	}
}

// Edges of the BidProject.
func (BidProject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("apply", BidApply.Type).Unique().Annotations(
			entsql.OnDelete(entsql.Cascade),
		).StorageKey(edge.Symbol("fk_pid_01")),

		edge.To("expense", BidExpense.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		).StorageKey(edge.Symbol("fk_pid_02")),
	}
}

// Indexes of the BidProject.
func (BidProject) Indexes() []ent.Index {
	return []ent.Index{
		// todo 历史数据项目名称存在重复（可能多次提交记录）
		//index.Fields("code", "name").Unique().StorageKey("uk_code_name"),
		index.Fields("code", "name").StorageKey("idx_code_name"),
		index.Fields("department_code").StorageKey("idx_department_code"),
	}
}

// Annotations of the BidProject.
func (BidProject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Table("bid_project"),
		schema.Comment("投标项目"),
	}
}

func (BidProject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
