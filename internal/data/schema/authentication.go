package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/avptp/brain/internal/api/types"
	"github.com/avptp/brain/internal/data/mutators"
	"github.com/avptp/brain/internal/data/rules"
	"github.com/avptp/brain/internal/generated/data/hook"
)

type Authentication struct {
	ent.Schema
}

func (Authentication) Fields() []ent.Field {
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
			StructTag(`faker:"slice_len=64"`),
		field.String("created_ip").
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			StructTag(`faker:"ipv6"`),
		field.String("last_used_ip").
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			StructTag(`faker:"ipv6"`),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Default(time.Now).
			Immutable(),
		field.Time("last_used_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Authentication) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("person", Person.Type).
			Ref("authentications").
			Field("person_id").
			Unique().
			Required().
			Immutable(),
	}
}

func (Authentication) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(mutators.AuthenticationToken, ent.OpCreate),
	}
}

func (Authentication) Policy() ent.Policy {
	return rules.FilterPersonOwnedRule()
}
