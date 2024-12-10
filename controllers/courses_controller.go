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
		panic(err)
	}
	c.JSON(http.StatusOK, courses)
}

func AllCourses(c *gin.Context) {
	courses, err := models.GetCourses()

	if err != nil {
		panic(err)
	}

	c.HTML(
		http.StatusOK,
		"courses.html",
		gin.H{
			"courses": courses,
		},
	)
}

func CoursesByID(c *gin.Context) {
	id := c.Param("id")
	course, err := models.GetCourseByID(id)

	if err != nil {
		panic(err)
	}

	posts, err := models.GetPostsByCourse(id)

	if err != nil {
		panic(err)
	}

	c.HTML(
		http.StatusOK,
		"course.html",
		gin.H{
			"course": course,
			"posts":  posts,
		},
	)
}
