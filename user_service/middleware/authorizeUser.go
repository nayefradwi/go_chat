package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/nayefradwi/go_chat/common/auth"
	"github.com/nayefradwi/go_chat/common/errorHandling"
	"github.com/nayefradwi/go_chat/common/goChatUtil"
	"github.com/nayefradwi/go_chat/user_service/config"
)

type UserIdKey struct{}

func AuthorizeHeaderMiddleware(f http.Handler) http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenSplit := strings.Split(authHeader, "Bearer ")
		if len(tokenSplit) < 2 {
			goChatUtil.WriteErrorResponse(w, errorHandling.NewUnAuthorizedError())
			return
		}
		token := tokenSplit[1]
		userId, err := auth.DecodeAccessToken(token, config.Secret)
		if err != nil {
			goChatUtil.WriteErrorResponse(w, err)
			return
		}
		rcopy := r.WithContext(context.WithValue(r.Context(), UserIdKey{}, userId))
		f.ServeHTTP(w, rcopy)
	})
	return handler
}
