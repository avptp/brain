package templates

import (
	"github.com/matcornic/hermes/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Template interface {
	Name() string
	Email() Composer
}

type Composer func(l *i18n.Localizer) hermes.Email
