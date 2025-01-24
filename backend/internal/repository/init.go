package repository

import "backend/internal/entity"

func NewUser(users map[string]entity.User) User {
	return &userImpl{
		users: users,
	}
}

func NewAssignment(assignments map[string]entity.Assignment) Assignment {
	return &assignmentImpl{
		assignments: assignments,
	}
}

func NewGrade(grades map[string]entity.Grade, assignment map[string]entity.Assignment) Grade {
	return &gradeImpl{
		grades:     grades,
		assigments: assignment,
	}
}
