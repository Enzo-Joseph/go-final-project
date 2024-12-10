package controllers

import (
	"go-final-project/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AllStudents(c *gin.Context) {
	students, err := models.GetStudents()

	if err != nil {
		panic("Error in controllers/student_controller.go/AllStudents")
	}
	c.JSON(http.StatusOK, students)
}

func Enroll(c *gin.Context) {
	courseID := c.Query("course_id")

	session := sessions.Default(c)
	studentID := session.Get("user_id").(int)

	err := models.EnrollToCourse(courseID, studentID)
	if err != nil {
		panic("Error in controllers/student_controller.go/Enroll")
	}
	c.JSON(http.StatusOK, gin.H{"message": "Enroll successful"})
}

func Unenroll(c *gin.Context) {
	courseID := c.Query("course_id")

	session := sessions.Default(c)
	studentID := session.Get("user_id").(int)

	err := models.RemoveEnrollment(courseID, studentID)
	if err != nil {
		panic("Error in controllers/student_controller.go/Unenroll")
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unenroll successful"})
}