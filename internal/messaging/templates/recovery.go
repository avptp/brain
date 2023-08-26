package templates

import (
	"github.com/matcornic/hermes/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Recovery struct {
	Link     string
	Validity string
}

func (t *Recovery) Name() string {
	return "recovery"
}

func (t *Recovery) Email() Composer {
	return func(l *i18n.Localizer) hermes.Email {
		return hermes.Email{
			Body: hermes.Body{
				Intros: []string{
					l.MustLocalize(&i18n.LocalizeConfig{MessageID: "recovery.intro"}),
				},
				Actions: []hermes.Action{
					{
						Instructions: l.MustLocalize(&i18n.LocalizeConfig{MessageID: "recovery.action.instructions"}),
						Button: hermes.Button{
							Text: l.MustLocalize(&i18n.LocalizeConfig{MessageID: "recovery.action.text"}),
							Link: t.Link,
						},
					},
				},
				Outros: []string{
					l.MustLocalize(&i18n.LocalizeConfig{MessageID: "recovery.outro"}),
				},
			},
		}
	}
}
