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

type Message struct {
	Id       int64
	Username string
	Body     string
}

type MessageHolder struct {
	Messages []Message
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "messages",
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
	messages, err := getMessages()
	message_holder := MessageHolder{messages}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Messages found: %v\n", messages)

	renderTemplate(w, "home", message_holder)
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

func getMessages() ([]Message, error) {
	// An albums slice to hold data from returned rows.
	var messages []Message

	rows, err := db.Query("SELECT messages.id, users.username, messages.body FROM messages JOIN users ON users.id=messages.user_id")
	// fmt.Print(rows)
	if err != nil {
		return nil, fmt.Errorf("getMessages : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.Id, &msg.Username, &msg.Body); err != nil {
			return nil, fmt.Errorf("getMessages : %v", err)
		}
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getMessages : %v", err)
	}
	return messages, nil
}
