package controllers

import (
	"go-final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	session := sessions.Default(c)
	userID := session.Get("user_id")
	userName := session.Get("user_name")
	userRole := session.Get("user_role")

	c.HTML(
		http.StatusOK,
		"courses.html",
		gin.H{
			"courses": courses,
			"user_id":  userID,
			"user_name": userName,
			"user_role": userRole,
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

	session := sessions.Default(c)

	if session.Get("user_id") == nil {
		c.HTML(
			http.StatusOK,
			"course.html",
			gin.H{
				"course": course,
			},
		)
		return
	}

	userID := session.Get("user_id").(int)
	userName := session.Get("user_name")
	userRole := session.Get("user_role")

	// check if user is enrolled in course
	isEnrolled := false
	enrollment, err := models.CheckEnrollment(id, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	if enrollment.ID != 0 {
		isEnrolled = true
	}

	c.HTML(
		http.StatusOK,
		"course.html",
		gin.H{
			"course":    course,
			"posts":     posts,
			"user_id":   userID,
			"user_name": userName,
			"user_role": userRole,
			"isEnrolled": isEnrolled,
		},
	)
}
