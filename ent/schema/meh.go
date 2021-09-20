package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Meh holds the schema definition for the Meh entity.
type Meh struct {
	ent.Schema
}

// Fields of the Meh.
func (Meh) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("user_id"),
		field.String("text").NotEmpty(),
	}
}

func (Meh) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixIn{},
		TimeMixin{},
	}
}

// Edges of the Meh.
func (Meh) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required(),
	}
}

func (Meh) Index() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
