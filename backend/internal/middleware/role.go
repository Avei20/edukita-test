package middleware

import (
	contextkeys "backend/pkg/contextKeys"
	"backend/pkg/response"
	"fmt"
	"net/http"
	"strings"
)

func RoleValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, isExist := contextkeys.GetValue[string](r.Context(), contextkeys.Role)

		if !isExist {
			response.SetError(w, http.StatusUnauthorized, fmt.Errorf("Unathorized"))
			return
		}

		if strings.Contains(r.URL.Path, "assignment") {
			if r.Method == "POST" && role != "STUDENT" {
				response.SetError(w, http.StatusUnauthorized, fmt.Errorf("Unathorized"))
				return
			}
		}

		if strings.Contains(r.URL.Path, "grade") {
			if r.Method == "POST" && role != "TEACHER" {
				response.SetError(w, http.StatusUnauthorized, fmt.Errorf("Unathorized"))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
