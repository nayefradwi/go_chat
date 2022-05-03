package auth

import (
	"context"
	"gochat/errorHandling"
	"gochat/goChatUtil"
	"net/http"
	"strings"
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
		userId, err := DecodeAccessToken(token)
		if err != nil {
			goChatUtil.WriteErrorResponse(w, err)
			return
		}
		rcopy := r.WithContext(context.WithValue(r.Context(), UserIdKey{}, userId))
		f.ServeHTTP(w, rcopy)
	})
	return handler
}
