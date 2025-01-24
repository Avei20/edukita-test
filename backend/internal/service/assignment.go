package service

import (
	"backend/internal/entity"
	contextkeys "backend/pkg/contextKeys"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *assignmentImpl) CreateAssignment(
	ctx context.Context,
	data CreateAssignmentBody,
) (*CreateAssignmentResponse, error) {
	userId, isExist := contextkeys.GetValue[string](ctx, contextkeys.UserId)

	if !isExist {
		return nil, fmt.Errorf("Unknown User")
	}

	assignment, err := s.repo.CreateAssignment(entity.Assignment{
		ID:      uuid.New().String(),
		Title:   data.Title,
		Content: data.Content,
		Subject: data.Subject,
		// UserID from Context.
		StudentId: userId,
	})

	if err != nil {
		return nil, err
	}

	return &CreateAssignmentResponse{
		Message:    "Success",
		Data:       assignment,
		StatusCode: http.StatusCreated,
		Success:    true,
	}, nil
}

func (s *assignmentImpl) GetAllAssignmentsByTeacher(
	ctx context.Context,
	filterBy *string,
) (*GetAllAssignmentsResponse, error) {
	assignments, err := s.repo.FindAllAssignments()

	if err != nil {
		return nil, err
	}

	if filterBy != nil {
		assignments, err = s.repo.FindAllAssignmentBySubject(*filterBy)
		if err != nil {
			return nil, err
		}
	}

	return &GetAllAssignmentsResponse{
		Message:    "Success",
		Data:       assignments,
		StatusCode: http.StatusOK,
		Success:    true,
	}, nil
}

// func (s *assignmentImpl) GetAllAssignmentsByStudent(
// 	ctx context.Context,
// ) (*GetAllAssignmentsResponse, error) {
