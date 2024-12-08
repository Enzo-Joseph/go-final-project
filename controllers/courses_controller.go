package controllers

import (
	"go-final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CoursesByStudentID(c *gin.Context) {
	studentID := c.Param("student_id")
	id, err := strconv.Atoi(studentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	courses, err := models.GetCoursesByStudentID(id)

	if err != nil {
		panic("Error in controllers/courses_controller.go/CoursesByStudentID")
	}
	c.JSON(http.StatusOK, courses)
}
