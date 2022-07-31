package Models

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	CourseID     uint
	Title        string
	Description  string
	VisibleAfter time.Time
	DueBy        time.Time
	Attachments  []AssignmentAttachment
}
