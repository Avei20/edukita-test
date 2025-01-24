package entity

type (
	Grade struct {
		ID           string `json:"id"`
		AssignmentID string `json:"assignment_id"`
		TeacherID    string `json:"teacher_id"`
		Grade        int    `json:"grade"`
		Feedback     string `json:"feedback"`
	}
)
