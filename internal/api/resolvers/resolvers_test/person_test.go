package resolvers_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/alexedwards/argon2id"
	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
	"github.com/avptp/brain/internal/messaging/templates"
	"github.com/stretchr/testify/mock"
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

		t.messenger.On(
			"Send",
			mock.IsType(&templates.Welcome{}),
			mock.IsType(&data.Person{}),
		).Return(nil).Once()

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

		id := t.toUUID(response.CreatePerson.Person.ID)

		p, err := t.data.Person.
			Query().
			Where(person.IDEQ(id)).
			First(t.allowCtx)

		t.NoError(err)

		match, err := argon2id.ComparePasswordAndHash(input.Password, p.Password)
		t.NoError(err)

		t.Equal(input.Email, p.Email)
		t.True(match)
		t.Equal(input.TaxID, p.TaxID)
		t.Equal(input.FirstName, p.FirstName)
		t.Equal(input.LastName, p.LastName)
		t.Equal(input.Language, p.Language)

		exists, err := t.data.Authorization.
			Query().
			Where(authorization.PersonIDEQ(p.ID)).
			Exist(t.allowCtx)

		t.NoError(err)
		t.True(exists)

		t.messenger.AssertExpectations(t.T())
	})

	t.Run("create_with_already_used_email", func() {
		factory := t.factory.Person()
		input := factory.Fields
		factory.Create(t.allowCtx)

		var response create
		err := t.api.Post(
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

		t.ErrorContains(err, reporting.ErrConstraint.Message)
	})

	const showQuery = `
		query {
			viewer {
				id
			}
		}
	`

	type show struct {
		Viewer struct {
			ID string
		}
	}

	t.Run("show", func() {
		authenticated, p, _, _, _ := t.authenticate()

		var response show
		t.api.MustPost(showQuery, &response, authenticated)

		id := t.toUUID(response.Viewer.ID)

		t.Equal(p.ID, id)
	})

	t.Run("show_without_authentication", func() {
		var response show
		err := t.api.Post(showQuery, &response)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthenticated.Message)
	})

	const updateMutation = `
		mutation(
			$id: ID!,
			$email: String,
			$phone: String,
			$taxId: String,
			$firstName: String,
			$lastName: String,
			$language: String,
			$birthdate: Time,
			$gender: Gender,
			$address: String,
			$postalCode: String,
			$city: String,
			$country: String,
		) {
			updatePerson(input: {
				id: $id
				email: $email
				phone: $phone
				taxId: $taxId
				firstName: $firstName
				lastName: $lastName
				language: $language
				birthdate: $birthdate
				gender: $gender
				address: $address
				postalCode: $postalCode
				city: $city
				country: $country
			}) {
				person {
					id
				}
			}
		}
	`

	type update struct {
		UpdatePerson struct {
			Person struct {
				ID string
			}
		}
	}

	t.Run("update", func() {
		authenticated, p, _, _, _ := t.authenticate()

		input := t.factory.Person().Fields

		var response update
		t.api.MustPost(
			updateMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(p.ID)),
			client.Var("email", input.Email),
			client.Var("phone", input.Phone),
			client.Var("taxId", input.TaxID),
			client.Var("firstName", input.FirstName),
			client.Var("lastName", input.LastName),
			client.Var("language", input.Language),
			client.Var("birthdate", input.Birthdate),
			client.Var("gender", input.Gender),
			client.Var("address", input.Address),
			client.Var("postalCode", input.PostalCode),
			client.Var("city", input.City),
			client.Var("country", input.Country),
		)

		t.Equal(p.ID, t.toUUID(response.UpdatePerson.Person.ID))

		p, err := t.data.Person.
			Query().
			Where(person.IDEQ(p.ID)).
			First(t.allowCtx)

		t.NoError(err)
		t.Equal(input.Email, p.Email)
		t.Nil(p.EmailVerifiedAt)
		t.NotNil(p.Phone)
		t.Equal(*input.Phone, *p.Phone)
		t.Equal(input.TaxID, p.TaxID)
		t.Equal(input.FirstName, p.FirstName)
		t.NotNil(p.LastName)
		t.Equal(*input.LastName, *p.LastName)
		t.Equal(input.Language, p.Language)
		t.NotNil(p.Birthdate)
		t.True(input.Birthdate.Equal(*p.Birthdate))
		t.NotNil(p.Gender)
		t.Equal(input.Gender.String(), p.Gender.String())
		t.NotNil(p.Address)
		t.Equal(*input.Address, *p.Address)
		t.NotNil(p.PostalCode)
		t.Equal(*input.PostalCode, *p.PostalCode)
		t.NotNil(p.City)
		t.Equal(*input.City, *p.City)
		t.NotNil(p.Country)
		t.Equal(*input.Country, *p.Country)
	})

	t.Run("update_nonexistent", func() {
		authenticated, _, _, _, _ := t.authenticate()

		var response update
		err := t.api.Post(
			updateMutation,
			&response,
			authenticated,
			client.Var("id", zeroID),
			client.Var("lastName", nil),
		)

		t.ErrorContains(err, reporting.ErrNotFound.Message)
	})

	t.Run("update_without_authorization", func() {
		authenticated, _, _, _, _ := t.authenticate()
		p := t.factory.Person().Create(t.allowCtx)

		var response update
		err := t.api.Post(
			updateMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(p.ID)),
			client.Var("lastName", nil),
		)

		t.ErrorContains(err, reporting.ErrUnauthorized.Message)
	})

	t.Run("update_without_authentication", func() {
		var response update
		err := t.api.Post(
			updateMutation,
			&response,
			client.Var("id", zeroID),
			client.Var("lastName", nil),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrUnauthorized.Message)
	})

	const updatePasswordMutation = `
		mutation(
			$id: ID!,
			$currentPassword: String!,
			$newPassword: String!,
			$captcha: String!,
		) {
			updatePersonPassword(input: {
				id: $id
				currentPassword: $currentPassword
				newPassword: $newPassword
				captcha: $captcha
			}) {
				person {
					id
				}
			}
		}
	`

	type updatePassword struct {
		UpdatePersonPassword struct {
			Person struct {
				ID string
			}
		}
	}

	t.Run("update_password", func() {
		authenticated, p, pf, _, _ := t.authenticate()

		input := t.factory.Person().Fields

		var response updatePassword
		t.api.MustPost(
			updatePasswordMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(p.ID)),
			client.Var("currentPassword", pf.Password),
			client.Var("newPassword", input.Password),
			client.Var("captcha", captchaResponseToken),
		)

		t.Equal(p.ID, t.toUUID(response.UpdatePersonPassword.Person.ID))

		p, err := t.data.Person.
			Query().
			Where(person.IDEQ(p.ID)).
			First(t.allowCtx)

		t.NoError(err)

		match, err := argon2id.ComparePasswordAndHash(input.Password, p.Password)
		t.NoError(err)
		t.True(match)
	})

	t.Run("update_password_with_wrong_password", func() {
		authenticated, p, pf, _, _ := t.authenticate()

		input := t.factory.Person().Fields

		var response updatePassword
		err := t.api.Post(
			updatePasswordMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(p.ID)),
			client.Var("currentPassword", pf.Password+"wrong"),
			client.Var("newPassword", input.Password),
			client.Var("captcha", captchaResponseToken),
		)

		t.ErrorContains(err, reporting.ErrWrongPassword.Message)

		p, err = t.data.Person.
			Query().
			Where(person.IDEQ(p.ID)).
			First(t.allowCtx)

		t.NoError(err)

		match, err := argon2id.ComparePasswordAndHash(pf.Password, p.Password)
		t.NoError(err)
		t.True(match)
	})

	t.Run("update_password_when_nonexistent", func() {
		authenticated, _, _, _, _ := t.authenticate()

		var response updatePassword
		err := t.api.Post(
			updatePasswordMutation,
			&response,
			authenticated,
			client.Var("id", zeroID),
			client.Var("currentPassword", "password"),
			client.Var("newPassword", ""),
			client.Var("captcha", captchaResponseToken),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)
	})

	t.Run("update_password_without_authorization", func() {
		authenticated, _, _, _, _ := t.authenticate()
		p := t.factory.Person().Create(t.allowCtx)

		var response updatePassword
		err := t.api.Post(
			updatePasswordMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(p.ID)),
			client.Var("currentPassword", "password"),
			client.Var("newPassword", ""),
			client.Var("captcha", captchaResponseToken),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)
	})

	t.Run("update_password_without_authentication", func() {
		var response updatePassword
		err := t.api.Post(
			updatePasswordMutation,
			&response,
			client.Var("id", zeroID),
			client.Var("currentPassword", ""),
			client.Var("newPassword", ""),
			client.Var("captcha", captchaResponseToken),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)
	})

	const deleteMutation = `
		mutation(
			$id: ID!,
			$currentPassword: String!,
			$captcha: String!,
		) {
			deletePerson(input: {
				id: $id
				currentPassword: $currentPassword
				captcha: $captcha
			}) {
				personId
			}
		}
	`

	type delete struct {
		DeletePerson struct {
			PersonID string
		}
	}

	t.Run("delete", func() {
		authenticated, p, pf, _, _ := t.authenticate()

		var response delete
		t.api.MustPost(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(p.ID)),
			client.Var("currentPassword", pf.Password),
			client.Var("captcha", captchaResponseToken),
		)

		t.Equal(p.ID, t.toUUID(response.DeletePerson.PersonID))

		exists, err := t.data.Person.
			Query().
			Where(person.IDEQ(p.ID)).
			Exist(t.allowCtx)

		t.NoError(err)
		t.False(exists)
	})

	t.Run("delete_with_wrong_password", func() {
		authenticated, p, pf, _, _ := t.authenticate()

		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(p.ID)),
			client.Var("currentPassword", pf.Password+"wrong"),
			client.Var("captcha", captchaResponseToken),
		)

		t.ErrorContains(err, reporting.ErrWrongPassword.Message)

		exists, err := t.data.Person.
			Query().
			Where(person.IDEQ(p.ID)).
			Exist(t.allowCtx)

		t.NoError(err)
		t.True(exists)
	})

	t.Run("delete_when_nonexistent", func() {
		authenticated, _, _, _, _ := t.authenticate()

		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", zeroID),
			client.Var("currentPassword", "password"),
			client.Var("captcha", captchaResponseToken),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)
	})

	t.Run("delete_without_authorization", func() {
		authenticated, _, _, _, _ := t.authenticate()
		p := t.factory.Person().Create(t.allowCtx)

		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			authenticated,
			client.Var("id", t.toULID(p.ID)),
			client.Var("currentPassword", "password"),
			client.Var("captcha", captchaResponseToken),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)
	})

	t.Run("delete_without_authentication", func() {
		var response delete
		err := t.api.Post(
			deleteMutation,
			&response,
			client.Var("id", zeroID),
			client.Var("currentPassword", "password"),
			client.Var("captcha", captchaResponseToken),
		)

		t.NotNil(err)
		t.ErrorContains(err, reporting.ErrWrongPassword.Message)
	})
}
