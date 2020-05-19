package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type Post struct {
	Id int
	Title string
	Body string
}

var db, err = sql.Open("mysql", "root:farsoft01@/go_course?charset=utf8")

// Função para facilitar o tratamento do erro
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	stmt, err := db.Prepare("insert into posts(title, body) values(?,?)")

	checkErr(err)

	_ , err = stmt.Exec("My First Post", "My first content")
	checkErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		post := Post{ Id: 1, Title: "Fist Post", Body: "Our content"}

		if title := r.FormValue("title"); title != ""{
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
