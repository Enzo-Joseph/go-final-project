package main

import (
	"go-final-project/controllers"
	"go-final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Teacher struct {
// 	Id   int64
// 	Name string
// }
// type Course struct {
// 	Id       int64
// 	Name     string
// 	Teacher  Teacher
// 	Students []models.Student
// }

// type CourseHolder struct {
// 	Courses []Course
// }

// type Post struct {
// 	Id     int64
// 	Course Course
// 	title  string
// 	body   string
// }

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/students", controllers.AllStudents)
	r.GET("/teachers", controllers.AllTeachers)
	r.GET("/courses/:student_id", controllers.CoursesByStudentID)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func getTeachers() ([]Teacher, error) {
// 	// An albums slice to hold data from returned rows.
// 	var teachers []Teacher

// 	rows, err := db.Query("SELECT * FROM Teachers")
// 	// fmt.Print(rows)
// 	if err != nil {
// 		return nil, fmt.Errorf("getTeachers : %v", err)
// 	}
// 	defer rows.Close()
// 	// Loop through rows, using Scan to assign column data to struct fields.
// 	for rows.Next() {
// 		var teacher Teacher
// 		if err := rows.Scan(&teacher.Id, &teacher.Name); err != nil {
// 			return nil, fmt.Errorf("getTeachers : %v", err)
// 		}
// 		teachers = append(teachers, teacher)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("getTeachers : %v", err)
// 	}
// 	return teachers, nil
// }

// func getCourses() ([]Course, error) {
// 	// An albums slice to hold data from returned rows.
// 	var courses []Course

// 	rows, err := db.Query("SELECT Courses.id, Courses.name, Teachers.id, Teachers.name FROM Courses JOIN Teachers ON Courses.teacher_id=Teachers.id")
// 	if err != nil {
// 		return nil, fmt.Errorf("getCourses : %v", err)
// 	}
// 	defer rows.Close()
// 	// Loop through rows, using Scan to assign column data to struct fields.
// 	for rows.Next() {
// 		var course Course
// 		if err := rows.Scan(&course.Id, &course.Name, &course.Teacher.Id, &course.Teacher.Name); err != nil {
// 			return nil, fmt.Errorf("getCourses : %v", err)
// 		}
// 		students, err := getStudentsByCourseId(int(course.Id))
// 		if err != nil {
// 			return nil, fmt.Errorf("getCourses : %v", err)
// 		}
// 		course.Students = students

// 		courses = append(courses, course)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("getCourses : %v", err)
// 	}
// 	return courses, nil
// }
