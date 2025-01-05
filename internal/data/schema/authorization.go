package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/avptp/brain/internal/api/types"
	"github.com/avptp/brain/internal/data/mutators"
	"github.com/avptp/brain/internal/data/rules"
	"github.com/avptp/brain/internal/generated/data/hook"
)

type Authorization struct {
	ent.Schema
}

func (Authorization) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", types.ID{}).
			Immutable().
			Annotations(
				entsql.DefaultExpr("uuid_generate_v4()"),
			),
		field.UUID("person_id", types.ID{}).
			Immutable(),
		field.Bytes("token").
			SchemaType(map[string]string{
				dialect.Postgres: "bytes",
			}).
			Immutable().
			Unique().
			StructTag(`fakesize:"64"`),
		field.Enum("kind").
			SchemaType(map[string]string{
				dialect.Postgres: "authorization_kind",
			}).
			NamedValues(
				"Email", "email",
				"Password", "password",
			).
			StructTag(`fake:"{randomstring:[email,password]}"`),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Default(time.Now).
			Immutable(),
	}
}

func (Authorization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("person", Person.Type).
			Ref("authorizations").
			Field("person_id").
			Unique().
			Required().
			Immutable(),
	}
}

func (Authorization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("person_id", "kind").
			Unique(),
	}
}

func (Authorization) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(mutators.AuthorizationToken, ent.OpCreate),
	}
}

func (Authorization) Policy() ent.Policy {
	return rules.FilterPersonOwnedRule()
}
