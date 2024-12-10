package controllers

import (
	"go-final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllLecturers(c *gin.Context) {
	lecturers, err := models.GetLecturers()

	if err != nil {
		panic(err)
	}

	c.HTML(
		http.StatusOK,
		"lecturers.html",
		gin.H{
			"lecturers": lecturers,
		},
	)
}

func LecturerByID(c *gin.Context) {
	id := c.Param("id")
	lecturer, err := models.GetLecturerByID(id)

	if err != nil {
		panic(err)
	}

	courses, err := models.GetCoursesByLecturerID(id)

	if err != nil {
		panic(err)
	}

	c.HTML(
		http.StatusOK,
		"lecturer.html",
		gin.H{
			"lecturer": lecturer,
			"courses":  courses,
		},
	)
}
