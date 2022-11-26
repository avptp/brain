package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/avptp/brain/internal/data/mutators"
	"github.com/google/uuid"
)

type Authentication struct {
	ent.Schema
}

func (Authentication) Mixin() []ent.Mixin {
	return []ent.Mixin{
		PersonOwnedMixin{},
	}
}

func (Authentication) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Annotations(
				entsql.Annotation{
					Default: "gen_random_ulid()",
				},
			),
		field.Bytes("token").
			SchemaType(map[string]string{
				dialect.Postgres: "bytes",
			}).
			Immutable().
			Unique(),
		field.String("created_ip").
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			StructTag(`fake:"{ipv6address}"`),
		field.String("last_used_ip").
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			StructTag(`fake:"{ipv6address}"`),
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

func (Authentication) Hooks() []ent.Hook {
	return []ent.Hook{
		mutators.AuthenticationToken,
	}
}
