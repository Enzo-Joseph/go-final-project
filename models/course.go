package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	TeacherID int
}

func GetCoursesByStudentID(studentID int) ([]Course, error) {
	var courses []Course

	res := DB.Table("courses").Joins("left join student_course on courses.id=student_course.course_id").Where("student_course.student_id=?", studentID).Find(&courses)

	return courses, res.Error
}
