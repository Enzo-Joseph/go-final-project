package models

import "gorm.io/gorm"

type Lecturer struct {
	gorm.Model
	ID   int64
	Name string
}

func GetLecturers() ([]Lecturer, error) {
	var lecturers []Lecturer

	res := DB.Find(&lecturers)

	return lecturers, res.Error
}

func GetLecturerByID(id string) (Lecturer, error) {
	var lecturer Lecturer

	res := DB.First(&lecturer, id)

	return lecturer, res.Error
}
