package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// BidInfo holds the schema definition for the BidInfo entity.
type BidInfo struct {
	ent.Schema
}

// Fields of the BidInfo.
func (BidInfo) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the BidInfo.
func (BidInfo) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the BidInfo.
func (BidInfo) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations of the BidInfo.
func (BidInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "bid_info"},
		schema.Comment("投标信息"),
	}
}
