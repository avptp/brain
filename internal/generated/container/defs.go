package container

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	slog "log/slog"

	config "github.com/avptp/brain/internal/config"
	data "github.com/avptp/brain/internal/generated/data"
	hcaptcha "github.com/kataras/hcaptcha"
	realclientipgo "github.com/realclientip/realclientip-go"
)

func getDiDefs(provider dingo.Provider) []di.Def {
	return []di.Def{
		{
			Name:  "captcha",
			Scope: "request",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("captcha")
				if err != nil {
					var eo *hcaptcha.Client
					return eo, err
				}
				pi0, err := ctn.SafeGet("config")
				if err != nil {
					var eo *hcaptcha.Client
					return eo, err
				}
				p0, ok := pi0.(*config.Config)
				if !ok {
					var eo *hcaptcha.Client
					return eo, errors.New("could not cast parameter 0 to *config.Config")
				}
				b, ok := d.Build.(func(*config.Config) (*hcaptcha.Client, error))
				if !ok {
					var eo *hcaptcha.Client
					return eo, errors.New("could not cast build function to func(*config.Config) (*hcaptcha.Client, error)")
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
	}
}
