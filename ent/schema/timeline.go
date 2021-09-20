package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Timeline holds the schema definition for the Timeline entity.
type Timeline struct {
	ent.Schema
}

// Fields of the Timeline.
func (Timeline) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("user_id"),
		field.Uint64("meh_id"),
	}
}

func (Timeline) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixIn{},
		TimeMixin{},
	}
}

// Edges of the Timeline.
func (Timeline) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required(),
		edge.To("meh", Meh.Type).
			Field("meh_id").
			Unique().
			Required(),
	}
}

func (Timeline) Index() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "meh_id").Unique(),
	}
}
