package Models

import "gorm.io/gorm"

type AssignmentSubmission struct {
	gorm.Model
	AssignmentID uint
	UserID       uint
	Description  string
	Attachment   []byte
	Grade        int
}
