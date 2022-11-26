package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/avptp/brain/internal/data/rules"
	"github.com/google/uuid"
)

type PersonOwnedMixin struct {
	mixin.Schema
}

func (PersonOwnedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("person_id", uuid.UUID{}).
			Immutable(),
	}
}

func (PersonOwnedMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("person", Person.Type).
			Ref("authentications").
			Field("person_id").
			Unique().
			Required().
			Immutable().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (PersonOwnedMixin) Policy() ent.Policy {
	return rules.FilterPersonOwnedRule()
}
