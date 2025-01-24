package middleware

import (
	"backend/pkg/response"
	"log"
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				log.Println("Preflight request")
				response.SetRawResponse(w, http.StatusOK, nil)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
