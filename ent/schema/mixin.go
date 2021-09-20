package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type IDMixIn struct {
	mixin.Schema
}

func (IDMixIn) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
	}
}

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Annotations(&entsql.Annotation{
			Default: "CURRENT_TIMESTAMP",
		}).Immutable(),
		field.Time("updated_at").Default(time.Now).Annotations(&entsql.Annotation{
			Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
		}).Immutable(),
	}
}
