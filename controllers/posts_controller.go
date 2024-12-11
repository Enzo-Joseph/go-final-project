package controllers

import (
	"net/http"

	"go-final-project/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostByID(c *gin.Context) {
	id := c.Param("id")
	post, err := models.GetPostByID(id)

	if err != nil {
		panic(err)
	}

	session := sessions.Default(c)
	userID := session.Get("user_id")
	userName := session.Get("user_name")
	userRole := session.Get("user_role")

	c.HTML(
		http.StatusOK,
		"post.html",
		gin.H{
			"post": post,
			"user_id":  userID,
			"user_name": userName,
			"user_role": userRole,
		},
	)
}

func PostsByCourse(c *gin.Context) {
	courseID := c.Param("course_id")
	posts, err := models.GetPostsByCourse(courseID)

	if err != nil {
		panic(err)
	}

	session := sessions.Default(c)
	userID := session.Get("user_id")
	userName := session.Get("user_name")
	userRole := session.Get("user_role")

	c.HTML(
		http.StatusOK,
		"posts.html",
		gin.H{
			"posts": posts,
			"user_id":  userID,
			"user_name": userName,
			"user_role": userRole,
		},
	)
}

func EditPost(c *gin.Context) {
	postID := c.PostForm("post_id")
	title := c.PostForm("title")
	body := c.PostForm("body")

	session := sessions.Default(c)
	userRole := session.Get("user_role")

	if userRole != "lecturer" {
		// return failed
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	// edit description
	err := models.EditPost(postID, title, body)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Post edited",
		})
	}
}