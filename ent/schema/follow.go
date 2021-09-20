package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Follow holds the schema definition for the Follow entity.
type Follow struct {
	ent.Schema
}

// Fields of the Follow.
func (Follow) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("user_id"),
		field.Uint64("followee_id"),
	}
}

func (Follow) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixIn{},
		TimeMixin{},
	}
}

// Edges of the Follow.
func (Follow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required(),
		edge.To("followee", User.Type).
			Field("followee_id").
			Unique().
			Required(),
	}
}

func (User) Index() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "followee_id").
			Unique(),
		index.Fields("followee_id"),
	}
}
