package services

import (
	"github.com/avptp/brain/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Ses = "ses"

var SesDef = dingo.Def{
	Name:  Ses,
	Scope: di.App,
	Build: func(cfg *config.Config) (*ses.SES, error) {
		session, err := session.NewSession(&aws.Config{
			Region: aws.String(cfg.AwsRegion),
			Credentials: credentials.NewStaticCredentials(
				cfg.AwsKeyId,
				cfg.AwsKeySecret,
				"",
			),
		})

		if err != nil {
			return nil, err
		}

		return ses.New(session), nil
	},
}
