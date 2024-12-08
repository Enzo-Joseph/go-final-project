package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID   int64
	Name string
}

func GetStudents() ([]Student, error) {
	// An albums slice to hold data from returned rows.
	var students []Student

	res := DB.Find(&students)

	return students, res.Error
}

func GetStudentsByCourseID(course_id int) ([]Student, error) {
	// Get the students who follow the course
	var students []Student

	res := DB.Joins("left join Courses on StudentCourse.course_id=Courses.id").Joins("left join Students on StudentCourse.student_id=Student.id").Where("Courses.id=?", course_id).Find(&students)

	return students, res.Error

	// rows, err := db.Query("SELECT Students.id, Students.name FROM StudentCourse JOIN Courses ON StudentCourse.course_id=Courses.id JOIN Students ON StudentCourse.student_id=Students.id WHERE Courses.id=?", course_id)
	// if err != nil {
	// 	return nil, fmt.Errorf("getStudentsByCourseID : %v", err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var student Student
	// 	if err := rows.Scan(&student.ID, &student.Name); err != nil {
	// 		return nil, fmt.Errorf("getStudentsByCourseID : %v", err)
	// 	}
	// 	students = append(students, student)
	// }
	// if err := rows.Err(); err != nil {
	// 	return nil, fmt.Errorf("getStudentsByCourseID : %v", err)
	// }
	// return students, nil
}
