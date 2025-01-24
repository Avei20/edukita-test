package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"net/http"
)

func InitRouter(h handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	// Use Middleware

	mux.HandleFunc("POST /users", h.CreateUser)
	mux.HandleFunc("POST /login", h.Login)
	mux.HandleFunc("POST /assignment", h.CreateAssignment)
	mux.HandleFunc("GET /assignment", h.GetAssignment)
	mux.HandleFunc("POST /grade", h.CreateGrade)
	mux.HandleFunc("GET /grade/{studentid}", h.GetGrade)

	muxCors := middleware.CORS(mux)
	muxJwt := middleware.JwtValidator(muxCors)
	muxMiddleware := middleware.RoleValidator(muxJwt)

	return muxMiddleware
}
