package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"size:255"`
	Description string
	LecturerID int
	LecturerName string
	StudentCount int
}

type StudentCourse struct {
	gorm.Model
	CourseID string
	StudentID int
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
		Joins("left join users on courses.lecturer_id=users.id").
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

func EnrollToCourse(courseID string, studentID int) error {
	enrollment := StudentCourse{CourseID: courseID, StudentID: studentID}

	return DB.Create(&enrollment).Error
}

func RemoveEnrollment(courseID string, studentID int) error {
	return DB.Unscoped().Where("course_id = ?", courseID).Where("student_id = ?", studentID).Delete(&StudentCourse{}).Error
}

func CheckEnrollment(courseID string, studentID int) (StudentCourse, error) {
	var enrollment StudentCourse

	res := DB.Table("student_courses").
		Where("course_id = ?", courseID).
		Where("student_id = ?", studentID).
		First(&enrollment)

	if res.Error != nil {
		return StudentCourse{}, res.Error
	}

	return enrollment, nil
}

func EditDescription(courseID string, description string) error {
	result := DB.Table("courses").Where("id = ?", courseID).Update("description", description)
	if result.Error != nil {
		return result.Error
	}

	return nil
}