package container

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	slog "log/slog"

	resolvers "github.com/avptp/brain/internal/api/resolvers"
	auth "github.com/avptp/brain/internal/auth"
	config "github.com/avptp/brain/internal/config"
	data "github.com/avptp/brain/internal/generated/data"
	messaging "github.com/avptp/brain/internal/messaging"
	ses "github.com/aws/aws-sdk-go/service/ses"
	tasks "github.com/madflojo/tasks"
	in "github.com/nicksnyder/go-i18n/v2/i18n"
	realclientipgo "github.com/realclientip/realclientip-go"
)

func getDiDefs(provider dingo.Provider) []di.Def {
	return []di.Def{
		{
			Name:  "captcha",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("captcha")
				if err != nil {
					var eo auth.Captcha
					return eo, err
				}
				pi0, err := ctn.SafeGet("config")
				if err != nil {
					var eo auth.Captcha
					return eo, err
				}
				p0, ok := pi0.(*config.Config)
				if !ok {
					var eo auth.Captcha
					return eo, errors.New("could not cast parameter 0 to *config.Config")
				}
				b, ok := d.Build.(func(*config.Config) (auth.Captcha, error))
				if !ok {
					var eo auth.Captcha
					return eo, errors.New("could not cast build function to func(*config.Config) (auth.Captcha, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "config",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("config")
				if err != nil {
					var eo *config.Config
					return eo, err
				}
				b, ok := d.Build.(func() (*config.Config, error))
				if !ok {
					var eo *config.Config
					return eo, errors.New("could not cast build function to func() (*config.Config, error)")
				}
				return b()
			},
			Unshared: false,
		},
		{
			Name:  "data",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("data")
				if err != nil {
					var eo *data.Client
					return eo, err
				}
				pi0, err := ctn.SafeGet("config")
				if err != nil {
					var eo *data.Client
					return eo, err
				}
				p0, ok := pi0.(*config.Config)
				if !ok {
					var eo *data.Client
					return eo, errors.New("could not cast parameter 0 to *config.Config")
				}
				b, ok := d.Build.(func(*config.Config) (*data.Client, error))
				if !ok {
					var eo *data.Client
					return eo, errors.New("could not cast build function to func(*config.Config) (*data.Client, error)")
				}
				return b(p0)
			},
			Close: func(obj interface{}) error {
				d, err := provider.Get("data")
				if err != nil {
					return err
				}
				c, ok := d.Close.(func(*data.Client) error)
				if !ok {
					return errors.New("could not cast close function to 'func(*data.Client) error'")
				}
				o, ok := obj.(*data.Client)
				if !ok {
					return errors.New("could not cast object to '*data.Client'")
				}
				return c(o)
			},
			Unshared: false,
		},
		{
			Name:  "i18n",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("i18n")
				if err != nil {
					var eo *in.Bundle
					return eo, err
				}
				b, ok := d.Build.(func() (*in.Bundle, error))
				if !ok {
					var eo *in.Bundle
					return eo, errors.New("could not cast build function to func() (*in.Bundle, error)")
				}
				return b()
			},
			Unshared: false,
		},
		{
			Name:  "ipStrategy",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("ipStrategy")
				if err != nil {
					var eo realclientipgo.Strategy
					return eo, err
				}
				pi0, err := ctn.SafeGet("config")
				if err != nil {
					var eo realclientipgo.Strategy
					return eo, err
				}
				p0, ok := pi0.(*config.Config)
				if !ok {
					var eo realclientipgo.Strategy
					return eo, errors.New("could not cast parameter 0 to *config.Config")
				}
				b, ok := d.Build.(func(*config.Config) (realclientipgo.Strategy, error))
				if !ok {
					var eo realclientipgo.Strategy
					return eo, errors.New("could not cast build function to func(*config.Config) (realclientipgo.Strategy, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "logger",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("logger")
				if err != nil {
					var eo *slog.Logger
					return eo, err
				}
				pi0, err := ctn.SafeGet("config")
				if err != nil {
					var eo *slog.Logger
					return eo, err
				}
				p0, ok := pi0.(*config.Config)
				if !ok {
					var eo *slog.Logger
					return eo, errors.New("could not cast parameter 0 to *config.Config")
				}
				b, ok := d.Build.(func(*config.Config) (*slog.Logger, error))
				if !ok {
					var eo *slog.Logger
					return eo, errors.New("could not cast build function to func(*config.Config) (*slog.Logger, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "messenger",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("messenger")
				if err != nil {
					var eo messaging.Messenger
					return eo, err
				}
				pi0, err := ctn.SafeGet("config")
				if err != nil {
					var eo messaging.Messenger
					return eo, err
				}
				p0, ok := pi0.(*config.Config)
				if !ok {
					var eo messaging.Messenger
					return eo, errors.New("could not cast parameter 0 to *config.Config")
				}
				pi1, err := ctn.SafeGet("ses")
				if err != nil {
					var eo messaging.Messenger
					return eo, err
				}
				p1, ok := pi1.(*ses.SES)
				if !ok {
					var eo messaging.Messenger
					return eo, errors.New("could not cast parameter 1 to *ses.SES")
				}
				pi2, err := ctn.SafeGet("i18n")
				if err != nil {
					var eo messaging.Messenger
					return eo, err
				}
				p2, ok := pi2.(*in.Bundle)
				if !ok {
					var eo messaging.Messenger
					return eo, errors.New("could not cast parameter 2 to *in.Bundle")
				}
				b, ok := d.Build.(func(*config.Config, *ses.SES, *in.Bundle) (messaging.Messenger, error))
				if !ok {
					var eo messaging.Messenger
					return eo, errors.New("could not cast build function to func(*config.Config, *ses.SES, *in.Bundle) (messaging.Messenger, error)")
				}
				return b(p0, p1, p2)
			},
			Unshared: false,
		},
		{
			Name:  "resolver",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("resolver")
				if err != nil {
					var eo *resolvers.Resolver
					return eo, err
				}
				pi0, err := ctn.SafeGet("config")
				if err != nil {
					var eo *resolvers.Resolver
					return eo, err
				}
				p0, ok := pi0.(*config.Config)
				if !ok {
					var eo *resolvers.Resolver
					return eo, errors.New("could not cast parameter 0 to *config.Config")
				}
				pi1, err := ctn.SafeGet("captcha")
				if err != nil {
					var eo *resolvers.Resolver
					return eo, err
				}
				p1, ok := pi1.(auth.Captcha)
				if !ok {
					var eo *resolvers.Resolver
					return eo, errors.New("could not cast parameter 1 to auth.Captcha")
				}
				pi2, err := ctn.SafeGet("data")
				if err != nil {
					var eo *resolvers.Resolver
					return eo, err
				}
				p2, ok := pi2.(*data.Client)
				if !ok {
					var eo *resolvers.Resolver
					return eo, errors.New("could not cast parameter 2 to *data.Client")
				}
				pi3, err := ctn.SafeGet("messenger")
				if err != nil {
					var eo *resolvers.Resolver
					return eo, err
				}
				p3, ok := pi3.(messaging.Messenger)
				if !ok {
					var eo *resolvers.Resolver
					return eo, errors.New("could not cast parameter 3 to messaging.Messenger")
				}
				b, ok := d.Build.(func(*config.Config, auth.Captcha, *data.Client, messaging.Messenger) (*resolvers.Resolver, error))
				if !ok {
					var eo *resolvers.Resolver
					return eo, errors.New("could not cast build function to func(*config.Config, auth.Captcha, *data.Client, messaging.Messenger) (*resolvers.Resolver, error)")
				}
				return b(p0, p1, p2, p3)
			},
			Unshared: false,
		},
		{
			Name:  "scheduler",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("scheduler")
				if err != nil {
					var eo *tasks.Scheduler
					return eo, err
				}
				b, ok := d.Build.(func() (*tasks.Scheduler, error))
				if !ok {
					var eo *tasks.Scheduler
					return eo, errors.New("could not cast build function to func() (*tasks.Scheduler, error)")
				}
				return b()
			},
			Close: func(obj interface{}) error {
				d, err := provider.Get("scheduler")
				if err != nil {
					return err
				}
				c, ok := d.Close.(func(*tasks.Scheduler) error)
				if !ok {
					return errors.New("could not cast close function to 'func(*tasks.Scheduler) error'")
				}
				o, ok := obj.(*tasks.Scheduler)
				if !ok {
					return errors.New("could not cast object to '*tasks.Scheduler'")
				}
				return c(o)
			},
			Unshared: false,
		},
		{
			Name:  "ses",
			Scope: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("ses")
				if err != nil {
					var eo *ses.SES
					return eo, err
				}
				pi0, err := ctn.SafeGet("config")
				if err != nil {
					var eo *ses.SES
					return eo, err
				}
				p0, ok := pi0.(*config.Config)
				if !ok {
					var eo *ses.SES
					return eo, errors.New("could not cast parameter 0 to *config.Config")
				}
				b, ok := d.Build.(func(*config.Config) (*ses.SES, error))
				if !ok {
					var eo *ses.SES
					return eo, errors.New("could not cast build function to func(*config.Config) (*ses.SES, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
	}
}
