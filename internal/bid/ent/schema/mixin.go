package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/coffee377/entcc"
)

type BaseMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_at").Comment("创建时间").Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime(3)",
		}).Annotations(entcc.Sort(1), entcc.Tail(true)),
		field.String("create_by").Comment("创建人").Optional().Nillable().MaxLen(32).
			Annotations(entcc.TailSort(2)),
		field.Time("update_at").Comment("更新时间").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime(3)",
		}).Annotations(entcc.TailSort(3)),
		field.String("update_by").Comment("更新人").Optional().Nillable().MaxLen(32).Annotations(entcc.TailSort(4)),
	}
}

func (BaseMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entcc.WithFieldSort(true),
	}
}
