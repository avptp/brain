package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/api"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/avptp/brain/internal/generated/data/privacy"
	"github.com/avptp/brain/internal/messaging/templates"
	"github.com/go-redis/redis_rate/v10"
)

// CreateEmailAuthorization is the resolver for the createEmailAuthorization field.
func (r *mutationResolver) CreateEmailAuthorization(ctx context.Context, input api.CreateEmailAuthorizationInput) (*api.CreateEmailAuthorizationPayload, error) {
	d := data.FromContext(ctx) // transactional data client for mutations

	// Retrieve person and check constraints
	p, err := d.Person.
		Query().
		Where(person.IDEQ(input.PersonID)).
		First(ctx)

	if err != nil {
		return nil, reporting.ErrNotFound
	}

	if p.EmailVerifiedAt != nil {
		return nil, reporting.ErrConstraint
	}

	// Rate limit by normalized email
	rlKey := fmt.Sprintf(
		"createEmailAuthorization:%s",
		strings.ToLower(p.Email),
	)

	res, err := r.limiter.Allow(ctx, rlKey, redis_rate.PerHour(r.cfg.AuthorizationEmailRateLimit))

	if err != nil {
		return nil, err
	}

	if res.Allowed <= 0 {
		return nil, reporting.ErrRateLimit
	}

	// Delete existing person's email authorizations
	_, err = d.Authorization.
		Delete().
		Where(
			authorization.PersonIDEQ(input.PersonID),
			authorization.KindEQ(authorization.KindEmail),
		).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	// Create the email authorization
	a, err := d.Authorization.
		Create().
		SetPersonID(input.PersonID).
		SetKind(authorization.KindEmail).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	// Send an email with the token
	err = r.messenger.Send(&templates.Verification{
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

	return &api.CreateEmailAuthorizationPayload{
		Authorization: a,
	}, nil
}

// ApplyEmailAuthorization is the resolver for the applyEmailAuthorization field.
func (r *mutationResolver) ApplyEmailAuthorization(ctx context.Context, input api.ApplyEmailAuthorizationInput) (*api.ApplyEmailAuthorizationPayload, error) {
	d := data.FromContext(ctx) // transactional data client for mutations
	allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

	token, err := base64.URLEncoding.DecodeString(input.Token)

	if err != nil {
		return nil, err
	}

	a, err := d.Authorization.
		Query().
		Where(
			authorization.TokenEQ(token),
			authorization.KindEQ(authorization.KindEmail),
			authorization.CreatedAtGT(
				time.Now().Add(-r.cfg.AuthorizationMaxAge),
			),
		).
		First(allowCtx)

	if err != nil {
		return nil, err
	}

	err = d.Person.
		Update().
		Where(person.IDEQ(a.PersonID)).
		SetEmailVerifiedAt(time.Now()).
		Exec(allowCtx)

	if err != nil {
		return nil, err
	}

	err = d.Authorization.
		DeleteOne(a).
		Exec(allowCtx)

	if err != nil {
		return nil, err
	}

	return &api.ApplyEmailAuthorizationPayload{
		AuthorizationID: a.ID,
	}, nil
}

// CreatePasswordAuthorization is the resolver for the createPasswordAuthorization field.
func (r *mutationResolver) CreatePasswordAuthorization(ctx context.Context, input api.CreatePasswordAuthorizationInput) (*api.CreatePasswordAuthorizationPayload, error) {
	if !r.captcha.Verify(input.Captcha) {
		return nil, reporting.ErrCaptcha
	}

	// Rate limit by normalized email
	// (to avoid exposing whether an email is in use)
	rlKey := fmt.Sprintf(
		"createPasswordAuthorization:%s",
		strings.ToLower(input.Email),
	)

	res, err := r.limiter.Allow(ctx, rlKey, redis_rate.PerHour(r.cfg.AuthorizationPasswordRateLimit))

	if err != nil {
		return nil, err
	}

	if res.Allowed <= 0 {
		return nil, reporting.ErrRateLimit
	}

	// Retrieve person by email
	// (it always returns "success" to avoid exposing whether an email is in use)
	d := data.FromContext(ctx) // transactional data client for mutations
	allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

	p, err := d.Person.
		Query().
		Where(person.EmailEQ(input.Email)).
		First(allowCtx)

	if err != nil {
		return &api.CreatePasswordAuthorizationPayload{
			Success: true,
		}, nil
	}

	// Delete existing person's password authorizations
	_, err = d.Authorization.
		Delete().
		Where(
			authorization.PersonIDEQ(p.ID),
			authorization.KindEQ(authorization.KindPassword),
		).
		Exec(allowCtx)

	if err != nil {
		return nil, err
	}

	// Create the password authorization
	a, err := d.Authorization.
		Create().
		SetPersonID(p.ID).
		SetKind(authorization.KindPassword).
		Save(allowCtx)

	if err != nil {
		return nil, err
	}

	// Send an email with the token
	err = r.messenger.Send(&templates.Recovery{
		Link: fmt.Sprintf(
			"%s/%s/%s",
			r.cfg.FrontUrl,
			r.cfg.FrontPasswordAuthorizationPath,
			a.TokenEncoded(),
		),
		Validity: fmt.Sprintf("%d", r.cfg.AuthorizationMaxAge/time.Hour),
	}, p)

	if err != nil {
		return nil, err
	}

	return &api.CreatePasswordAuthorizationPayload{
		Success: true,
	}, nil
}

// ApplyPasswordAuthorization is the resolver for the applyPasswordAuthorization field.
func (r *mutationResolver) ApplyPasswordAuthorization(ctx context.Context, input api.ApplyPasswordAuthorizationInput) (*api.ApplyPasswordAuthorizationPayload, error) {
	d := data.FromContext(ctx) // transactional data client for mutations
	allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

	token, err := base64.URLEncoding.DecodeString(input.Token)

	if err != nil {
		return nil, err
	}

	a, err := d.Authorization.
		Query().
		Where(
			authorization.TokenEQ(token),
			authorization.KindEQ(authorization.KindPassword),
			authorization.CreatedAtGT(
				time.Now().Add(-r.cfg.AuthorizationMaxAge),
			),
		).
		First(allowCtx)

	if err != nil {
		return nil, err
	}

	err = d.Person.
		Update().
		Where(person.IDEQ(a.PersonID)).
		SetPassword(input.NewPassword).
		Exec(allowCtx)

	if err != nil {
		return nil, err
	}

	err = d.Authorization.
		DeleteOne(a).
		Exec(allowCtx)

	if err != nil {
		return nil, err
	}

	return &api.ApplyPasswordAuthorizationPayload{
		AuthorizationID: a.ID,
	}, nil
}
