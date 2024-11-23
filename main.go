package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var templates = template.Must(template.ParseFiles("edit.html", "view.html", "home.html"))

// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Student struct {
	Id   int64
	Name string
}

type Teacher struct {
	Id   int64
	Name string
}
type Course struct {
	Id       int64
	Name     string
	Teacher  Teacher
	Students []Student
}

type CourseHolder struct {
	Courses []Course
}

type Post struct {
	Id     int64
	Course Course
	title  string
	body   string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "classrooms",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/view/", makeHandler(viewHandler))
	// http.HandleFunc("/edit/", makeHandler(editHandler))
	// http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := getCourses()
	course_holder := CourseHolder{courses}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Courses found: %v\n", courses)

	renderTemplate(w, "home", course_holder)
}

// func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	_, err := loadPage(title)
// 	if err != nil {
// 		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 		return
// 	}
// 	renderTemplate(w, "view")
// }

// func editHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	_, err := loadPage(title)
// 	if err != nil {
// 		// p = &Page{Title: title}
// 	}
// 	renderTemplate(w, "edit")
// }

// func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	body := r.FormValue("body")
// 	p := &Page{Title: title, Body: []byte(body)}
// 	err := p.save()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

// func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m := validPath.FindStringSubmatch(r.URL.Path)
// 		if m == nil {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		fn(w, r, m[2])
// 	}
// }

func getStudents() ([]Student, error) {
	// An albums slice to hold data from returned rows.
	var students []Student

	rows, err := db.Query("SELECT * FROM Students")
	// fmt.Print(rows)
	if err != nil {
		return nil, fmt.Errorf("getStudents : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Id, &student.Name); err != nil {
			return nil, fmt.Errorf("getStudents : %v", err)
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getStudents : %v", err)
	}
	return students, nil
}

func getStudentsByCourseId(course_id int) ([]Student, error) {
	// Get the students who follow the course
	var students []Student
	rows, err := db.Query("SELECT Students.id, Students.name FROM StudentCourse JOIN Courses ON StudentCourse.course_id=Courses.id JOIN Students ON StudentCourse.student_id=Students.id WHERE Courses.id=?", course_id)
	if err != nil {
		return nil, fmt.Errorf("getStudentsByCourseId : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Id, &student.Name); err != nil {
			return nil, fmt.Errorf("getStudentsByCourseId : %v", err)
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getStudentsByCourseId : %v", err)
	}
	return students, nil
}

func getTeachers() ([]Teacher, error) {
	// An albums slice to hold data from returned rows.
	var teachers []Teacher

	rows, err := db.Query("SELECT * FROM Teachers")
	// fmt.Print(rows)
	if err != nil {
		return nil, fmt.Errorf("getTeachers : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var teacher Teacher
		if err := rows.Scan(&teacher.Id, &teacher.Name); err != nil {
			return nil, fmt.Errorf("getTeachers : %v", err)
		}
		teachers = append(teachers, teacher)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getTeachers : %v", err)
	}
	return teachers, nil
}

func getCourses() ([]Course, error) {
	// An albums slice to hold data from returned rows.
	var courses []Course

	rows, err := db.Query("SELECT Courses.id, Courses.name, Teachers.id, Teachers.name FROM Courses JOIN Teachers ON Courses.teacher_id=Teachers.id")
	if err != nil {
		return nil, fmt.Errorf("getCourses : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.Id, &course.Name, &course.Teacher.Id, &course.Teacher.Name); err != nil {
			return nil, fmt.Errorf("getCourses : %v", err)
		}
		students, err := getStudentsByCourseId(int(course.Id))
		if err != nil {
			return nil, fmt.Errorf("getCourses : %v", err)
		}
		course.Students = students

		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getCourses : %v", err)
	}
	return courses, nil
}
