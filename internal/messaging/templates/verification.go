package templates

import (
	"github.com/matcornic/hermes/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Verification struct {
	Link     string
	Validity string
}

func (t *Verification) Name() string {
	return "verification"
}

func (t *Verification) Email() Composer {
	return func(l *i18n.Localizer) hermes.Email {
		return hermes.Email{
			Body: hermes.Body{
				Intros: []string{
					l.MustLocalize(&i18n.LocalizeConfig{MessageID: "verification.intro"}),
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
				},
			},
		}
	}
}
