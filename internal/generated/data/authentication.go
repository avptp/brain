// Code generated by ent, DO NOT EDIT.

package data

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/avptp/brain/internal/api/types"
	"github.com/avptp/brain/internal/encoding"
	"github.com/avptp/brain/internal/generated/data/authentication"
	"github.com/avptp/brain/internal/generated/data/person"
)

// Authentication is the model entity for the Authentication schema.
type Authentication struct {
	config `faker:"-" json:"-"`
	// ID of the ent.
	ID types.ID `json:"id,omitempty"`
	// PersonID holds the value of the "person_id" field.
	PersonID types.ID `json:"person_id,omitempty"`
	// Token holds the value of the "token" field.
	Token []byte `json:"token,omitempty" faker:"slice_len=64"`
	// CreatedIP holds the value of the "created_ip" field.
	CreatedIP string `json:"created_ip,omitempty" faker:"ipv6"`
	// LastUsedIP holds the value of the "last_used_ip" field.
	LastUsedIP string `json:"last_used_ip,omitempty" faker:"ipv6"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// LastUsedAt holds the value of the "last_used_at" field.
	LastUsedAt time.Time `json:"last_used_at,omitempty"`
	// LastPasswordChallengeAt holds the value of the "last_password_challenge_at" field.
	LastPasswordChallengeAt *time.Time `json:"last_password_challenge_at,omitempty" faker:"-"`
	// LastCaptchaChallengeAt holds the value of the "last_captcha_challenge_at" field.
	LastCaptchaChallengeAt *time.Time `json:"last_captcha_challenge_at,omitempty" faker:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthenticationQuery when eager-loading is set.
	Edges        AuthenticationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AuthenticationEdges holds the relations/edges for other nodes in the graph.
type AuthenticationEdges struct {
	// Person holds the value of the person edge.
	Person *Person `json:"person,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// PersonOrErr returns the Person value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthenticationEdges) PersonOrErr() (*Person, error) {
	if e.Person != nil {
		return e.Person, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "person"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Authentication) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authentication.FieldToken:
			values[i] = new([]byte)
		case authentication.FieldCreatedIP, authentication.FieldLastUsedIP:
			values[i] = new(sql.NullString)
		case authentication.FieldCreatedAt, authentication.FieldLastUsedAt, authentication.FieldLastPasswordChallengeAt, authentication.FieldLastCaptchaChallengeAt:
			values[i] = new(sql.NullTime)
		case authentication.FieldID, authentication.FieldPersonID:
			values[i] = new(types.ID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Authentication fields.
func (a *Authentication) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authentication.FieldID:
			if value, ok := values[i].(*types.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case authentication.FieldPersonID:
			if value, ok := values[i].(*types.ID); !ok {
				return fmt.Errorf("unexpected type %T for field person_id", values[i])
			} else if value != nil {
				a.PersonID = *value
			}
		case authentication.FieldToken:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value != nil {
				a.Token = *value
			}
		case authentication.FieldCreatedIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_ip", values[i])
			} else if value.Valid {
				a.CreatedIP = value.String
			}
		case authentication.FieldLastUsedIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_used_ip", values[i])
			} else if value.Valid {
				a.LastUsedIP = value.String
			}
		case authentication.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case authentication.FieldLastUsedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used_at", values[i])
			} else if value.Valid {
				a.LastUsedAt = value.Time
			}
		case authentication.FieldLastPasswordChallengeAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_password_challenge_at", values[i])
			} else if value.Valid {
				a.LastPasswordChallengeAt = new(time.Time)
				*a.LastPasswordChallengeAt = value.Time
			}
		case authentication.FieldLastCaptchaChallengeAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_captcha_challenge_at", values[i])
			} else if value.Valid {
				a.LastCaptchaChallengeAt = new(time.Time)
				*a.LastCaptchaChallengeAt = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Authentication.
// This includes values selected through modifiers, order, etc.
func (a *Authentication) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryPerson queries the "person" edge of the Authentication entity.
func (a *Authentication) QueryPerson() *PersonQuery {
	return NewAuthenticationClient(a.config).QueryPerson(a)
}

// Update returns a builder for updating this Authentication.
// Note that you need to call Authentication.Unwrap() before calling this method if this Authentication
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Authentication) Update() *AuthenticationUpdateOne {
	return NewAuthenticationClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Authentication entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Authentication) Unwrap() *Authentication {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("data: Authentication is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Authentication) String() string {
	var builder strings.Builder
	builder.WriteString("Authentication(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("person_id=")
	builder.WriteString(fmt.Sprintf("%v", a.PersonID))
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(fmt.Sprintf("%v", a.Token))
	builder.WriteString(", ")
	builder.WriteString("created_ip=")
	builder.WriteString(a.CreatedIP)
	builder.WriteString(", ")
	builder.WriteString("last_used_ip=")
	builder.WriteString(a.LastUsedIP)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_used_at=")
	builder.WriteString(a.LastUsedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := a.LastPasswordChallengeAt; v != nil {
		builder.WriteString("last_password_challenge_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := a.LastCaptchaChallengeAt; v != nil {
		builder.WriteString("last_captcha_challenge_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

func (a *Authentication) TokenEncoded() string {
	return encoding.Base32.EncodeToString(a.Token)
}

// Authentications is a parsable slice of Authentication.
type Authentications []*Authentication
