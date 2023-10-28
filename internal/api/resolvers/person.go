package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/alexedwards/argon2id"
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/api"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/avptp/brain/internal/generated/data/privacy"
	"github.com/avptp/brain/internal/messaging/templates"
	"github.com/avptp/brain/internal/transport/request"
	"github.com/google/uuid"
)

// CreatePerson is the resolver for the createPerson field.
func (r *mutationResolver) CreatePerson(ctx context.Context, input api.CreatePersonInput) (*api.CreatePersonPayload, error) {
	if !r.captcha.Verify(input.Captcha) {
		return nil, reporting.ErrCaptcha
	}

	d := data.FromContext(ctx) // transactional data client for mutations
	allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

	// Create person
	create := d.Person.
		Create().
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetTaxID(input.TaxID).
		SetFirstName(input.FirstName).
		SetLanguage(input.Language)

	if input.LastName.IsSet() {
		create.SetNillableLastName(input.LastName.Value())
	}

	p, err := create.Save(allowCtx)

	if err != nil {
		return nil, err
	}

	// Create authorization
	a, err := d.Authorization.
		Create().
		SetPersonID(p.ID).
		SetKind(authorization.KindEmail).
		Save(allowCtx)

	if err != nil {
		return nil, err
	}

	// Send welcome message
	err = r.messenger.Send(&templates.Welcome{
		Link: fmt.Sprintf(
			"%s/%s/%s",
			r.cfg.FrontUrl,
			r.cfg.FrontEmailAuthorizationPath,
			a.TokenEncoded(),
		),
		Validity: fmt.Sprintf("%d", r.cfg.AuthorizationMaxAge/time.Hour),
	}, p)

	if err != nil {
		return nil, err
	}

	// Return payload
	return &api.CreatePersonPayload{
		Person: p,
	}, nil
}

// UpdatePerson is the resolver for the updatePerson field.
func (r *mutationResolver) UpdatePerson(ctx context.Context, input api.UpdatePersonInput) (*api.UpdatePersonPayload, error) {
	d := data.FromContext(ctx) // transactional data client for mutations

	// Update person
	update := d.Person.UpdateOneID(input.ID)

	if input.Email.IsSet() {
		if v := input.Email.Value(); v != nil {
			update.SetEmail(*v)
		}
	}

	if input.Phone.IsSet() {
		if v := input.Phone.Value(); v != nil {
			update.SetPhone(*v)
		} else {
			update.ClearPhone()
		}
	}

	if input.TaxID.IsSet() {
		if v := input.TaxID.Value(); v != nil {
			update.SetTaxID(*v)
		}
	}

	if input.FirstName.IsSet() {
		if v := input.FirstName.Value(); v != nil {
			update.SetFirstName(*v)
		}
	}

	if input.LastName.IsSet() {
		if v := input.LastName.Value(); v != nil {
			update.SetLastName(*v)
		} else {
			update.ClearLastName()
		}
	}

	if input.Language.IsSet() {
		if v := input.Language.Value(); v != nil {
			update.SetLanguage(*v)
		}
	}

	if input.Birthdate.IsSet() {
		if v := input.Birthdate.Value(); v != nil {
			update.SetBirthdate(*v)
		} else {
			update.ClearBirthdate()
		}
	}

	if input.Gender.IsSet() {
		if v := input.Gender.Value(); v != nil {
			update.SetGender(*v)
		} else {
			update.ClearGender()
		}
	}

	if input.Address.IsSet() {
		if v := input.Address.Value(); v != nil {
			update.SetAddress(*v)
		} else {
			update.ClearAddress()
		}
	}

	if input.PostalCode.IsSet() {
		if v := input.PostalCode.Value(); v != nil {
			update.SetPostalCode(*v)
		} else {
			update.ClearPostalCode()
		}
	}

	if input.City.IsSet() {
		if v := input.City.Value(); v != nil {
			update.SetCity(*v)
		} else {
			update.ClearCity()
		}
	}

	if input.Country.IsSet() {
		if v := input.Country.Value(); v != nil {
			update.SetCountry(*v)
		} else {
			update.ClearCountry()
		}
	}

	person, err := update.Save(ctx)

	if err != nil {
		return nil, err
	}

	// Sync with biller
	err = r.biller.SyncPerson(person)

	if err != nil {
		return nil, err
	}

	// Return payload
	return &api.UpdatePersonPayload{
		Person: person,
	}, nil
}

// UpdatePersonPassword is the resolver for the updatePersonPassword field.
func (r *mutationResolver) UpdatePersonPassword(ctx context.Context, input api.UpdatePersonPasswordInput) (*api.UpdatePersonPasswordPayload, error) {
	if !r.captcha.Verify(input.Captcha) {
		return nil, reporting.ErrCaptcha
	}

	d := data.FromContext(ctx) // transactional data client for mutations

	person, err := d.Person.
		Query().
		Where(person.IDEQ(input.ID)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	match, err := argon2id.ComparePasswordAndHash(input.CurrentPassword, person.Password)

	if err != nil {
		return nil, err
	}

	if !match {
		return nil, reporting.ErrWrongPassword
	}

	person, err = person.
		Update().
		SetPassword(input.NewPassword).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return &api.UpdatePersonPasswordPayload{
		Person: person,
	}, nil
}

// DeletePerson is the resolver for the deletePerson field.
func (r *mutationResolver) DeletePerson(ctx context.Context, input api.DeletePersonInput) (*api.DeletePersonPayload, error) {
	if !r.captcha.Verify(input.Captcha) {
		return nil, reporting.ErrCaptcha
	}

	d := data.FromContext(ctx) // transactional data client for mutations

	person, err := d.Person.
		Query().
		Where(person.IDEQ(input.ID)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	match, err := argon2id.ComparePasswordAndHash(input.CurrentPassword, person.Password)

	if err != nil {
		return nil, err
	}

	if !match {
		return nil, reporting.ErrWrongPassword
	}

	err = d.Person.
		DeleteOneID(input.ID).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &api.DeletePersonPayload{
		PersonID: input.ID,
	}, nil
}

// Authentications is the resolver for the authentications field.
func (r *personResolver) Authentications(ctx context.Context, obj *data.Person, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int) (*data.AuthenticationConnection, error) {
	return obj.
		QueryAuthentications().
		Paginate(ctx, after, first, before, last)
}

// Viewer is the resolver for the viewer field.
func (r *queryResolver) Viewer(ctx context.Context) (*data.Person, error) {
	viewer := request.ViewerFromCtx(ctx)

	if viewer == nil {
		return nil, reporting.ErrUnauthenticated
	}

	return viewer, nil
}

// Person returns api.PersonResolver implementation.
func (r *Resolver) Person() api.PersonResolver { return &personResolver{r} }

type personResolver struct{ *Resolver }
