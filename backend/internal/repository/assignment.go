package repository

import (
	"backend/internal/entity"
	"fmt"
)

func (r *assignmentImpl) CreateAssignment(assignment entity.Assignment) (entity.Assignment, error) {
	r.assignments[assignment.ID] = assignment
	return assignment, nil
}

func (r *assignmentImpl) IsAssignmentExist(assignmentID string) bool {
	_, ok := r.assignments[assignmentID]
	return ok
}

func (r *assignmentImpl) FindAssignmentByID(assignmentID string) (entity.Assignment, error) {
	assignment, ok := r.assignments[assignmentID]
	if !ok {
		return entity.Assignment{}, fmt.Errorf("assignment with ID %s not found", assignmentID)
	}
	return assignment, nil
}

func (r *assignmentImpl) FindAllAssignments() ([]entity.Assignment, error) {
	var assignments []entity.Assignment
	for _, assignment := range r.assignments {
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

func (r *assignmentImpl) FindAllAssignmentBySubject(subject string) ([]entity.Assignment, error) {
	var assignments []entity.Assignment
	for _, assignment := range r.assignments {
		if assignment.Subject == subject {
			assignments = append(assignments, assignment)
		}
	}
	return assignments, nil
}
