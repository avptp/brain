package resolvers_test

import (
	"encoding/base64"
	"fmt"

	"github.com/99designs/gqlgen/client"
	"github.com/avptp/brain/internal/api/reporting"
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
		_, p, pf, _, _ := t.authenticate()

		var response create
		err := t.api.Post(
			createMutation,
			&response,
			client.Var("email", pf.Email),
			client.Var("password", pf.Password),
		)

		t.NoError(err)

		token, err := base64.URLEncoding.DecodeString(response.CreateAuthentication.Token)
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
		_, p, pf, _, _ := t.authenticate()

		rlKey := fmt.Sprintf("createAuthentication:%s", p.Email)

		res, err := t.limiter.AllowN(
			t.allowCtx,
			rlKey,
			redis_rate.PerHour(t.cfg.AuthenticationRateLimit),
			t.cfg.AuthenticationRateLimit,
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
		_, p, pf, _, _ := t.authenticate()

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
		authenticated, _, _, a, _ := t.authenticate()

		var response list
		err := t.api.Post(listQuery, &response, authenticated)

		t.NoError(err)
		t.Len(response.Viewer.Authentications.Edges, 1)

		id := t.toUUID(response.Viewer.Authentications.Edges[0].Node.ID)
		t.Equal(a.ID, id)
	})

	t.Run("list_without_authentication", func() {
		var response list
		err := t.api.Post(listQuery, &response)

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
		authenticated, _, _, a, _ := t.authenticate()

		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(a.ID)),
		)

		t.NoError(err)
		t.Equal(a.ID, t.toUUID(response.DeleteAuthentication.AuthenticationID))

		exists, err := t.data.Authentication.
			Query().
			Where(authentication.IDEQ(a.ID)).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exists)
	})

	t.Run("delete_when_nonexistent", func() {
		authenticated, _, _, _, _ := t.authenticate()

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
		authenticated, _, _, _, _ := t.authenticate()
		a := t.factory.Authentication().Create(t.allowCtx)

		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(a.ID)),
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
		t.ErrorContains(err, reporting.ErrUnauthorized.Message)
	})

	// Performing this test on a single action of a model is enough.
	t.Run("delete_with_invalid_id", func() {
		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			client.Var("id", "aeiou"),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrInput.Message)
	})
}
