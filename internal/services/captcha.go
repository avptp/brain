package services

import (
	"github.com/avptp/brain/internal/config"
	"github.com/kataras/hcaptcha"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

const Captcha = "captcha"

var CaptchaDef = dingo.Def{
	Name:  Captcha,
	Scope: di.Request,
	Build: func(cfg *config.Config) (*hcaptcha.Client, error) {
		return hcaptcha.New(cfg.HcaptchaSecret), nil
	},
}
