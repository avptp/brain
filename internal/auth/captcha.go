package auth

import (
	"context"

	"github.com/avptp/brain/internal/transport/request"
)

func Captcha(ctx context.Context, token string) bool {
	captcha := request.ContainerFromCtx(ctx).GetCaptcha()
	captcha.RemoteIP = string(request.IPFromCtx(ctx))

	return captcha.VerifyToken(token).Success
}
