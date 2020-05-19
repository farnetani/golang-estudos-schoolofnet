package main

import (
	"fmt"
	"github.com/farnetani/exemplo-rotas-simples/modules/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	// Criando uma rota static
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", controller.HomeHandler)
	r.HandleFunc("/{id}/view", controller.ViewHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}

