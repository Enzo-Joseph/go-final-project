package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"size:255"`
	Description string
	LecturerID int
	LecturerName string
	StudentCount int
}

func GetCourses() ([]Course, error) {
	var courses []Course

	res := DB.Table("courses").
		Select("courses.id, courses.name, courses.description, courses.lecturer_id, users.name as lecturer_name, COUNT(student_courses.student_id) as student_count").
		Joins("left join users on courses.lecturer_id=users.id").
		Joins("left join student_courses on courses.id=student_courses.course_id").
		Group("courses.id").
		Find(&courses)

	return courses, res.Error
}

func GetCourseByID(id string) (Course, error) {
	var course Course

	res := DB.Table("courses").
		Select("courses.id, courses.name, courses.description, courses.lecturer_id, users.name as lecturer_name, COUNT(student_courses.student_id) as student_count").
		Joins("left join users on courses.lecturer_id=users.id").
		Joins("left join student_courses on courses.id=student_courses.course_id").
		Where("courses.id=?", id).
		Group("courses.id").
		Find(&course)

	return course, res.Error
}

func GetCoursesByStudentID(studentID int) ([]Course, error) {
	var courses []Course

	res := DB.Table("courses").
		Select("courses.id, courses.name, courses.description, courses.lecturer_id, users.name as lecturer_name, COUNT(student_courses.student_id) as student_count").
		Joins("left join users as lecturer on courses.lecturer_id=users.id").
		Joins("left join student_courses on courses.id=student_courses.course_id").
		Where("student_courses.student_id=?", studentID).
		Group("courses.id").
		Find(&courses)

	return courses, res.Error
}

func GetCoursesByLecturerID(lecturerID string) ([]Course, error) {
	var courses []Course

	res := DB.Find(&courses, "lecturer_id=?", lecturerID)

	return courses, res.Error
}
