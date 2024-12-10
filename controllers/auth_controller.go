package controllers

import (
	"go-final-project/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Login(c *gin.Context) {
	// Get username and password from request
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check if username and password are correct
	user, err := models.CheckAuth(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}
	if user == (models.User{}) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	} else {
		// Set session
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_name", user.Name)
		session.Set("user_role", user.Role)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	}
}

func SignUpPage(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	
	c.HTML(http.StatusOK, "sign-up.html", gin.H{})
}

func SignUp(c *gin.Context) {
	// Get username and password from request
	name := c.PostForm("name")
	username := c.PostForm("username")
	role := c.PostForm("role")
	password := c.PostForm("password")

	// Check if role is valid
	if role != "student" && role != "lecturer" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid role"})
		return
	}

	// Check if username already exists
	user, err := models.CheckUsername(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}
	if user != (models.User{}) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username already exists"})
		return
	}

	// Create new user
	err = models.CreateUser(name, username, password, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	// Redirect to login page
	c.Redirect(http.StatusSeeOther, "/login")
}
