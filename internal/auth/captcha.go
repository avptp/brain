package auth

import (
	"github.com/kataras/hcaptcha"
)

type Captcha interface {
	Verify(token string) bool
}

type CaptchaHandler struct {
	client *hcaptcha.Client
}

func NewCaptchaHandler(secret string) Captcha {
	return &CaptchaHandler{
		client: hcaptcha.New(secret),
	}
}

func (c *CaptchaHandler) Verify(token string) bool {
	return c.client.VerifyToken(token).Success
}
