package resolvers_test

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/avptp/brain/internal/api/auth/auth_test"
	"github.com/avptp/brain/internal/api/types"
	"github.com/avptp/brain/internal/billing/billing_test"
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/data/validation"
	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/factories"
	"github.com/avptp/brain/internal/generated/data/privacy"
	_ "github.com/avptp/brain/internal/generated/data/runtime"
	"github.com/avptp/brain/internal/messaging/messaging_test"
	"github.com/avptp/brain/internal/services"
	"github.com/avptp/brain/internal/transport"
	"github.com/go-redis/redis_rate/v10"

	"github.com/99designs/gqlgen/client"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/suite"
)

const zeroID = "00000000000000000000000000"
const spanishTaxIdLetters = "TRWAGMYFPDXBNJZSQVHLCKE"

func init() {
	_ = faker.AddProvider("phone", func(_ reflect.Value) (any, error) {
		phone := fmt.Sprintf(
			"+346%08d",
			rand.Intn(99999999),
		)

		return &phone, nil
	})

	_ = faker.AddProvider("tax_id", func(_ reflect.Value) (any, error) {
		number := rand.Intn(99999999)
		letter := spanishTaxIdLetters[number%23]

		return fmt.Sprintf("%08d%c", number, letter), nil
	})

	_ = faker.AddProvider("birthdate", func(_ reflect.Value) (any, error) {
		start := time.Now().AddDate(-100, 0, 0)
		end := time.Now()

		days := int(end.Sub(start).Hours() / 24)
		randomDays := rand.Intn(days)

		birthdate := start.AddDate(0, 0, randomDays)

		return &birthdate, nil
	})

	_ = faker.AddProvider("address", func(_ reflect.Value) (any, error) {
		addr := faker.GetRealAddress()

		return &addr.Address, nil
	})

	_ = faker.AddProvider("postal_code", func(_ reflect.Value) (any, error) {
		addr := faker.GetRealAddress()

		return &addr.PostalCode, nil
	})

	_ = faker.AddProvider("city", func(_ reflect.Value) (any, error) {
		addr := faker.GetRealAddress()

		return &addr.City, nil
	})

	_ = faker.AddProvider("country", func(_ reflect.Value) (any, error) {
		all := validation.Countries.FindAllCountries()

		keys := make([]string, 0, len(all))

		for key := range all {
			keys = append(keys, key)
		}

		index := rand.Intn(len(keys))
		key := keys[index]

		country := all[key]

		return &country.Alpha2, nil
	})
}

type TestSuite struct {
	suite.Suite

	ctn       *container.Container
	biller    *billing_test.MockedBiller
	captcha   *auth_test.MockedCaptcha
	cfg       *config.Config
	data      *data.Client
	limiter   *redis_rate.Limiter
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

	t.biller = &billing_test.MockedBiller{}
	err = builder.Set(services.Biller, t.biller)

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
	t.cfg = ctn.GetConfig()
	t.data = ctn.GetData()
	t.limiter = ctn.GetLimiter()

	t.factory = factories.New(t.data)
	t.api = client.New(transport.GraphHandler(ctn))
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

func (t *TestSuite) authenticate() (client.Option, *data.Person, factories.PersonFields, *data.Authentication) {
	return t.authenticateWith(nil, nil)
}

func (t *TestSuite) authenticateWith(prsCb func(*data.PersonCreate), authnCb func(*data.AuthenticationCreate)) (client.Option, *data.Person, factories.PersonFields, *data.Authentication) {
	// Create person
	personFactory := t.factory.Person()

	if prsCb != nil {
		personFactory.With(prsCb)
	}

	person := personFactory.Create(t.allowCtx)
	personFields := personFactory.Fields

	// Create authentication
	authnFactory := t.factory.
		Authentication().
		With(func(a *data.AuthenticationCreate) {
			a.SetPerson(person)
		})

	if authnCb != nil {
		authnFactory.With(authnCb)
	}

	authn := authnFactory.Create(t.allowCtx)

	// Create client options
	option := func(bd *client.Request) {
		bd.HTTP.Header.Add(
			"Authorization",
			fmt.Sprintf("Bearer %s", authn.TokenEncoded()),
		)
	}

	return option, person, personFields, authn
}

func (t *TestSuite) parseID(str string) types.ID {
	id, err := types.ParseID(str)
	t.NoError(err)

	return id
}
