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

	//stmt, err := db.Prepare("insert into posts(title, body) values(?,?)")
	//
	//checkErr(err)
	//
	//_ , err = stmt.Exec("My First Post", "My first content")
	//checkErr(err)
	//
	//db.Close()

	rows, err := db.Query("Select * from posts")
	checkErr(err)

	items := []Post{} // Slices de posts

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}
	db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// fmt.Fprintf(w, "Hello World")
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
