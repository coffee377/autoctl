package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// BidCACertificate holds the schema definition for the BidCACertificate entity.
type BidCACertificate struct {
	ent.Schema
}

// Fields of the BidCACertificate.
func (BidCACertificate) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxRuneLen(32).Default(uuid.New().String()),
		field.String("code").Comment("CA证书编码").NotEmpty(),
		field.String("name").Comment("CA证书名称").NotEmpty(),
		field.Time("expiry_time").Comment("过期时间"),
		field.String("password").Sensitive().Optional().Comment("CA证书密码"),
		field.Text("remark").Optional().Comment("备注"),
		field.Bool("primary").Default(false).Comment("是否为主证书"),
		field.Time("last_renewal_at").Optional().Comment("最后续费时间"),
	}
}

// Edges of the BidCACertificate.
func (BidCACertificate) Edges() []ent.Edge {
	return []ent.Edge{
		// 一个CA证书可以被多个账号使用（通过中间表）
		edge.To("account_relations", BidAccountRelation.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		).StorageKey(edge.Symbol("fk_ca_id")),
	}
}

// Indexes of the BidCACertificate.
func (BidCACertificate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").StorageKey("idx_code"),
		index.Fields("name").StorageKey("idx_name"),
	}
}

// Annotations of the BidCACertificate.
func (BidCACertificate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("bid_ca_certificate"),
		entsql.WithComments(true),
		schema.Comment("CA 证书"),
	}
}

func (BidCACertificate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
