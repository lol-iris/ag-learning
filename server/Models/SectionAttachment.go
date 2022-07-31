package Models

import "gorm.io/gorm"

type SectionAttachment struct {
	gorm.Model
	SectionID  uint
	Name       string
	Attachment []byte
}
