package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"size:255"`
	LecturerID int
}

func GetCourses() ([]Course, error) {
	var courses []Course

	res := DB.Find(&courses)

	return courses, res.Error
}

func GetCourseByID(id string) (Course, error) {
	var course Course

	res := DB.First(&course, id)

	return course, res.Error
}

func GetCoursesByStudentID(studentID int) ([]Course, error) {
	var courses []Course

	res := DB.Table("courses").Joins("left join student_course on courses.id=student_course.course_id").Where("student_course.student_id=?", studentID).Find(&courses)

	return courses, res.Error
}

func GetCoursesByLecturerID(lecturerID string) ([]Course, error) {
	var courses []Course

	res := DB.Find(&courses, "lecturer_id=?", lecturerID)

	return courses, res.Error
}
