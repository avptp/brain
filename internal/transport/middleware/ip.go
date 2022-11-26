package middleware

import (
	"context"
	"net"
	"net/http"

	"github.com/avptp/brain/internal/transport/request"
	"github.com/realclientip/realclientip-go"
)

func NewSetIP(strategy realclientip.Strategy) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := strategy.ClientIP(r.Header, "")

			if ip == "" {
				remote, _, err := net.SplitHostPort(r.RemoteAddr)

				if err != nil {
					panic(err) // unrecoverable situation
				}

				ip = remote
			}

			ctx := context.WithValue(
				r.Context(),
				request.IPCtxKey{},
				ip,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
