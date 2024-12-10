package main

import (
	"go-final-project/controllers"
	"go-final-project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "static")
	r.GET("/", controllers.HomePage)

	r.GET("/students", controllers.AllStudents)
	r.GET("/lecturers", controllers.AllLecturers)
	r.GET("/lecturers/:id", controllers.LecturerByID)
	r.GET("/courses", controllers.AllCourses)
	r.GET("/courses/:id", controllers.CoursesByID)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
