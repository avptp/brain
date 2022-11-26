// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthenticationsColumns holds the columns for the "authentications" table.
	AuthenticationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Default: "gen_random_ulid()"},
		{Name: "token", Type: field.TypeBytes, Unique: true, SchemaType: map[string]string{"postgres": "bytes"}},
		{Name: "created_ip", Type: field.TypeString, SchemaType: map[string]string{"postgres": "inet"}},
		{Name: "last_used_ip", Type: field.TypeString, SchemaType: map[string]string{"postgres": "inet"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "last_used_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "person_id", Type: field.TypeUUID},
	}
	// AuthenticationsTable holds the schema information for the "authentications" table.
	AuthenticationsTable = &schema.Table{
		Name:       "authentications",
		Columns:    AuthenticationsColumns,
		PrimaryKey: []*schema.Column{AuthenticationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "authentications_persons_authentications",
				Columns:    []*schema.Column{AuthenticationsColumns[6]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Default: "gen_random_ulid()"},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 254, SchemaType: map[string]string{"postgres": "string(254)"}},
		{Name: "email_verified_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true, Size: 16, SchemaType: map[string]string{"postgres": "string(16)"}},
		{Name: "password", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(97)"}},
		{Name: "tax_id", Type: field.TypeString, Unique: true, Size: 9, SchemaType: map[string]string{"postgres": "char(9)"}},
		{Name: "first_name", Type: field.TypeString, Size: 50, SchemaType: map[string]string{"postgres": "string(50)"}},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 50, SchemaType: map[string]string{"postgres": "string(50)"}},
		{Name: "language", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(2)"}},
		{Name: "birthdate", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "gender", Type: field.TypeEnum, Nullable: true, Enums: []string{"woman", "man", "nonbinary"}, SchemaType: map[string]string{"postgres": "gender"}},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 100, SchemaType: map[string]string{"postgres": "string(100)"}},
		{Name: "postal_code", Type: field.TypeString, Nullable: true, Size: 10, SchemaType: map[string]string{"postgres": "string(10)"}},
		{Name: "city", Type: field.TypeString, Nullable: true, Size: 58, SchemaType: map[string]string{"postgres": "string(58)"}},
		{Name: "country", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "char(2)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthenticationsTable,
		PersonsTable,
	}
)

func init() {
	AuthenticationsTable.ForeignKeys[0].RefTable = PersonsTable
}
