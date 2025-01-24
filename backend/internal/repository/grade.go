package repository

import (
	"backend/internal/entity"
	"fmt"
)

func (r *gradeImpl) CreateGrade(grade entity.Grade) (entity.Grade, error) {
	r.grades[grade.ID] = grade
	return grade, nil
}

func (r *gradeImpl) IsGradeExist(gradeID string) bool {
	_, ok := r.grades[gradeID]
	return ok
}

func (r *gradeImpl) FindGradeByID(gradeID string) (entity.Grade, error) {
	grade, ok := r.grades[gradeID]
	if !ok {
		return entity.Grade{}, fmt.Errorf("grade with ID %s not found", gradeID)
	}

	return grade, nil
}

func (r *gradeImpl) FindAllGradesByStudentID(studentId string) ([]entity.Grade, error) {
	var grades []entity.Grade

	for _, assigment := range r.assigments {
		if assigment.StudentId == studentId {
			for _, grade := range r.grades {
				if grade.AssignmentID == assigment.ID {
					grades = append(grades, grade)
				}
			}
		}
	}

	return grades, nil
}
