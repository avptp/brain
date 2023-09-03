package resolvers_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/avptp/brain/internal/auth/auth_test"
	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/factories"
	"github.com/avptp/brain/internal/generated/data/privacy"
	_ "github.com/avptp/brain/internal/generated/data/runtime"
	"github.com/avptp/brain/internal/messaging/messaging_test"
	"github.com/avptp/brain/internal/services"
	"github.com/avptp/brain/internal/transport"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/suite"
)

const zeroID = "00000000000000000000000000"
const spanishTaxIdLetters = "TRWAGMYFPDXBNJZSQVHLCKE"

func init() {
	gofakeit.AddFuncLookup("tax_id", gofakeit.Info{
		Category: "custom",
		Output:   "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			number := rand.Intn(99999999)
			letter := spanishTaxIdLetters[number%23]

			return fmt.Sprintf("%08d%c", number, letter), nil
		},
	})

	gofakeit.AddFuncLookup("phone_e164", gofakeit.Info{
		Category: "custom",
		Output:   "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			number := rand.Intn(99999999)

			return fmt.Sprintf("+346%08d", number), nil
		},
	})
}

type TestSuite struct {
	suite.Suite

	ctn  *container.Container
	data *data.Client

	captcha   *auth_test.MockedCaptcha
	messenger *messaging_test.MockedMessenger

	factory  *factories.Factory
	api      *client.Client
	allowCtx context.Context
}

func (t *TestSuite) SetupSuite() {
	builder, err := container.NewBuilder()

	if err != nil {
		panic(err) // unrecoverable situation
	}

	t.captcha = &auth_test.MockedCaptcha{}
	err = builder.Set(services.Captcha, t.captcha)

	if err != nil {
		panic(err) // unrecoverable situation
	}

	t.messenger = &messaging_test.MockedMessenger{}
	err = builder.Set(services.Messenger, t.messenger)

	if err != nil {
		panic(err) // unrecoverable situation
	}

	ctn := builder.Build()

	t.ctn = ctn
	t.data = ctn.GetData()

	t.factory = factories.New(t.data)
	t.api = client.New(transport.Mux(ctn))
	t.allowCtx = privacy.DecisionContext(context.Background(), privacy.Allow)
}

func TestResolvers(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) TearDownSuite() {
	err := t.ctn.Delete()

	if err != nil {
		panic(err) // unrecoverable situation
	}
}

func (t *TestSuite) authenticate() (client.Option, *data.Person, factories.PersonFields, *data.Authentication, factories.AuthenticationFields) {
	personFactory := t.factory.Person()
	person := personFactory.Create(t.allowCtx)
	personFields := personFactory.Fields

	authFactory := t.factory.Authentication().With(func(a *data.AuthenticationCreate) {
		a.SetPerson(person)
	})
	auth := authFactory.Create(t.allowCtx)
	authFields := authFactory.Fields

	option := func(bd *client.Request) {
		bd.HTTP.Header.Add(
			"Authorization",
			fmt.Sprintf("Bearer %s", auth.TokenEncoded()),
		)
	}

	return option, person, personFields, auth, authFields
}

func (t *TestSuite) toUUID(id string) uuid.UUID {
	ulid := ulid.ULID{}

	err := ulid.Scan(id)
	t.NoError(err)

	return uuid.UUID(ulid)
}

func (t *TestSuite) toULID(id uuid.UUID) string {
	ulid := ulid.ULID(id)

	return ulid.String()
}
