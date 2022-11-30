package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"loyalty-system/pkg/logger"
	"net/http"
	"strconv"
	"strings"
)

type Middleware func(http.Handler) http.Handler

func Authenticate(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := logger.GetLogger()
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Warnf("Token was not passed.")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			authHeaderParts := strings.Split(authHeader, " ")
			if len(authHeaderParts) != 2 {
				log.Warnf("Token must consist of 2 parts, but %d were given", len(authHeaderParts))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			token, err := jwt.Parse(authHeaderParts[1], func(jwtToken *jwt.Token) (interface{}, error) {
				if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
					err := fmt.Errorf("unexpected signing method %s", jwtToken.Header["alg"])
					return nil, err
				}
				return []byte(secret), nil
			})
			if err != nil {
				log.Error(err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if !token.Valid {
				log.Warn("token is invalid")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			claims := token.Claims.(jwt.MapClaims)
			userID, err := strconv.Atoi(claims["sub"].(string))
			if err != nil {
				log.Errorf("could not get user_id from token claim %s", err.Error())
			}
			ctx := context.WithValue(r.Context(), "user_id", userID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
