package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	// Criando uma rota static
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{id}/view", ViewHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/list.html"))
	if err := t.ExecuteTemplate(w, "layout.html", ListPosts()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/view.html"))
	//if err := t.ExecuteTemplate(w, "layout.html", GetPostById(id)); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	// Refatorado sem o tratamento do erro
	t.ExecuteTemplate(w, "layout.html", GetPostById(id))
}

func ListPosts() []Post {
	rows, err := db.Query("Select * from posts")
	checkErr(err)

	items := []Post{} // Slices de posts

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}
	db.Close()
	return items
}

func GetPostById(id string) Post {
	row := db.QueryRow("select * from posts where id=?", id)
	post := Post{}
	row.Scan(&post.Id, &post.Title, &post.Body)
	return post
}

