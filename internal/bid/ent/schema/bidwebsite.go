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

// BidWebSite holds the schema definition for the BidWebSite entity.
type BidWebSite struct {
	ent.Schema
}

// Fields of the BidWebSite.
func (BidWebSite) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxRuneLen(32).Default(uuid.New().String()),
		field.String("province").Comment("省份").MaxRuneLen(8),
		field.String("city").Comment("地市").MaxRuneLen(8).Nillable().Optional(),
		field.String("name").Comment("网站名称").MaxRuneLen(32).NotEmpty(),
		field.Text("url").Comment("网站地址"),
		field.Bool("active").Comment("是否启用").Default(true),
		field.Text("remark").Comment("备注").Optional(),
	}
}

// Edges of the BidWebSite.
func (BidWebSite) Edges() []ent.Edge {
	return []ent.Edge{
		// 一个网站有多个会员账号
		edge.To("member_accounts", BidMemberAccount.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		).StorageKey(edge.Symbol("fk_website_id")),
	}
}

// Indexes 网站信息索引定义
func (BidWebSite) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("province", "city").StorageKey("idx_province_city"),
		index.Fields("name").StorageKey("idx_name"),
	}
}

// Annotations of the BidWebSite.
func (BidWebSite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("bid_web_site"),
		entsql.WithComments(true),
		schema.Comment("网站信息"),
	}
}

func (BidWebSite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
