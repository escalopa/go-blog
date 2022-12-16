package main

import (
	"log"
	"net/http"

	"github.com/escalopa/goblog/config"
	"github.com/escalopa/goblog/controller"
	"github.com/escalopa/goblog/database"
	"github.com/gorilla/mux"
)

func RegisterBlogRoutes(router *mux.Router) {
	router.HandleFunc("/api/posts", controller.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controller.GetPostById).Methods("GET")
	router.HandleFunc("/api/posts", controller.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", controller.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", controller.DeletePost).Methods("DELETE")
}

func main() {
	config := config.New()
	database.Connect(config.Get("DATABASE_URL"))
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterBlogRoutes(router)

	log.Print("Server up and running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
