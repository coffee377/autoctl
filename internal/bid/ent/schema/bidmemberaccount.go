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

// BidMemberAccount holds the schema definition for the BidMemberAccount entity.
type BidMemberAccount struct {
	ent.Schema
}

// Fields of the BidMemberAccount.
func (BidMemberAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxRuneLen(32).Default(uuid.New().String()),
		field.String("website_id").Comment("归属网站").MaxRuneLen(32),
		field.String("owner_code").Comment("归属主体编码").Nillable().MaxRuneLen(32),
		field.String("owner_name").Comment("归属主体名称").Nillable().MaxRuneLen(32),
		field.String("username").Comment("账号").NotEmpty().MaxRuneLen(32),
		field.String("password").Sensitive().Optional().Comment("密码").MaxRuneLen(32),

		field.String("register_person").Optional().Comment("注册人员").MaxRuneLen(16),
		field.String("register_mobile").Optional().Comment(" 注册手机号").MaxRuneLen(16),

		field.String("primary_ca_id").Comment("主证书 ID").MaxRuneLen(32).Optional().Nillable(),

		field.Enum("account_status").
			Values("active", "inactive", "abandoned", "suspended").
			Default("active").
			Comment("账号状态: active-正常/inactive-未激活/abandoned-废弃/suspended-暂停"),
		field.Text("abandon_reason").Optional().Comment("废弃原因"),
		field.Text("remark").Optional().Comment("备注"),
	}
}

// Edges of the BidMemberAccount.
func (BidMemberAccount) Edges() []ent.Edge {
	return []ent.Edge{
		// 定义多对一关系：一个 MemberAccount 属于一个 Website
		edge.From("website", BidWebSite.Type). // edge.From(边名, 关联实体类型)
			// 关联Website的member_accounts边（双向绑定）
			Ref("member_accounts").
			// 关键：确保一个账号只属于一个网站（一对多核心）
			Unique().
			// 强制外键非空（账号必须归属某个网站）
			Required().
			// 显式指定外键字段名（默认是website_id）
			Field("website_id"),

		// 一个账号可以关联多个CA证书（通过中间表）
		edge.To("ca_relations", BidAccountRelation.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		).StorageKey(edge.Symbol("fk_account_id")),
	}
}

// Indexes of the BidMemberAccount.
func (BidMemberAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("website_id").StorageKey("idx_website_id"),
	}
}

// Annotations of the BidMemberAccount.
func (BidMemberAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("bid_member_account"),
		entsql.WithComments(true),
		schema.Comment("会员账号"),
	}
}

func (BidMemberAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
