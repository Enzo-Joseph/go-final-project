package main

import (
	"go-final-project/controllers"
	"go-final-project/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "static")

	// Set up session store
	store := cookie.NewStore([]byte("I_LOVE_FRAMEWORK_PROGRAMMING"))
	r.Use(sessions.Sessions("session", store))

	r.GET("/", controllers.HomePage)
	r.GET("/students", controllers.AllStudents)
	r.GET("/lecturers", controllers.AllLecturers)
	r.GET("/lecturers/:id", controllers.LecturerByID)
	r.GET("/courses", controllers.AllCourses)
	r.GET("/courses/:id", controllers.CoursesByID)
	r.GET("/my-courses", controllers.MyCourses)

	r.GET("/post/:id", controllers.PostByID)
	r.GET("/post/new", controllers.NewPostPage)
	r.POST("/api/post", controllers.CreatePost)
	r.POST("/api/edit/post", controllers.EditPost)
	r.POST("/api/delete/post", controllers.DeletePost)

	r.GET("/lecturer-zone", controllers.LecturerZone)
	r.POST("/api/edit/description", controllers.EditDescription)

	r.GET("/login", controllers.LoginPage)
	r.GET("/sign-up", controllers.SignUpPage)
	r.POST("/api/login", controllers.Login)
	r.POST("/api/signup", controllers.SignUp)
	r.GET("/logout", controllers.Logout)

	r.GET("/api/enroll", controllers.Enroll)
	r.GET("/api/unenroll", controllers.Unenroll)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
