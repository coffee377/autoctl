package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// BidExpense holds the schema definition for the BidExpense entity.
type BidExpense struct {
	ent.Schema
}

// Fields of the BidExpense.
func (BidExpense) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the BidExpense.
func (BidExpense) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the BidExpense.
func (BidExpense) Indexes() []ent.Index {
	return []ent.Index{}
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
