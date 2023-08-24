package resolvers_test

import (
	"encoding/base64"

	"github.com/99designs/gqlgen/client"
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/data/authentication"
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
		_, p, input, _, _ := t.authenticate()

		var response create
		t.api.MustPost(
			createMutation,
			&response,
			client.Var("email", input.Email),
			client.Var("password", input.Password),
		)

		token, err := base64.URLEncoding.DecodeString(response.CreateAuthentication.Token)
		t.Nil(err)

		atc, err := t.data.Authentication.
			Query().
			WithPerson().
			Where(authentication.TokenEQ(token)).
			First(t.ctx)

		t.Nil(err)
		t.Equal(p.ID, atc.Edges.Person.ID)
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
		_, p, current, _, _ := t.authenticate()

		var response create
		err := t.api.Post(
			createMutation,
			&response,
			client.Var("email", p.Email),
			client.Var("password", current.Password+"wrong"),
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
		t.api.MustPost(listQuery, &response, authenticated)

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
		t.api.MustPost(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(a.ID)),
		)

		t.Equal(a.ID, t.toUUID(response.DeleteAuthentication.AuthenticationID))

		exists, err := t.data.Authentication.
			Query().
			Where(authentication.IDEQ(a.ID)).
			Exist(t.ctx)

		t.Nil(err)
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
		a := t.factory.Authentication().Create(t.ctx)

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