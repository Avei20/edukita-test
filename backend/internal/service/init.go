package service

import "backend/internal/repository"

func NewUser(userRepository repository.User) User {
	return &userImpl{
		repo: userRepository,
	}
}

func NewAssignment(assignmentRepository repository.Assignment) Assignment {
	return &assignmentImpl{
		repo: assignmentRepository,
	}
}

func NewGrade(gradeRepository repository.Grade) Grade {
	return &gradeImpl{
		repo: gradeRepository,
	}
}
