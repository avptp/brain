package resolvers_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/factories"
	"github.com/avptp/brain/internal/generated/data/privacy"
	_ "github.com/avptp/brain/internal/generated/data/runtime"
	"github.com/avptp/brain/internal/transport"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

const zeroID = "00000000000000000000000000"
const captchaResponseToken = "10000000-aaaa-bbbb-cccc-000000000001"
const spanishTaxIdLetters = "TRWAGMYFPDXBNJZSQVHLCKE"

type TestSuite struct {
	suite.Suite

	ctn     *container.Container
	log     *zap.SugaredLogger
	data    *data.Client
	factory *factories.Factory
	api     *client.Client
	ctx     context.Context
}

func (t *TestSuite) SetupSuite() {
	ctn, err := container.NewContainer()

	if err != nil {
		panic(err) // unrecoverable situation
	}

	t.ctn = ctn
	t.log = ctn.GetLogger()
	t.data = ctn.GetData()
	t.factory = factories.New(t.data)
	t.api = client.New(transport.Mux(ctn))
	t.ctx = privacy.DecisionContext(context.Background(), privacy.Allow)

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

func TestResolvers(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (t *TestSuite) TearDownSuite() {
	err := t.ctn.Delete()

	if err != nil {
		t.log.Fatal(err)

		// flush buffers again, since container has just been deleted
		// intentionally ignoring error here, see https://github.com/uber-go/zap/issues/328
		_ = t.log.Sync()
	}
}

func (t *TestSuite) authenticate() (client.Option, *data.Person, factories.PersonFields, *data.Authentication, factories.AuthenticationFields) {
	personFactory := t.factory.Person()
	person := personFactory.Create(t.ctx)
	personFields := personFactory.Fields

	authFactory := t.factory.Authentication().With(func(a *data.AuthenticationCreate) {
		a.SetPerson(person)
	})
	auth := authFactory.Create(t.ctx)
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
	t.Nil(err)

	return uuid.UUID(ulid)
}

func (t *TestSuite) toULID(id uuid.UUID) string {
	ulid := ulid.ULID(id)

	return ulid.String()
}
