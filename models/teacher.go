package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	ID   int64
	Name string
}

func GetTeachers() ([]Teacher, error) {
	// An albums slice to hold data from returned rows.
	var teachers []Teacher

	res := DB.Find(&teachers)

	return teachers, res.Error
}
