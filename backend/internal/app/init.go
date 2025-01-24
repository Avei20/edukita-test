package app

import (
	"backend/internal/entity"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
)

func InitHttp() *Server {
	users := map[string]entity.User{}
	assignments := map[string]entity.Assignment{}
	grades := map[string]entity.Grade{}

	userRepository := repository.NewUser(users)
	assignmentRepository := repository.NewAssignment(assignments)
	gradesRepository := repository.NewGrade(grades, assignments)

	userService := service.NewUser(userRepository)
	assignmentService := service.NewAssignment(assignmentRepository)
	gradeService := service.NewGrade(gradesRepository)

	handler := handler.NewHandler(userService, assignmentService, gradeService)

	server := NewServer(handler)
	return server
}
