package Models

import "gorm.io/gorm"

type Announcement struct {
	gorm.Model
	CourseID    uint
	Title       string
	Description string
}
