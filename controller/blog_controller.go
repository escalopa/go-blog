package controller

import (
	"encoding/json"
	"github.com/escalopa/go-blog/database"
	"github.com/escalopa/go-blog/entities"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPosts(w http.ResponseWriter, _ *http.Request) {
	// fetch post
	var posts []entities.Post
	database.Instance.Find(&posts)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	// check if post exists
	var postId = mux.Vars(r)["id"]
	if isPostExists(postId) == false {
		err := json.NewEncoder(w).Encode("Post Not Found")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// fetch `post` from `db`
	var post entities.Post
	var err error
	database.Instance.First(&post, postId)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// parse post
	var post entities.Post
	var err error
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create post
	database.Instance.Create(&post)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// set response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// check if post exists
	var postId = mux.Vars(r)["id"]
	if isPostExists(postId) == false {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("Post Not Found")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		return
	}

	// save post
	var post entities.Post
	database.Instance.First(&post, postId)
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	database.Instance.Save(&post)

	// set response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Post Updates")
	w.WriteHeader(http.StatusOK)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	var postId = mux.Vars(r)["id"]
	if isPostExists(postId) == false {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("Post Not Found")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		return
	}
	var post entities.Post
	database.Instance.Delete(&post, postId)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode("Post Deleted")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func isPostExists(postId string) bool {
	var post entities.Post
	database.Instance.First(&post, postId)
	return post.Id != 0
}
