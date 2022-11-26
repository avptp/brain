package resolvers_test

import (
	"fmt"

	"github.com/99designs/gqlgen/client"
	"github.com/alexedwards/argon2id"
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/data/person"
)

func (t *TestSuite) TestPerson() {
	const createMutation = `
		mutation(
			$email: String!
			$password: String!
			$taxId: String!
			$firstName: String!
			$lastName: String
			$language: String!
			$captcha: String!
		) {
			createPerson(input: {
				email: $email
				password: $password
				taxId: $taxId
				firstName: $firstName
				lastName: $lastName
				language: $language
				captcha: $captcha
			}) {
				person {
					id
				}
			}
		}
	`

	type create struct {
		CreatePerson struct {
			Person struct {
				ID string
			}
		}
	}

	t.Run("create", func() {
		input := t.factory.Person().Fields

		var response create
		t.api.MustPost(
			createMutation,
			&response,
			client.Var("email", input.Email),
			client.Var("password", input.Password),
			client.Var("taxId", input.TaxID),
			client.Var("firstName", input.FirstName),
			client.Var("lastName", input.LastName),
			client.Var("language", input.Language),
			client.Var("captcha", captchaResponseToken),
		)

		id := t.parseID(response.CreatePerson.Person.ID)

		p, err := t.data.Person.
			Query().
			Where(person.IDEQ(id)).
			First(t.ctx)

		t.Nil(err)

		match, err := argon2id.ComparePasswordAndHash(input.Password, p.Password)
		t.Nil(err)

		t.Equal(input.Email, p.Email)
		t.True(match)
		t.Equal(input.TaxID, p.TaxID)
		t.Equal(input.FirstName, p.FirstName)
		t.Equal(input.LastName, p.LastName)
		t.Equal(input.Language, p.Language)
	})
		t.Equal(input.Language, p.Language)
	})
}
