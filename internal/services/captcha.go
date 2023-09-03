package services

import (
	"github.com/avptp/brain/internal/auth"
	"github.com/avptp/brain/internal/config"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Captcha = "captcha"

var CaptchaDef = dingo.Def{
	Name:  Captcha,
	Scope: di.App,
	Build: func(cfg *config.Config) (auth.Captcha, error) {
		return auth.NewCaptchaHandler(cfg.HcaptchaSecret), nil
	},
}
