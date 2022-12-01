package main

import (
	"fmt"
	"github.com/escalopa/goblog/controller"
	"github.com/escalopa/goblog/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterBlogRoutes(router *mux.Router) {
	router.HandleFunc("/api/posts", controller.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controller.GetPostById).Methods("GET")
	router.HandleFunc("/api/posts", controller.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", controller.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", controller.DeletePost).Methods("DELETE")
}

func main() {
	LoadConfiguration()
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterBlogRoutes(router)

	log.Println(fmt.Sprintf("Server up and running on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
