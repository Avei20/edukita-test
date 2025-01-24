package handler

import (
	"backend/internal/service"
	"net/http"
)

type (
	Handler struct{}

	Handlers interface {
		CreateUser(w http.ResponseWriter, r *http.Request)
		CreateAssignment(w http.ResponseWriter, r *http.Request)
		GetAssignment(w http.ResponseWriter, r *http.Request)
		CreateGrade(w http.ResponseWriter, r *http.Request)
		GetGrade(w http.ResponseWriter, r *http.Request)
		Login(w http.ResponseWriter, r *http.Request)
	}

	handlerImpl struct {
		userService       service.User
		assignmentService service.Assignment
		gradeService      service.Grade
	}
)
