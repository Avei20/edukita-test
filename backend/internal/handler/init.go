package handler

import "backend/internal/service"

func NewHandler(
	userService service.User,
	assignmentService service.Assignment,
	gradeService service.Grade,
) Handlers {
	return &handlerImpl{
		userService:       userService,
		assignmentService: assignmentService,
		gradeService:      gradeService,
	}
}
