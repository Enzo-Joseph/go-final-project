package controllers

import (
	"go-final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllTeachers(c *gin.Context) {
	teachers, err := models.GetTeachers()

	if err != nil {
		panic("Error in controllers/teachers_controller.go/AllTeachers")
	}
	c.JSON(http.StatusOK, teachers)
}
