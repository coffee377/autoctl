package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// BidAccountRelation holds the schema definition for the BidAccountRelation entity.
type BidAccountRelation struct {
	ent.Schema
}

// Fields of the BidAccountRelation.
func (BidAccountRelation) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").Comment("会员账号ID").NotEmpty(),
		field.String("ca_id").Comment("CA证书ID").NotEmpty(),
		field.String("remark").Comment("关联备注").Optional(),
		field.Time("bind_at").Default(time.Now),
	}
}

func (BidAccountRelation) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一索引：防止重复关联
		index.Fields("account_id", "ca_id").Unique().StorageKey("uk_account_id_ca_id"),
		// 普通索引：优化查询性能
		index.Fields("account_id").StorageKey("idx_account_id"),
		index.Fields("ca_id").StorageKey("idx_ca_id"),
	}
}

// Edges of the BidAccountRelation.
func (BidAccountRelation) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联到会员账号
		edge.From("account", BidMemberAccount.Type).
			Ref("ca_relations").
			Field("account_id").
			Unique().
			Required(),
		// 关联到CA证书
		edge.From("ca_certificate", BidCACertificate.Type).
			Ref("account_relations").
			Field("ca_id").
			Unique().
			Required(),
	}
}

// Annotations of the BidWebSite.
func (BidAccountRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("bid_account_ca_rel"),
		entsql.WithComments(true),
		schema.Comment("账号证书关联关系"),
	}
}
