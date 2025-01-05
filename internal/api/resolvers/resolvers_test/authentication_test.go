package resolvers_test

import (
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/client"
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/encoding"
	"github.com/avptp/brain/internal/generated/data/authentication"
	"github.com/go-redis/redis_rate/v10"
)

func (t *TestSuite) TestAuthentication() {
	const createMutation = `
		mutation CreateAuthentication(
            $email: String!,
            $password: String!
        ) {
            createAuthentication(input: {
                email: $email
				password: $password
            }) {
				token
			}
        }
	`

	type create struct {
		CreateAuthentication struct {
			Token string
		}
	}

	t.Run("create", func() {
		_, p, pf, _ := t.authenticate()

		var response create
		err := t.api.Post(
			createMutation,
			&response,
			client.Var("email", pf.Email),
			client.Var("password", pf.Password),
		)

		t.NoError(err)

		token, err := encoding.Base32.DecodeString(response.CreateAuthentication.Token)
		t.NoError(err)

		atc, err := t.data.Authentication.
			Query().
			WithPerson().
			Where(authentication.TokenEQ(token)).
			First(t.allowCtx)

		t.NoError(err)
		t.Equal(p.ID, atc.Edges.Person.ID)
	})

	t.Run("create_with_too_many_attempts", func() {
		_, p, pf, _ := t.authenticate()

		rlKey := fmt.Sprintf(
			"createAuthentication:%s",
			strings.ToLower(p.Email),
		)

		res, err := t.limiter.AllowN(
			t.allowCtx,
			rlKey,
			redis_rate.PerHour(t.cfg.AuthenticationPasswordChallengeRateLimit),
			t.cfg.AuthenticationPasswordChallengeRateLimit,
		)

		t.NoError(err)
		t.LessOrEqual(res.Remaining, 0)

		var response create
		err = t.api.Post(
			createMutation,
			&response,
			client.Var("email", pf.Email),
			client.Var("password", pf.Password),
		)

		t.ErrorContains(err, reporting.ErrRateLimit.Message)
	})

	t.Run("create_with_wrong_email", func() {
		input := t.factory.Person().Fields

		var response create
		err := t.api.Post(
			createMutation,
			&response,
			client.Var("email", input.Email),
			client.Var("password", input.Password),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)
	})

	t.Run("create_with_wrong_password", func() {
		_, p, pf, _ := t.authenticate()

		var response create
		err := t.api.Post(
			createMutation,
			&response,
			client.Var("email", p.Email),
			client.Var("password", pf.Password+"wrong"),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)
	})

	const listQuery = `
 		query {
            viewer {
                authentications {
                    edges {
                        node {
                            id
                        }
                    }
                }
            }
        }
	`

	type list struct {
		Viewer struct {
			Authentications struct {
				Edges []struct {
					Node struct {
						ID string
					}
				}
			}
		}
	}

	t.Run("list", func() {
		authenticated, _, _, a := t.authenticate()

		var response list
		err := t.api.Post(listQuery, &response, authenticated)

		t.NoError(err)
		t.Len(response.Viewer.Authentications.Edges, 1)

		id := t.parseID(response.Viewer.Authentications.Edges[0].Node.ID)
		t.Equal(a.ID, id)
	})

	t.Run("list_without_authentication", func() {
		var response list
		err := t.api.Post(listQuery, &response)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthenticated.Message)
	})

	const passPasswordChallengeMutation = `
		mutation PassAuthenticationPasswordChallenge(
            $id: ID!
            $password: String!
        ) {
            passAuthenticationPasswordChallenge(input: {
                id: $id
                password: $password
            }) {
				success
			}
        }
	`

	type passPasswordChallenge struct {
		PassAuthenticationPasswordChallenge struct {
			Success bool
		}
	}

	t.Run("pass_password_challenge", func() {
		authenticated, _, pf, auth := t.authenticate()

		t.Nil(auth.LastPasswordChallengeAt)

		var response passPasswordChallenge
		err := t.api.Post(
			passPasswordChallengeMutation,
			&response,
			authenticated,
			client.Var("id", auth.ID.String()),
			client.Var("password", pf.Password),
		)

		t.NoError(err)
		t.True(response.PassAuthenticationPasswordChallenge.Success)

		auth, err = t.data.Authentication.Get(t.allowCtx, auth.ID)

		t.NoError(err)
		t.NotNil(auth.LastPasswordChallengeAt)
		t.Nil(auth.LastCaptchaChallengeAt)
	})

	t.Run("pass_password_challenge_with_wrong_password", func() {
		authenticated, _, pf, auth := t.authenticate()

		var response passPasswordChallenge
		err := t.api.Post(
			passPasswordChallengeMutation,
			&response,
			authenticated,
			client.Var("id", auth.ID.String()),
			client.Var("password", pf.Password+"wrong"),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)

		auth, err = t.data.Authentication.Get(t.allowCtx, auth.ID)

		t.NoError(err)
		t.Nil(auth.LastPasswordChallengeAt)
	})

	t.Run("pass_password_challenge_with_too_many_attempts", func() {
		authenticated, p, pf, auth := t.authenticate()

		rlKey := fmt.Sprintf(
			"passAuthenticationPasswordChallenge:%s",
			p.ID,
		)

		res, err := t.limiter.AllowN(
			t.allowCtx,
			rlKey,
			redis_rate.PerHour(t.cfg.AuthenticationPasswordChallengeRateLimit),
			t.cfg.AuthenticationPasswordChallengeRateLimit,
		)

		t.NoError(err)
		t.LessOrEqual(res.Remaining, 0)

		var response passPasswordChallenge
		err = t.api.Post(
			passPasswordChallengeMutation,
			&response,
			authenticated,
			client.Var("id", auth.ID.String()),
			client.Var("password", pf.Password),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrRateLimit.Message)

		auth, err = t.data.Authentication.Get(t.allowCtx, auth.ID)

		t.NoError(err)
		t.Nil(auth.LastPasswordChallengeAt)
	})

	t.Run("pass_password_challenge_with_nonexistent_authentication", func() {
		authenticated, _, _, _ := t.authenticate()

		var response passPasswordChallenge
		err := t.api.Post(
			passPasswordChallengeMutation,
			&response,
			authenticated,
			client.Var("id", zeroID),
			client.Var("password", ""),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrNotFound.Message)
	})

	t.Run("pass_password_challenge_without_authorization", func() {
		authenticated, _, _, _ := t.authenticate()
		auth := t.factory.Authentication().Create(t.allowCtx)

		var response passPasswordChallenge
		err := t.api.Post(
			passPasswordChallengeMutation,
			&response,
			authenticated,
			client.Var("id", auth.ID.String()),
			client.Var("password", ""),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthorized.Message)
	})

	t.Run("pass_password_challenge_without_authentication", func() {
		var response passPasswordChallenge
		err := t.api.Post(
			passPasswordChallengeMutation,
			&response,
			client.Var("id", zeroID),
			client.Var("password", ""),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthenticated.Message)
	})

	const passCaptchaChallengeMutation = `
		mutation PassAuthenticationCaptchaChallenge(
            $id: ID!
            $captcha: String!
        ) {
            passAuthenticationCaptchaChallenge(input: {
                id: $id
                captcha: $captcha
            }) {
				success
			}
        }
	`

	type passCaptchaChallenge struct {
		PassAuthenticationCaptchaChallenge struct {
			Success bool
		}
	}

	t.Run("pass_captcha_challenge", func() {
		authenticated, _, _, auth := t.authenticate()

		t.Nil(auth.LastCaptchaChallengeAt)

		t.captcha.On("Verify", "").Return(true).Once()

		var response passCaptchaChallenge
		err := t.api.Post(
			passCaptchaChallengeMutation,
			&response,
			authenticated,
			client.Var("id", auth.ID.String()),
			client.Var("captcha", ""),
		)

		t.captcha.AssertExpectations(t.T())

		t.NoError(err)
		t.True(response.PassAuthenticationCaptchaChallenge.Success)

		auth, err = t.data.Authentication.Get(t.allowCtx, auth.ID)

		t.NoError(err)
		t.NotNil(auth.LastCaptchaChallengeAt)
		t.Nil(auth.LastPasswordChallengeAt)
	})

	t.Run("pass_captcha_challenge_with_wrong_captcha", func() {
		authenticated, _, _, auth := t.authenticate()

		t.captcha.On("Verify", "").Return(false).Once()

		var response passCaptchaChallenge
		err := t.api.Post(
			passCaptchaChallengeMutation,
			&response,
			authenticated,
			client.Var("id", auth.ID.String()),
			client.Var("captcha", ""),
		)

		t.captcha.AssertExpectations(t.T())

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrCaptcha.Message)

		auth, err = t.data.Authentication.Get(t.allowCtx, auth.ID)

		t.NoError(err)
		t.Nil(auth.LastCaptchaChallengeAt)
	})

	t.Run("pass_captcha_challenge_with_too_many_attempts", func() {
		authenticated, p, _, auth := t.authenticate()

		rlKey := fmt.Sprintf(
			"passAuthenticationCaptchaChallenge:%s",
			p.ID,
		)

		res, err := t.limiter.AllowN(
			t.allowCtx,
			rlKey,
			redis_rate.PerHour(t.cfg.AuthenticationCaptchaChallengeRateLimit),
			t.cfg.AuthenticationCaptchaChallengeRateLimit,
		)

		t.NoError(err)
		t.LessOrEqual(res.Remaining, 0)

		var response passCaptchaChallenge
		err = t.api.Post(
			passCaptchaChallengeMutation,
			&response,
			authenticated,
			client.Var("id", auth.ID.String()),
			client.Var("captcha", ""),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrRateLimit.Message)

		auth, err = t.data.Authentication.Get(t.allowCtx, auth.ID)

		t.NoError(err)
		t.Nil(auth.LastCaptchaChallengeAt)
	})

	t.Run("pass_captcha_challenge_with_nonexistent_authentication", func() {
		authenticated, _, _, _ := t.authenticate()

		var response passCaptchaChallenge
		err := t.api.Post(
			passCaptchaChallengeMutation,
			&response,
			authenticated,
			client.Var("id", zeroID),
			client.Var("captcha", ""),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrNotFound.Message)
	})

	t.Run("pass_captcha_challenge_without_authorization", func() {
		authenticated, _, _, _ := t.authenticate()
		auth := t.factory.Authentication().Create(t.allowCtx)

		var response passCaptchaChallenge
		err := t.api.Post(
			passCaptchaChallengeMutation,
			&response,
			authenticated,
			client.Var("id", auth.ID.String()),
			client.Var("captcha", ""),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthorized.Message)
	})

	t.Run("pass_captcha_challenge_without_authentication", func() {
		var response passCaptchaChallenge
		err := t.api.Post(
			passCaptchaChallengeMutation,
			&response,
			client.Var("id", zeroID),
			client.Var("captcha", ""),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthenticated.Message)
	})

	const deleteMutation = `
		mutation(
			$id: ID!,
		) {
			deleteAuthentication(input: {
				id: $id
			}) {
				authenticationId
			}
		}
	`

	type delete struct {
		DeleteAuthentication struct {
			AuthenticationID string
		}
	}

	t.Run("delete", func() {
		authenticated, _, _, a := t.authenticate()

		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", a.ID.String()),
		)

		t.NoError(err)
		t.Equal(a.ID, t.parseID(response.DeleteAuthentication.AuthenticationID))

		exists, err := t.data.Authentication.
			Query().
			Where(authentication.IDEQ(a.ID)).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exists)
	})

	t.Run("delete_when_nonexistent", func() {
		authenticated, _, _, _ := t.authenticate()

		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", zeroID),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrNotFound.Message)
	})

	t.Run("delete_without_authorization", func() {
		authenticated, _, _, _ := t.authenticate()
		a := t.factory.Authentication().Create(t.allowCtx)

		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", a.ID.String()),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthorized.Message)
	})

	t.Run("delete_without_authentication", func() {
		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			client.Var("id", zeroID),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthenticated.Message)
	})

	// Performing this test on a single mutation of a model is enough.
	t.Run("delete_with_invalid_id", func() {
		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			client.Var("id", "i"), // the "i" character is invalid for ID encoding
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrInput.Message)
	})
}
