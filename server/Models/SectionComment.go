package Models

import "gorm.io/gorm"

type SectionComment struct {
	gorm.Model
	SectionID   uint
	Description string
}
