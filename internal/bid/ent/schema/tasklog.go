package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TaskLog holds the schema definition for the TaskLog entity.
type TaskLog struct {
	ent.Schema
}

// Fields of the TaskLog.
func (TaskLog) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("biz_type").Comment("任务流转业务类型 1:商品采购 2:项目投标 0:其他").
			Values("1", "2", "0").Default("0"),
		field.String("biz_id").Comment("业务标识").MaxRuneLen(32),
		field.Uint32("assign_seq").Comment("同一业务标识下的指派序号（从1开始递增）"),

		field.Time("assign_time").Comment("任务指派时间").Optional().Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.String("handler_no").Comment("受理人工号").Optional().Nillable().MaxRuneLen(8),

		field.Time("start_time").Comment("任务开始时间").Optional().Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("end_time").Comment("任务结束时间（任务完成/终止的时间）").Optional().Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),

		field.Text("remark").Comment("备注（如重新指派原因、任务终止说明等）").Optional().Nillable(),
	}
}

// Edges of the TaskLog.
func (TaskLog) Edges() []ent.Edge {
	return nil
}

func (TaskLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("biz_type", "biz_id", "assign_seq").Unique().StorageKey("uk_bti_as"),
	}
}

func (TaskLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_task_log"),
		entsql.WithComments(true),
		schema.Comment("任务流转记录"),
	}
}

func (TaskLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
