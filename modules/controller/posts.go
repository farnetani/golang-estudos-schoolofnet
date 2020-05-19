package controller

import (
	"github.com/farnetani/exemplo-rotas-simples/db"
	"github.com/farnetani/exemplo-rotas-simples/modules/model"
	"github.com/farnetani/exemplo-rotas-simples/utils"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("modules/views/layout.html", "modules/views/list.html"))

	if err := t.ExecuteTemplate(w, "layout.html", ListPosts()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	t := template.Must(template.ParseFiles("modules/views/layout.html", "modules/views/view.html"))

	t.ExecuteTemplate(w, "layout.html", GetPostById(id))
}

func ListPosts() []model.Post {
	rows, err := db.Connection().Query("Select * from posts")
	utils.CheckErr(err)

	items := []model.Post{} // Slices de posts

	for rows.Next() {
		post := model.Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}
	db.Connection().Close()
	return items
}

func GetPostById(id string) model.Post {
	row := db.Connection().QueryRow("select * from posts where id=?", id)
	post := model.Post{}
	row.Scan(&post.Id, &post.Title, &post.Body)
	return post
}

