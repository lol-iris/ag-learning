package Models

import (
	"time"

	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	ParentSectionID uint
	ParentSection   *Section
	Title           string
	Description     string
	Attachments     []SectionAttachment
	Visible         bool
	VisibleAfter    time.Time
	Comments        []SectionComment
}
