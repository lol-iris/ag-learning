package Models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	UserID        uint
	User          User
	Name          string
	Description   string
	Banner        []byte
	Tags          []Tag `gorm:"many2many:course_tags"`
	Announcements []Announcement
	Teachers      []User `gorm:"many2many:course_teachers"`
	Students      []User `gorm:"many2many:course_students"`
	Assignments   []Assignment
}
