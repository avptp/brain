package services

import (
	"github.com/BurntSushi/toml"
	"github.com/avptp/brain/internal/i18n"
	i "github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"golang.org/x/text/language"
)

const I18n = "i18n"

var I18nDef = dingo.Def{
	Name:  I18n,
	Scope: di.App,
	Build: func() (*i.Bundle, error) {
		bundle := i.NewBundle(language.Catalan)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

		files, err := i18n.FS.ReadDir(".")

		if err != nil {
			return nil, err
		}

		for _, file := range files {
			bundle.LoadMessageFileFS(i18n.FS, file.Name())
		}

		return bundle, nil
	},
}
