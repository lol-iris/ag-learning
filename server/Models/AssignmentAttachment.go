package Models

import "gorm.io/gorm"

type AssignmentAttachment struct {
	gorm.Model
	AssignmentID uint
	Name         string
	Attachment   []byte
}
