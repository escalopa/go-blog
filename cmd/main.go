package main

import (
	"log"
	"net/http"
	"os"

	"github.com/escalopa/goblog/config"
	"github.com/escalopa/goblog/controller"
	"github.com/escalopa/goblog/database"
	"github.com/escalopa/goblog/docs"
	_ "github.com/escalopa/goblog/docs"
	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           GoBlog API
// @version         1.0
// @description     This is a simple blog CRUD server.

// @contact.name   API Support

// @license.name  Apache 2.0
func main() {
	config := config.New()
	database.Connect(config.Get("DATABASE_URL"))
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)
	RegisterBlogRoutes(router)
	RegisterSwaggerRoutes(router)

	log.Print("Server up and running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func RegisterBlogRoutes(router *mux.Router) {
	router.HandleFunc("/api/posts/", controller.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controller.GetPostById).Methods("GET")
	router.HandleFunc("/api/posts/", controller.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", controller.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", controller.DeletePost).Methods("DELETE")
}

func RegisterSwaggerRoutes(router *mux.Router) {
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	if docs.SwaggerInfo.Host == "" {
		docs.SwaggerInfo.Host = "localhost:9000"
	}

	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	if docs.SwaggerInfo.BasePath == "" {
		docs.SwaggerInfo.BasePath = "/api"
	}

	prefix := os.Getenv("SWAGGER_PREFIX")
	if prefix == "" {
		prefix = "/swagger/"
	}

	router.PathPrefix(prefix).Handler(httpSwagger.WrapHandler)
}
