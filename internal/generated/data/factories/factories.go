// Code generated by ent, DO NOT EDIT.

package factories

import (
	"context"
	"time"

	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/google/uuid"

	"github.com/brianvoe/gofakeit/v6"
)

// Base factory
type Factory struct {
	data *data.Client
}

func New(data *data.Client) *Factory {
	return &Factory{
		data: data,
	}
}

// Authentication factory
type AuthenticationFactory struct {
	*Factory

	Fields  AuthenticationFields
	builder *data.AuthenticationCreate
}

type AuthenticationFields struct {
	PersonID   uuid.UUID `json:"person_id,omitempty"`
	Token      []byte    `json:"token,omitempty" fakesize:"64"`
	CreatedIP  string    `json:"created_ip,omitempty" fake:"{ipv6address}"`
	LastUsedIP string    `json:"last_used_ip,omitempty" fake:"{ipv6address}"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	LastUsedAt time.Time `json:"last_used_at,omitempty"`
}

func (bf *Factory) Authentication() *AuthenticationFactory {
	f := &AuthenticationFactory{
		Factory: bf,
	}

	gofakeit.Struct(&f.Fields)

	f.builder = f.data.Authentication.
		Create().
		SetToken(f.Fields.Token).
		SetCreatedIP(f.Fields.CreatedIP).
		SetLastUsedIP(f.Fields.LastUsedIP)

	return f
}

func (f *AuthenticationFactory) With(cb func(*data.AuthenticationCreate)) *AuthenticationFactory {
	cb(f.builder)

	return f
}

func (f *AuthenticationFactory) Create(ctx context.Context) *data.Authentication {
	if _, exists := f.builder.Mutation().PersonID(); !exists {
		f.builder.SetPerson(
			f.Person().Create(ctx),
		)
	}

	return f.builder.SaveX(ctx)
}

// Authorization factory
type AuthorizationFactory struct {
	*Factory

	Fields  AuthorizationFields
	builder *data.AuthorizationCreate
}

type AuthorizationFields struct {
	PersonID  uuid.UUID          `json:"person_id,omitempty"`
	Token     []byte             `json:"token,omitempty" fakesize:"64"`
	Kind      authorization.Kind `json:"kind,omitempty" fake:"{randomstring:[email,password]}"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
}

func (bf *Factory) Authorization() *AuthorizationFactory {
	f := &AuthorizationFactory{
		Factory: bf,
	}

	gofakeit.Struct(&f.Fields)

	f.builder = f.data.Authorization.
		Create().
		SetToken(f.Fields.Token).
		SetKind(f.Fields.Kind)

	return f
}

func (f *AuthorizationFactory) With(cb func(*data.AuthorizationCreate)) *AuthorizationFactory {
	cb(f.builder)

	return f
}

func (f *AuthorizationFactory) Create(ctx context.Context) *data.Authorization {
	if _, exists := f.builder.Mutation().PersonID(); !exists {
		f.builder.SetPerson(
			f.Person().Create(ctx),
		)
	}

	return f.builder.SaveX(ctx)
}

// Person factory
type PersonFactory struct {
	*Factory

	Fields  PersonFields
	builder *data.PersonCreate
}

type PersonFields struct {
	StripeID        *string        `json:"stripe_id,omitempty"`
	Email           string         `json:"email,omitempty" fake:"{email}"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at,omitempty"`
	Phone           *string        `json:"phone,omitempty" fake:"{phone_e164}"`
	Password        string         `json:"password,omitempty"`
	TaxID           string         `json:"tax_id,omitempty" fake:"{tax_id}"`
	FirstName       string         `json:"first_name,omitempty" fake:"{firstname}"`
	LastName        *string        `json:"last_name,omitempty" fake:"{lastname}"`
	Language        string         `json:"language,omitempty" fake:"{randomstring:[ca,es,en]}"`
	Birthdate       *time.Time     `json:"birthdate,omitempty" fake:"{date}"`
	Gender          *person.Gender `json:"gender,omitempty" fake:"{randomstring:[woman,man,nonbinary]}"`
	Address         *string        `json:"address,omitempty" fake:"{street}"`
	PostalCode      *string        `json:"postal_code,omitempty" fake:"{zip}"`
	City            *string        `json:"city,omitempty" fake:"{city}"`
	Country         *string        `json:"country,omitempty" fake:"{countryabr}"`
	Subscribed      bool           `json:"subscribed,omitempty"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
}

func (bf *Factory) Person() *PersonFactory {
	f := &PersonFactory{
		Factory: bf,
	}

	gofakeit.Struct(&f.Fields)

	f.builder = f.data.Person.
		Create().
		SetNillableStripeID(f.Fields.StripeID).
		SetEmail(f.Fields.Email).
		SetNillableEmailVerifiedAt(f.Fields.EmailVerifiedAt).
		SetNillablePhone(f.Fields.Phone).
		SetPassword(f.Fields.Password).
		SetTaxID(f.Fields.TaxID).
		SetFirstName(f.Fields.FirstName).
		SetNillableLastName(f.Fields.LastName).
		SetLanguage(f.Fields.Language).
		SetNillableBirthdate(f.Fields.Birthdate).
		SetNillableGender(f.Fields.Gender).
		SetNillableAddress(f.Fields.Address).
		SetNillablePostalCode(f.Fields.PostalCode).
		SetNillableCity(f.Fields.City).
		SetNillableCountry(f.Fields.Country)

	return f
}

func (f *PersonFactory) With(cb func(*data.PersonCreate)) *PersonFactory {
	cb(f.builder)

	return f
}

func (f *PersonFactory) Create(ctx context.Context) *data.Person {
	return f.builder.SaveX(ctx)
}
