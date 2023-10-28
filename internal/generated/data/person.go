// Code generated by ent, DO NOT EDIT.

package data

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/google/uuid"
)

// Person is the model entity for the Person schema.
type Person struct {
	config `fake:"-" json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// StripeID holds the value of the "stripe_id" field.
	StripeID *string `json:"stripe_id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty" fake:"{email}"`
	// EmailVerifiedAt holds the value of the "email_verified_at" field.
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone *string `json:"phone,omitempty" fake:"{phone_e164}"`
	// Password holds the value of the "password" field.
	Password string `fake:"-" json:"-"`
	// TaxID holds the value of the "tax_id" field.
	TaxID string `json:"tax_id,omitempty" fake:"{tax_id}"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty" fake:"{firstname}"`
	// LastName holds the value of the "last_name" field.
	LastName *string `json:"last_name,omitempty" fake:"{lastname}"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty" fake:"{randomstring:[ca,es,en]}"`
	// Birthdate holds the value of the "birthdate" field.
	Birthdate *time.Time `json:"birthdate,omitempty" fake:"{date}"`
	// Gender holds the value of the "gender" field.
	Gender *person.Gender `json:"gender,omitempty" fake:"{randomstring:[woman,man,nonbinary]}"`
	// Address holds the value of the "address" field.
	Address *string `json:"address,omitempty" fake:"{street}"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode *string `json:"postal_code,omitempty" fake:"{zip}"`
	// City holds the value of the "city" field.
	City *string `json:"city,omitempty" fake:"{city}"`
	// Country holds the value of the "country" field.
	Country *string `json:"country,omitempty" fake:"{countryabr}"`
	// Subscribed holds the value of the "subscribed" field.
	Subscribed bool `json:"subscribed,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonQuery when eager-loading is set.
	Edges        PersonEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PersonEdges holds the relations/edges for other nodes in the graph.
type PersonEdges struct {
	// Authentications holds the value of the authentications edge.
	Authentications []*Authentication `json:"authentications,omitempty"`
	// Authorizations holds the value of the authorizations edge.
	Authorizations []*Authorization `json:"authorizations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedAuthentications map[string][]*Authentication
	namedAuthorizations  map[string][]*Authorization
}

// AuthenticationsOrErr returns the Authentications value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) AuthenticationsOrErr() ([]*Authentication, error) {
	if e.loadedTypes[0] {
		return e.Authentications, nil
	}
	return nil, &NotLoadedError{edge: "authentications"}
}

// AuthorizationsOrErr returns the Authorizations value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) AuthorizationsOrErr() ([]*Authorization, error) {
	if e.loadedTypes[1] {
		return e.Authorizations, nil
	}
	return nil, &NotLoadedError{edge: "authorizations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Person) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case person.FieldSubscribed:
			values[i] = new(sql.NullBool)
		case person.FieldStripeID, person.FieldEmail, person.FieldPhone, person.FieldPassword, person.FieldTaxID, person.FieldFirstName, person.FieldLastName, person.FieldLanguage, person.FieldGender, person.FieldAddress, person.FieldPostalCode, person.FieldCity, person.FieldCountry:
			values[i] = new(sql.NullString)
		case person.FieldEmailVerifiedAt, person.FieldBirthdate, person.FieldCreatedAt, person.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case person.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Person fields.
func (pe *Person) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case person.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pe.ID = *value
			}
		case person.FieldStripeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stripe_id", values[i])
			} else if value.Valid {
				pe.StripeID = new(string)
				*pe.StripeID = value.String
			}
		case person.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pe.Email = value.String
			}
		case person.FieldEmailVerifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified_at", values[i])
			} else if value.Valid {
				pe.EmailVerifiedAt = new(time.Time)
				*pe.EmailVerifiedAt = value.Time
			}
		case person.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				pe.Phone = new(string)
				*pe.Phone = value.String
			}
		case person.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				pe.Password = value.String
			}
		case person.FieldTaxID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tax_id", values[i])
			} else if value.Valid {
				pe.TaxID = value.String
			}
		case person.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				pe.FirstName = value.String
			}
		case person.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				pe.LastName = new(string)
				*pe.LastName = value.String
			}
		case person.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				pe.Language = value.String
			}
		case person.FieldBirthdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthdate", values[i])
			} else if value.Valid {
				pe.Birthdate = new(time.Time)
				*pe.Birthdate = value.Time
			}
		case person.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pe.Gender = new(person.Gender)
				*pe.Gender = person.Gender(value.String)
			}
		case person.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				pe.Address = new(string)
				*pe.Address = value.String
			}
		case person.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				pe.PostalCode = new(string)
				*pe.PostalCode = value.String
			}
		case person.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				pe.City = new(string)
				*pe.City = value.String
			}
		case person.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				pe.Country = new(string)
				*pe.Country = value.String
			}
		case person.FieldSubscribed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field subscribed", values[i])
			} else if value.Valid {
				pe.Subscribed = value.Bool
			}
		case person.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Time
			}
		case person.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Time
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Person.
// This includes values selected through modifiers, order, etc.
func (pe *Person) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryAuthentications queries the "authentications" edge of the Person entity.
func (pe *Person) QueryAuthentications() *AuthenticationQuery {
	return NewPersonClient(pe.config).QueryAuthentications(pe)
}

// QueryAuthorizations queries the "authorizations" edge of the Person entity.
func (pe *Person) QueryAuthorizations() *AuthorizationQuery {
	return NewPersonClient(pe.config).QueryAuthorizations(pe)
}

// Update returns a builder for updating this Person.
// Note that you need to call Person.Unwrap() before calling this method if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return NewPersonClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Person entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("data: Person is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	if v := pe.StripeID; v != nil {
		builder.WriteString("stripe_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(pe.Email)
	builder.WriteString(", ")
	if v := pe.EmailVerifiedAt; v != nil {
		builder.WriteString("email_verified_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pe.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("tax_id=")
	builder.WriteString(pe.TaxID)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(pe.FirstName)
	builder.WriteString(", ")
	if v := pe.LastName; v != nil {
		builder.WriteString("last_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(pe.Language)
	builder.WriteString(", ")
	if v := pe.Birthdate; v != nil {
		builder.WriteString("birthdate=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pe.Gender; v != nil {
		builder.WriteString("gender=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := pe.Address; v != nil {
		builder.WriteString("address=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pe.PostalCode; v != nil {
		builder.WriteString("postal_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pe.City; v != nil {
		builder.WriteString("city=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pe.Country; v != nil {
		builder.WriteString("country=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("subscribed=")
	builder.WriteString(fmt.Sprintf("%v", pe.Subscribed))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

func (pe *Person) FullName() string {
	parts := []string{
		pe.FirstName,
	}

	if v := pe.LastName; v != nil {
		parts = append(parts, *v)
	}

	return strings.Join(parts, " ")
}

func (pe *Person) CanSubscribe() bool {
	return pe.Phone != nil &&
		pe.Birthdate != nil &&
		pe.Address != nil &&
		pe.PostalCode != nil &&
		pe.City != nil &&
		pe.Country != nil
}

// NamedAuthentications returns the Authentications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedAuthentications(name string) ([]*Authentication, error) {
	if pe.Edges.namedAuthentications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedAuthentications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedAuthentications(name string, edges ...*Authentication) {
	if pe.Edges.namedAuthentications == nil {
		pe.Edges.namedAuthentications = make(map[string][]*Authentication)
	}
	if len(edges) == 0 {
		pe.Edges.namedAuthentications[name] = []*Authentication{}
	} else {
		pe.Edges.namedAuthentications[name] = append(pe.Edges.namedAuthentications[name], edges...)
	}
}

// NamedAuthorizations returns the Authorizations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Person) NamedAuthorizations(name string) ([]*Authorization, error) {
	if pe.Edges.namedAuthorizations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedAuthorizations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Person) appendNamedAuthorizations(name string, edges ...*Authorization) {
	if pe.Edges.namedAuthorizations == nil {
		pe.Edges.namedAuthorizations = make(map[string][]*Authorization)
	}
	if len(edges) == 0 {
		pe.Edges.namedAuthorizations[name] = []*Authorization{}
	} else {
		pe.Edges.namedAuthorizations[name] = append(pe.Edges.namedAuthorizations[name], edges...)
	}
}

// Persons is a parsable slice of Person.
type Persons []*Person
