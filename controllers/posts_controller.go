package controllers

import (
	"net/http"

	"go-final-project/models"

	"github.com/gin-gonic/gin"
)

func PostByID(c *gin.Context) {
	id := c.Param("id")
	post, err := models.GetPostByID(id)

	if err != nil {
		panic(err)
	}

	c.HTML(
		http.StatusOK,
		"post.html",
		gin.H{
			"post": post,
		},
	)
}

func PostsByCourse(c *gin.Context) {
	courseID := c.Param("course_id")
	posts, err := models.GetPostsByCourse(courseID)

	if err != nil {
		panic(err)
	}

	c.HTML(
		http.StatusOK,
		"posts.html",
		gin.H{
			"posts": posts,
		},
	)
}
