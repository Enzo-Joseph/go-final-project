package controllers

import (
	"go-final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllStudents(c *gin.Context) {
	students, err := models.GetStudents()

	if err != nil {
		panic("Error in controllers/student_controller.go/AllStudents")
	}
	c.JSON(http.StatusOK, students)
}
