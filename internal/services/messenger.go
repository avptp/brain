package services

import (
	"github.com/avptp/brain/internal/config"
	"github.com/avptp/brain/internal/messaging"
	"github.com/aws/aws-sdk-go/service/ses"
	i "github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Messenger = "messenger"

var MessengerDef = dingo.Def{
	Name:  Messenger,
	Scope: di.App,
	Build: func(cfg *config.Config, ses *ses.SES, i18n *i.Bundle) (messaging.Messenger, error) {
		return messaging.NewMessenger(
			cfg,
			ses,
			i18n,
		), nil
	},
}
