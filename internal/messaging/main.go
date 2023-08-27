package messaging

import (
	"fmt"
	"mime"

	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/messaging/templates"
	"github.com/avptp/brain/internal/messaging/themes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/matcornic/hermes/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const charset = "UTF-8"

type Messenger interface {
	Send(t templates.Template, p *data.Person) error
}

type Mailer struct {
	cfg  *config.Config
	ses  *ses.SES
	i18n *i18n.Bundle
}

func NewMessenger(cfg *config.Config, ses *ses.SES, i18n *i18n.Bundle) Messenger {
	return &Mailer{
		cfg,
		ses,
		i18n,
	}
}

func (m *Mailer) Send(t templates.Template, p *data.Person) error {
	// Instantiate localizer
	l := i18n.NewLocalizer(m.i18n, p.Language)

	// Compose email
	theme := themes.Theme("default")

	h := hermes.Hermes{
		Theme: &theme,
		Product: hermes.Product{
			Name:        m.cfg.OrgName,
			Link:        m.cfg.FrontUrl,
			Logo:        m.cfg.OrgLogo,
			Copyright:   l.MustLocalize(&i18n.LocalizeConfig{MessageID: "sent_by"}),
			TroubleText: l.MustLocalize(&i18n.LocalizeConfig{MessageID: "trouble"}),
		},
	}

	composer := t.Email()
	email := composer(l)

	// Set default title and signature
	email.Body.Title = l.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "greetings",
		TemplateData: map[string]string{
			"Name": p.FirstName,
		},
	})

	email.Body.Signature = l.MustLocalize(&i18n.LocalizeConfig{MessageID: "salutation"})

	// Build HTML and plain text versions
	html, err := h.GenerateHTML(email)

	if err != nil {
		return err
	}

	text, err := h.GeneratePlainText(email)

	if err != nil {
		return err
	}

	// Send email through AWS SES
	_, err = m.ses.SendEmail(&ses.SendEmailInput{
		Source: address(m.cfg.OrgName, m.cfg.MailSource),
		ReplyToAddresses: []*string{
			address(m.cfg.OrgName, m.cfg.MailReplyTo),
		},
		Destination: &ses.Destination{
			ToAddresses: []*string{
				address(p.FullName(), p.Email),
			},
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Charset: aws.String(charset),
				Data: aws.String(
					l.MustLocalize(&i18n.LocalizeConfig{MessageID: fmt.Sprintf(
						"%s.subject",
						t.Name(),
					)}),
				),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(html),
				},
				Text: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(text),
				},
			},
		},
	})

	return err
}

const addressFmt = `%s <%s>`

func address(n string, a string) *string {
	n = mime.QEncoding.Encode(charset, n)
	s := fmt.Sprintf(addressFmt, n, a)

	return &s
}
