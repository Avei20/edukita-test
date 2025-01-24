package service

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"backend/pkg/response"
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type (
	userImpl struct {
		repo repository.User
	}

	User interface {
		CreateUser(
			ctx context.Context,
			data CreateUserBody,
		) (*LoginResponse, error)
		Login(
			ctx context.Context,
			email string,
		) (*LoginResponse, error)
	}

	CreateUserResponse = response.BaseResponse[entity.User]

	CreateUserBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	LoginData struct {
		Token string      `json:"token"`
		User  entity.User `json:"user"`
	}

	LoginResponse = response.BaseResponse[LoginData]

	TokenClaims struct {
		jwt.Claims
		User entity.User
	}
)

type (
	assignmentImpl struct {
		repo repository.Assignment
	}

	Assignment interface {
		CreateAssignment(
			ctx context.Context,
			data CreateAssignmentBody,
		) (*CreateAssignmentResponse, error)
		GetAllAssignmentsByTeacher(
			ctx context.Context,
			filterby *string,
		) (*GetAllAssignmentsResponse, error)
	}

	CreateAssignmentResponse = response.BaseResponse[entity.Assignment]

	CreateAssignmentBody struct {
		Title   string `json:"title"`
		Subject string `json:"subject"`
		Content string `json:"content"`
	}

	GetAllAssignmentsResponse = response.BaseResponse[[]entity.Assignment]
)

type (
	gradeImpl struct {
		repo repository.Grade
	}

	Grade interface {
		CreateGrade(
			ctx context.Context,
			data CreateGradeBody,
		) (*CreateGradeResponse, error)
		GetAllGradesByStudent(
			ctx context.Context,
			studentID string,
		) (*GetAllGradesResponse, error)
	}

	CreateGradeResponse = response.BaseResponse[entity.Grade]

	CreateGradeBody struct {
		AssignmentID string `json:"assignment_id"`
		Grade        int    `json:"grade"`
		Feeback      string `json:"feedback"`
	}

	GetAllGradesResponse = response.BaseResponse[[]entity.Grade]
)
