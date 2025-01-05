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
			StructTag(`fake:"{email}"`),
		field.Time("email_verified_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Optional().
			Nillable(),
		field.String("phone").
			SchemaType(map[string]string{
				dialect.Postgres: "string(16)",
			}).
			MaxLen(16).
			Validate(validation.Phone).
			Optional().
			Nillable().
			Unique().
			StructTag(`fake:"{phone_e164}"`),
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
			StructTag(`fake:"{tax_id}"`),
		field.String("first_name").
			SchemaType(map[string]string{
				dialect.Postgres: "string(50)",
			}).
			MaxLen(50).
			NotEmpty().
			StructTag(`fake:"{firstname}"`),
		field.String("last_name").
			SchemaType(map[string]string{
				dialect.Postgres: "string(50)",
			}).
			MaxLen(50).
			NotEmpty().
			Optional().
			Nillable().
			StructTag(`fake:"{lastname}"`),
		field.String("language").
			SchemaType(map[string]string{
				dialect.Postgres: "char(2)",
			}).
			StructTag(`fake:"{randomstring:[ca,es,en]}"`),
		field.Time("birthdate").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}).
			Optional().
			Nillable().
			StructTag(`fake:"{date}"`),
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
			StructTag(`fake:"{randomstring:[woman,man,nonbinary]}"`),
		field.String("address").
			SchemaType(map[string]string{
				dialect.Postgres: "string(100)",
			}).
			MaxLen(100).
			NotEmpty().
			Optional().
			Nillable().
			StructTag(`fake:"{street}"`),
		field.String("postal_code").
			SchemaType(map[string]string{
				dialect.Postgres: "string(10)",
			}).
			MaxLen(10).
			NotEmpty().
			Optional().
			Nillable().
			StructTag(`fake:"{zip}"`),
		field.String("city").
			SchemaType(map[string]string{
				dialect.Postgres: "string(58)",
			}).
			MaxLen(58).
			NotEmpty().
			Optional().
			Nillable().
			StructTag(`fake:"{city}"`),
		field.String("country").
			SchemaType(map[string]string{
				dialect.Postgres: "char(2)",
			}).
			Validate(validation.Country).
			Optional().
			Nillable().
			StructTag(`fake:"{countryabr}"`),
		field.Bool("subscribed").
			Default(false),
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
