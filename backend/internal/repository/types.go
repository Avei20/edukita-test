package repository

import "backend/internal/entity"

type (
	User interface {
		CreateUser(user entity.User) (entity.User, error)
		FindUserByEmail(email string) (*entity.User, error)
	}
	userImpl struct {
		users map[string]entity.User
	}

	Assignment interface {
		CreateAssignment(assignment entity.Assignment) (entity.Assignment, error)
		IsAssignmentExist(assignmentId string) bool
		FindAssignmentByID(assignmentId string) (entity.Assignment, error)
		FindAllAssignments() ([]entity.Assignment, error)
		FindAllAssignmentBySubject(subject string) ([]entity.Assignment, error)
	}
	assignmentImpl struct {
		assignments map[string]entity.Assignment
	}

	Grade interface {
		CreateGrade(grade entity.Grade) (entity.Grade, error)
		IsGradeExist(gradeId string) bool
		FindGradeByID(gradeId string) (entity.Grade, error)
		FindAllGradesByStudentID(studentId string) ([]entity.Grade, error)
	}
	gradeImpl struct {
		grades     map[string]entity.Grade
		assigments map[string]entity.Assignment
	}
)
