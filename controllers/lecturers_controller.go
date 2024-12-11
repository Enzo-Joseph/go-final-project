package controllers

import (
	"fmt"
	"go-final-project/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AllLecturers(c *gin.Context) {
	lecturers, err := models.GetLecturers()

	if err != nil {
		panic(err)
	}

	session := sessions.Default(c)
	userID := session.Get("user_id")
	userName := session.Get("user_name")
	userRole := session.Get("user_role")

	c.HTML(
		http.StatusOK,
		"lecturers.html",
		gin.H{
			"lecturers": lecturers,
			"user_id":   userID,
			"user_name": userName,
			"user_role": userRole,
		},
	)

	// c.JSON(http.StatusOK, gin.H{
	// 	"lecturers": lecturers,
	// 	"user_id":   userID,
	// 	"user_name": userName,
	// 	"user_role": userRole,
	// })
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

	session := sessions.Default(c)
	userID := session.Get("user_id")
	userName := session.Get("user_name")
	userRole := session.Get("user_role")

	c.HTML(
		http.StatusOK,
		"lecturer.html",
		gin.H{
			"lecturer":  lecturer,
			"courses":   courses,
			"user_id":   userID,
			"user_name": userName,
			"user_role": userRole,
		},
	)
}

func LecturerZone(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	userName := session.Get("user_name")
	userRole := session.Get("user_role")

	if userID == nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	if userRole != "lecturer" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	// get courses from lecturer
	courses, err := models.GetCoursesByLecturerID(fmt.Sprintf("%d", userID))

	if err != nil {
		panic(err)
	}

	c.HTML(
		http.StatusOK,
		"lecturer-zone.html",
		gin.H{
			"courses": courses,
			"user_id":  userID,
			"user_name": userName,
			"user_role": userRole,
		},
	)
}