package templates

import (
	"github.com/matcornic/hermes/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Welcome struct {
	Link     string
	Validity string
}

func (t *Welcome) Name() string {
	return "welcome"
}

func (t *Welcome) Email() Composer {
	return func(l *i18n.Localizer) hermes.Email {
		return hermes.Email{
			Body: hermes.Body{
				Intros: []string{
					l.MustLocalize(&i18n.LocalizeConfig{MessageID: "welcome.intro"}),
				},
				Actions: []hermes.Action{
					{
						Instructions: l.MustLocalize(&i18n.LocalizeConfig{MessageID: "verification.action.instructions"}),
						Button: hermes.Button{
							Text: l.MustLocalize(&i18n.LocalizeConfig{MessageID: "verification.action.text"}),
							Link: t.Link,
						},
					},
				},
				Outros: []string{
					l.MustLocalize(&i18n.LocalizeConfig{
						MessageID: "verification.outro",
						TemplateData: map[string]string{
							"Validity": t.Validity,
						},
					}),
					l.MustLocalize(&i18n.LocalizeConfig{MessageID: "acknowledgment"}),
				},
			},
		}
	}
}
