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

	// mux.Handle("/users/", InitUserRoute(h))
	// mux.Handle("/assignments/", InitAssignmentRoute(h))
	// mux.Handle("/grades/", InitGradeRoute(h))
	muxMiddleware := middleware.RoleValidator(mux)
	muxCors := middleware.CORS(muxMiddleware)
	muxJwt := middleware.JwtValidator(muxCors)

	return muxJwt
}
