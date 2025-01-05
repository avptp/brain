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
	"github.com/avptp/brain/internal/data/validation"
	"github.com/avptp/brain/internal/generated/data/privacy"
)

type Person struct {
	ent.Schema
}

func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", types.ID{}).
			Immutable().
			Annotations(
				entsql.DefaultExpr("uuid_generate_v4()"),
			),
		field.String("stripe_id").
			SchemaType(map[string]string{
				dialect.Postgres: "string(255)",
			}).
			Optional().
			Nillable().
			Unique(),
		field.String("email").
			SchemaType(map[string]string{
				dialect.Postgres: "string(254)",
			}).
			MaxLen(254).
			Match(validation.EmailRegexp).
			Unique().
			StructTag(`faker:"email"`),
		field.Time("email_verified_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Optional().
			Nillable().
			StructTag(`faker:"-"`),
		field.String("phone").
			SchemaType(map[string]string{
				dialect.Postgres: "string(16)",
			}).
			MaxLen(16).
			Validate(validation.Phone).
			Optional().
			Nillable().
			Unique().
			StructTag(`faker:"phone"`),
		field.String("password").
			SchemaType(map[string]string{
				dialect.Postgres: "char(97)",
			}).
			Sensitive(),
		field.String("tax_id").
			SchemaType(map[string]string{
				dialect.Postgres: "char(9)",
			}).
			MaxLen(9).
			Validate(validation.TaxID).
			Unique().
			StructTag(`faker:"tax_id"`),
		field.String("first_name").
			SchemaType(map[string]string{
				dialect.Postgres: "string(50)",
			}).
			MaxLen(50).
			NotEmpty().
			StructTag(`faker:"first_name"`),
		field.String("last_name").
			SchemaType(map[string]string{
				dialect.Postgres: "string(50)",
			}).
			MaxLen(50).
			NotEmpty().
			Optional().
			Nillable().
			StructTag(`faker:"last_name"`),
		field.String("language").
			SchemaType(map[string]string{
				dialect.Postgres: "char(2)",
			}).
			StructTag(`faker:"oneof:ca,en,es"`),
		field.Time("birthdate").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}).
			Optional().
			Nillable().
			StructTag(`faker:"birthdate"`),
		// We don't care about people's gender, but we ask for this field
		// voluntarily to make statistics on the underrepresentation of women.
		field.Enum("gender").
			SchemaType(map[string]string{
				dialect.Postgres: "gender",
			}).
			NamedValues(
				"Woman", "woman",
				"Man", "man",
				"NonBinary", "nonbinary",
			).
			Optional().
			Nillable().
			StructTag(`faker:"oneof:woman,man,nonbinary"`),
		field.String("address").
			SchemaType(map[string]string{
				dialect.Postgres: "string(100)",
			}).
			MaxLen(100).
			NotEmpty().
			Optional().
			Nillable().
			StructTag(`faker:"address"`),
		field.String("postal_code").
			SchemaType(map[string]string{
				dialect.Postgres: "string(10)",
			}).
			MaxLen(10).
			NotEmpty().
			Optional().
			Nillable().
			StructTag(`faker:"postal_code"`),
		field.String("city").
			SchemaType(map[string]string{
				dialect.Postgres: "string(58)",
			}).
			MaxLen(58).
			NotEmpty().
			Optional().
			Nillable().
			StructTag(`faker:"city"`),
		field.String("country").
			SchemaType(map[string]string{
				dialect.Postgres: "char(2)",
			}).
			Validate(validation.Country).
			Optional().
			Nillable().
			StructTag(`faker:"country"`),
		field.Bool("subscribed").
			Default(false).
			StructTag(`faker:"-"`),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("authentications", Authentication.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("authorizations", Authorization.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Person) Hooks() []ent.Hook {
	return []ent.Hook{
		mutators.PersonEmail,
		mutators.PersonPassword,
		mutators.PersonLanguage,
		mutators.PersonBirthdate,
	}
}

func (Person) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rules.FilterPersonRule(),
		},
		Mutation: privacy.MutationPolicy{
			rules.FilterPersonRule(),
		},
	}
}
