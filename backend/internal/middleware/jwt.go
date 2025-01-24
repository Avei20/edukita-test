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
		log.Println("JWTValidator")
		if strings.Contains(r.URL.Path, "login") || strings.Contains(r.URL.Path, "users") {
			next.ServeHTTP(w, r)
			return
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

		log.Println(claims)
		log.Println(claims)

		userMap, ok := claims["User"].(map[string]interface{})
		if !ok {
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: missing user data"))
			return
		}

		log.Println(userMap)

		// Check if required claims exist
		userId, ok := userMap["id"]
		if !ok || userId == nil {
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: missing user id"))
			return
		}

		role, ok := userMap["role"]
		if !ok || role == nil {
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: missing role"))
			return
		}

		// Convert claims to string
		userIdStr, ok := userId.(string)
		if !ok {
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: invalid user id format"))
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: invalid role format"))
			return
		}

		log.Printf("User ID: %s, Role: %s", userIdStr, roleStr)

		ctx := context.WithValue(r.Context(), contextkeys.UserId, userIdStr)
		ctx = context.WithValue(ctx, contextkeys.Role, roleStr)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
