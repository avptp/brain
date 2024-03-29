// Code generated by ent, DO NOT EDIT.

package person

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the person type in the database.
	Label = "person"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStripeID holds the string denoting the stripe_id field in the database.
	FieldStripeID = "stripe_id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldEmailVerifiedAt holds the string denoting the email_verified_at field in the database.
	FieldEmailVerifiedAt = "email_verified_at"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldTaxID holds the string denoting the tax_id field in the database.
	FieldTaxID = "tax_id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldBirthdate holds the string denoting the birthdate field in the database.
	FieldBirthdate = "birthdate"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldSubscribed holds the string denoting the subscribed field in the database.
	FieldSubscribed = "subscribed"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAuthentications holds the string denoting the authentications edge name in mutations.
	EdgeAuthentications = "authentications"
	// EdgeAuthorizations holds the string denoting the authorizations edge name in mutations.
	EdgeAuthorizations = "authorizations"
	// Table holds the table name of the person in the database.
	Table = "persons"
	// AuthenticationsTable is the table that holds the authentications relation/edge.
	AuthenticationsTable = "authentications"
	// AuthenticationsInverseTable is the table name for the Authentication entity.
	// It exists in this package in order to avoid circular dependency with the "authentication" package.
	AuthenticationsInverseTable = "authentications"
	// AuthenticationsColumn is the table column denoting the authentications relation/edge.
	AuthenticationsColumn = "person_id"
	// AuthorizationsTable is the table that holds the authorizations relation/edge.
	AuthorizationsTable = "authorizations"
	// AuthorizationsInverseTable is the table name for the Authorization entity.
	// It exists in this package in order to avoid circular dependency with the "authorization" package.
	AuthorizationsInverseTable = "authorizations"
	// AuthorizationsColumn is the table column denoting the authorizations relation/edge.
	AuthorizationsColumn = "person_id"
)

// Columns holds all SQL columns for person fields.
var Columns = []string{
	FieldID,
	FieldStripeID,
	FieldEmail,
	FieldEmailVerifiedAt,
	FieldPhone,
	FieldPassword,
	FieldTaxID,
	FieldFirstName,
	FieldLastName,
	FieldLanguage,
	FieldBirthdate,
	FieldGender,
	FieldAddress,
	FieldPostalCode,
	FieldCity,
	FieldCountry,
	FieldSubscribed,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/avptp/brain/internal/generated/data/runtime"
var (
	Hooks  [5]ent.Hook
	Policy ent.Policy
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// TaxIDValidator is a validator for the "tax_id" field. It is called by the builders before save.
	TaxIDValidator func(string) error
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
	// DefaultSubscribed holds the default value on creation for the "subscribed" field.
	DefaultSubscribed bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderWoman     Gender = "woman"
	GenderMan       Gender = "man"
	GenderNonBinary Gender = "nonbinary"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderWoman, GenderMan, GenderNonBinary:
		return nil
	default:
		return fmt.Errorf("person: invalid enum value for gender field: %q", ge)
	}
}

// OrderOption defines the ordering options for the Person queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStripeID orders the results by the stripe_id field.
func ByStripeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStripeID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByEmailVerifiedAt orders the results by the email_verified_at field.
func ByEmailVerifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailVerifiedAt, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByTaxID orders the results by the tax_id field.
func ByTaxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaxID, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByBirthdate orders the results by the birthdate field.
func ByBirthdate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthdate, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByPostalCode orders the results by the postal_code field.
func ByPostalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCode, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// BySubscribed orders the results by the subscribed field.
func BySubscribed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscribed, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByAuthenticationsCount orders the results by authentications count.
func ByAuthenticationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthenticationsStep(), opts...)
	}
}

// ByAuthentications orders the results by authentications terms.
func ByAuthentications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthenticationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAuthorizationsCount orders the results by authorizations count.
func ByAuthorizationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorizationsStep(), opts...)
	}
}

// ByAuthorizations orders the results by authorizations terms.
func ByAuthorizations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorizationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthenticationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthenticationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthenticationsTable, AuthenticationsColumn),
	)
}
func newAuthorizationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorizationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthorizationsTable, AuthorizationsColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Gender) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Gender) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Gender(str)
	if err := GenderValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}
