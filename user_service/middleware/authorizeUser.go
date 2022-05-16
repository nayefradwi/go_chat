package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/nayefradwi/go_chat/user_service/config"
	"github.com/nayefradwi/go_chat_common"
)

type UserIdKey struct{}

func AuthorizeHeaderMiddleware(f http.Handler) http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenSplit := strings.Split(authHeader, "Bearer ")
		if len(tokenSplit) < 2 {
			gochatcommon.WriteErrorResponse(w, gochatcommon.NewUnAuthorizedError())
			return
		}
		token := tokenSplit[1]
		userId, err := gochatcommon.DecodeAccessToken(token, config.Secret)
		if err != nil {
			gochatcommon.WriteErrorResponse(w, err)
			return
		}
		rcopy := r.WithContext(context.WithValue(r.Context(), UserIdKey{}, userId))
		f.ServeHTTP(w, rcopy)
	})
	return handler
}
