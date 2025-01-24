package middleware

import (
	contextkeys "backend/pkg/contextKeys"
	"backend/pkg/response"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"os"

	"github.com/golang-jwt/jwt/v5"
)

func JwtValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		publicPath := []string{
			"/login",
			"/users",
		}

		for _, path := range publicPath {
			if r.URL.Path == path {
				next.ServeHTTP(w, r)
				return
			}
		}

		authorizationHeader := r.Header.Get("Authorization")

		// Validating JWT Token
		if !strings.Contains(authorizationHeader, "Bearer") {
			log.Println(authorizationHeader)
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			method, ok := t.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, fmt.Errorf("invalid Signing Method")
			}

			if method != jwt.SigningMethodHS256 {
				return nil, nil
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			log.Printf("[Middleware][JwtValidator] Error: %v\n", err)
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
			return
		}

		// Save UserId and Role to context from clams

		userId := claims["sub"].(string)
		role := claims["role"].(string)

		log.Println(role)

		ctx := context.WithValue(r.Context(), contextkeys.UserId, userId)
		ctx = context.WithValue(ctx, contextkeys.Role, role)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
