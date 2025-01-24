package service

import (
	"backend/internal/entity"
	contextkeys "backend/pkg/contextKeys"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *gradeImpl) CreateGrade(
	ctx context.Context,
	data CreateGradeBody,
) (*CreateGradeResponse, error) {
	userId, isExist := contextkeys.GetValue[string](ctx, contextkeys.UserId)

	if !isExist {
		return nil, fmt.Errorf("Unknown User")
	}

	grade, err := s.repo.CreateGrade(entity.Grade{
		ID:           uuid.New().String(),
		AssignmentID: data.AssignmentID,
		TeacherID:    userId,
		Grade:        data.Grade,
		Feedback:     data.Feeback,
	})

	if err != nil {
		return nil, err
	}

	return &CreateGradeResponse{
		Message:    "Success",
		Data:       grade,
		StatusCode: http.StatusCreated,
		Success:    true,
	}, nil
}

func (s *gradeImpl) GetAllGradesByStudent(
	ctx context.Context,
	studentID string,
) (*GetAllGradesResponse, error) {
	userId, isExist := contextkeys.GetValue[string](ctx, contextkeys.UserId)

	if !isExist {
		return nil, fmt.Errorf("Unknown User")
	}

	if userId != studentID {
		return nil, fmt.Errorf("Unauthorized")
	}

	grades, err := s.repo.FindAllGradesByStudentID(userId)

	if err != nil {
		return nil, err
	}

	return &GetAllGradesResponse{
		Message:    "Success",
		Data:       grades,
		StatusCode: http.StatusOK,
		Success:    true,
	}, nil
}
