package resolvers_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/alexedwards/argon2id"
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/avptp/brain/internal/messaging/templates"
	"github.com/go-redis/redis_rate/v10"
	"github.com/stretchr/testify/mock"
)

func (t *TestSuite) TestAuthorization() {
	const emailCreateMutation = `
		mutation CreateEmailAuthorization(
            $personId: ID!
        ) {
            createEmailAuthorization(input: {
                personId: $personId
            }) {
				authorization {
					id
				}
			}
        }
	`

	type emailCreate struct {
		CreateEmailAuthorization struct {
			Authorization struct {
				ID string
			}
		}
	}

	t.Run("email_create", func() {
		authenticated, p, _, _ := t.authenticate()

		t.messenger.On(
			"Send",
			mock.IsType(&templates.Verification{}),
			mock.IsType(&data.Person{}),
		).Return(nil).Once()

		var response emailCreate
		err := t.api.Post(
			emailCreateMutation,
			&response,
			authenticated,
			client.Var("personId", p.ID.String()),
		)

		t.NoError(err)

		t.messenger.AssertExpectations(t.T())

		a, err := t.data.Authorization.
			Query().
			Where(
				authorization.PersonIDEQ(p.ID),
				authorization.KindEQ(authorization.KindEmail),
			).
			First(t.allowCtx)

		t.NoError(err)
		t.Equal(a.ID, t.parseID(response.CreateEmailAuthorization.Authorization.ID))
	})

	t.Run("email_create_when_already_verified", func() {
		authenticated, p, _, _ := t.authenticate()

		// It is not done with "authenticateWith" because if the email is set,
		// "email_verified_at" field is automatically set to nil.
		err := p.
			Update().
			SetEmailVerifiedAt(time.Now()).
			Exec(t.allowCtx)

		t.NoError(err)

		var response emailCreate
		err = t.api.Post(
			emailCreateMutation,
			&response,
			authenticated,
			client.Var("personId", p.ID.String()),
		)

		t.ErrorContains(err, reporting.ErrConstraint.Message)

		t.messenger.AssertNotCalled(t.T(), "Send")

		exist, err := t.data.Authorization.
			Query().
			Where(
				authorization.PersonIDEQ(p.ID),
				authorization.KindEQ(authorization.KindEmail),
			).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exist)
	})

	t.Run("email_create_with_too_many_attempts", func() {
		authenticated, p, _, _ := t.authenticate()

		rlKey := fmt.Sprintf(
			"createEmailAuthorization:%s",
			strings.ToLower(p.Email),
		)

		res, err := t.limiter.AllowN(
			t.allowCtx,
			rlKey,
			redis_rate.PerHour(t.cfg.AuthorizationEmailRateLimit),
			t.cfg.AuthorizationEmailRateLimit,
		)

		t.NoError(err)
		t.LessOrEqual(res.Remaining, 0)

		var response emailCreate
		err = t.api.Post(
			emailCreateMutation,
			&response,
			authenticated,
			client.Var("personId", p.ID.String()),
		)

		t.ErrorContains(err, reporting.ErrRateLimit.Message)

		t.messenger.AssertNotCalled(t.T(), "Send")

		exist, err := t.data.Authorization.
			Query().
			Where(
				authorization.PersonIDEQ(p.ID),
				authorization.KindEQ(authorization.KindEmail),
			).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exist)
	})

	t.Run("email_create_with_nonexistent_person", func() {
		authenticated, _, _, _ := t.authenticate()

		var response emailCreate
		err := t.api.Post(
			emailCreateMutation,
			&response,
			authenticated,
			client.Var("personId", zeroID),
		)

		t.ErrorContains(err, reporting.ErrNotFound.Message)

		t.messenger.AssertNotCalled(t.T(), "Send")
	})

	t.Run("email_create_without_authorization", func() {
		authenticated, _, _, _ := t.authenticate()
		p := t.factory.Person().Create(t.allowCtx)

		var response emailCreate
		err := t.api.Post(
			emailCreateMutation,
			&response,
			authenticated,
			client.Var("personId", p.ID.String()),
		)

		t.ErrorContains(err, reporting.ErrUnauthorized.Message)

		t.messenger.AssertNotCalled(t.T(), "Send")
	})

	t.Run("email_create_without_authentication", func() {
		var response emailCreate
		err := t.api.Post(
			emailCreateMutation,
			&response,
			client.Var("personId", zeroID),
		)

		t.ErrorContains(err, reporting.ErrUnauthenticated.Message)

		t.messenger.AssertNotCalled(t.T(), "Send")
	})

	const emailApplyMutation = `
		mutation ApplyEmailAuthorization(
            $token: String!
        ) {
            applyEmailAuthorization(input: {
                token: $token
            }) {
				authorizationId
			}
        }
	`

	type emailApply struct {
		ApplyEmailAuthorization struct {
			AuthorizationID string
		}
	}

	t.Run("email_apply", func() {
		a := t.factory.
			Authorization().
			With(func(ac *data.AuthorizationCreate) {
				ac.SetKind(authorization.KindEmail)
			}).
			Create(t.allowCtx)

		p, err := a.Person(t.allowCtx)

		t.NoError(err)
		t.Nil(p.EmailVerifiedAt)

		var response emailApply
		err = t.api.Post(
			emailApplyMutation,
			&response,
			client.Var("token", a.TokenEncoded()),
		)

		t.NoError(err)
		t.Equal(a.ID, t.parseID(response.ApplyEmailAuthorization.AuthorizationID))

		exist, err := t.data.Person.
			Query().
			Where(
				person.IDEQ(p.ID),
				person.EmailVerifiedAtNotNil(),
			).
			Exist(t.allowCtx)

		t.NoError(err)
		t.True(exist)

		exist, err = t.data.Authorization.
			Query().
			Where(authorization.IDEQ(a.ID)).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exist)
	})

	t.Run("email_apply_with_expired_token", func() {
		a := t.factory.
			Authorization().
			With(func(ac *data.AuthorizationCreate) {
				ac.SetKind(authorization.KindEmail)
				ac.SetCreatedAt(
					time.Now().Add(-t.ctn.GetConfig().AuthorizationMaxAge),
				)
			}).
			Create(t.allowCtx)

		p, err := a.Person(t.allowCtx)

		t.NoError(err)
		t.Nil(p.EmailVerifiedAt)

		var response emailApply
		err = t.api.Post(
			emailApplyMutation,
			&response,
			client.Var("token", a.TokenEncoded()),
		)

		t.ErrorContains(err, reporting.ErrNotFound.Message)

		exist, err := t.data.Person.
			Query().
			Where(
				person.IDEQ(p.ID),
				person.EmailVerifiedAtIsNil(),
			).
			Exist(t.allowCtx)

		t.NoError(err)
		t.True(exist)
	})

	t.Run("email_apply_with_nonexistent_token", func() {
		var response emailApply
		err := t.api.Post(
			emailApplyMutation,
			&response,
			client.Var("token", ""),
		)

		t.ErrorContains(err, reporting.ErrNotFound.Message)
	})

	const passwordCreateMutation = `
		mutation CreatePasswordAuthorization(
            $email: String!,
            $captcha: String!
        ) {
            createPasswordAuthorization(input: {
                email: $email
				captcha: $captcha
            }) {
				success
			}
        }
	`

	type passwordCreate struct {
		CreatePasswordAuthorization struct {
			Success bool
		}
	}

	t.Run("password_create", func() {
		_, p, _, _ := t.authenticate()

		t.captcha.On("Verify", "").Return(true).Once()

		t.messenger.On(
			"Send",
			mock.IsType(&templates.Recovery{}),
			mock.IsType(&data.Person{}),
		).Return(nil).Once()

		var response passwordCreate
		err := t.api.Post(
			passwordCreateMutation,
			&response,
			client.Var("email", p.Email),
			client.Var("captcha", ""),
		)

		t.NoError(err)
		t.True(response.CreatePasswordAuthorization.Success)

		t.captcha.AssertExpectations(t.T())
		t.messenger.AssertExpectations(t.T())

		exist, err := t.data.Authorization.
			Query().
			Where(
				authorization.PersonIDEQ(p.ID),
				authorization.KindEQ(authorization.KindPassword),
			).
			Exist(t.allowCtx)

		t.NoError(err)
		t.True(exist)
	})

	t.Run("password_create_with_wrong_captcha", func() {
		_, p, _, _ := t.authenticate()

		t.captcha.On("Verify", "").Return(false).Once()

		var response passwordCreate
		err := t.api.Post(
			passwordCreateMutation,
			&response,
			client.Var("email", p.Email),
			client.Var("captcha", ""),
		)

		t.ErrorContains(err, reporting.ErrCaptcha.Message)

		t.captcha.AssertExpectations(t.T())
		t.messenger.AssertNotCalled(t.T(), "Send")

		exist, err := t.data.Authorization.
			Query().
			Where(
				authorization.PersonIDEQ(p.ID),
				authorization.KindEQ(authorization.KindPassword),
			).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exist)
	})

	t.Run("password_create_with_too_many_attempts", func() {
		_, p, _, _ := t.authenticate()

		rlKey := fmt.Sprintf(
			"createPasswordAuthorization:%s",
			strings.ToLower(p.Email),
		)

		res, err := t.limiter.AllowN(
			t.allowCtx,
			rlKey,
			redis_rate.PerHour(t.cfg.AuthorizationPasswordRateLimit),
			t.cfg.AuthorizationPasswordRateLimit,
		)

		t.NoError(err)
		t.LessOrEqual(res.Remaining, 0)

		t.captcha.On("Verify", "").Return(true).Once()

		var response passwordCreate
		err = t.api.Post(
			passwordCreateMutation,
			&response,
			client.Var("email", p.Email),
			client.Var("captcha", ""),
		)

		t.ErrorContains(err, reporting.ErrRateLimit.Message)

		t.captcha.AssertExpectations(t.T())
		t.messenger.AssertNotCalled(t.T(), "Send")

		exist, err := t.data.Authorization.
			Query().
			Where(
				authorization.PersonIDEQ(p.ID),
				authorization.KindEQ(authorization.KindPassword),
			).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exist)
	})

	t.Run("password_create_with_nonexistent_email", func() {
		input := t.factory.Person().Fields

		t.captcha.On("Verify", "").Return(true).Once()

		var response passwordCreate
		err := t.api.Post(
			passwordCreateMutation,
			&response,
			client.Var("email", input.Email),
			client.Var("captcha", ""),
		)

		t.NoError(err)
		t.True(response.CreatePasswordAuthorization.Success)

		t.captcha.AssertExpectations(t.T())
		t.messenger.AssertNotCalled(t.T(), "Send")
	})

	const passwordApplyMutation = `
		mutation ApplyPasswordAuthorization(
            $token: String!
			$newPassword: String!
        ) {
            applyPasswordAuthorization(input: {
                token: $token
				newPassword: $newPassword
            }) {
				authorizationId
			}
        }
	`

	type passwordApply struct {
		ApplyPasswordAuthorization struct {
			AuthorizationID string
		}
	}

	t.Run("password_apply", func() {
		a := t.factory.
			Authorization().
			With(func(ac *data.AuthorizationCreate) {
				ac.SetKind(authorization.KindPassword)
			}).
			Create(t.allowCtx)

		input := t.factory.Person().Fields

		var response passwordApply
		err := t.api.Post(
			passwordApplyMutation,
			&response,
			client.Var("token", a.TokenEncoded()),
			client.Var("newPassword", input.Password),
		)

		t.NoError(err)
		t.Equal(a.ID, t.parseID(response.ApplyPasswordAuthorization.AuthorizationID))

		p, err := t.data.Person.
			Query().
			Where(person.IDEQ(a.PersonID)).
			First(t.allowCtx)

		t.NoError(err)

		match, err := argon2id.ComparePasswordAndHash(input.Password, p.Password)
		t.NoError(err)
		t.True(match)

		exist, err := t.data.Authorization.
			Query().
			Where(authorization.IDEQ(a.ID)).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exist)
	})

	t.Run("password_apply_with_expired_token", func() {
		a := t.factory.
			Authorization().
			With(func(ac *data.AuthorizationCreate) {
				ac.SetKind(authorization.KindPassword)
				ac.SetCreatedAt(
					time.Now().Add(-t.ctn.GetConfig().AuthorizationMaxAge),
				)
			}).
			Create(t.allowCtx)

		input := t.factory.Person().Fields

		var response passwordApply
		err := t.api.Post(
			passwordApplyMutation,
			&response,
			client.Var("token", a.TokenEncoded()),
			client.Var("newPassword", input.Password),
		)

		t.ErrorContains(err, reporting.ErrNotFound.Message)

		p, err := t.data.Person.
			Query().
			Where(person.IDEQ(a.PersonID)).
			First(t.allowCtx)

		t.NoError(err)

		match, err := argon2id.ComparePasswordAndHash(input.Password, p.Password)
		t.NoError(err)
		t.False(match)
	})

	t.Run("password_apply_with_nonexistent_token", func() {
		input := t.factory.Person().Fields

		var response passwordApply
		err := t.api.Post(
			passwordApplyMutation,
			&response,
			client.Var("token", ""),
			client.Var("newPassword", input.Password),
		)

		t.ErrorContains(err, reporting.ErrNotFound.Message)
	})
}
