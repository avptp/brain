package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/authentication"
	"github.com/avptp/brain/internal/generated/data/privacy"
	"github.com/avptp/brain/internal/transport/request"
)

func NewSetViewer(data *data.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := strings.Split(
				r.Header.Get("authorization"),
				"Bearer ",
			)

			// The authorization header is invalid
			if len(header) != 2 {
				next.ServeHTTP(w, r)
				return
			}

			token, err := base64.URLEncoding.DecodeString(
				strings.TrimSpace(header[1]),
			)

			// The token is not correctly encoded
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// The token does not exist
			ctx := r.Context()
			allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

			auth, err := data.
				Authentication.
				Query().
				Where(authentication.TokenEQ(token)).
				WithPerson().
				First(allowCtx)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ip, ok := ctx.Value(request.IPCtxKey{}).(string)

			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			ctx = context.WithValue(ctx, request.ViewerCtxKey{}, auth.Edges.Person)

			auth.
				Update().
				SetLastUsedIP(ip).
				SetLastUsedAt(time.Now()).
				SaveX(allowCtx)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
