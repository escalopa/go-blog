package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/escalopa/goblog/database"
	"github.com/escalopa/goblog/entities"
	"github.com/gorilla/mux"
)

// GetPosts godoc
// @Summary      Show all Posts
// @Description  get all posts
// @Tags         posts
// @Produce      json
// @Success      200  {object} 	 []entities.Post
// @Failure      400  {object} 	string
// @Failure      404  {object} 	string
// @Failure      500  {object}	string
// @Router       /posts/ [get]
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
}

// GetPostByID godoc
// @Summary      Show one Post
// @Description  get post by id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Post ID"
// @Success      200  {object} 	entities.Post
// @Failure      400  {object} 	string
// @Failure      404  {object}	string
// @Failure      500  {object}	string
// @Router       /posts/{id} [get]
func GetPostById(w http.ResponseWriter, r *http.Request) {
	// check if post exists
	postId := mux.Vars(r)["id"]
	if !isPostExists(postId) {
		handleNotFound(w, r)
		return
	}

	// fetch `post` from `db`
	var post entities.Post
	var err error
	tx := database.Instance.First(&post, postId)
	if tx.Error != nil {
		if tx.Error.Error() == "record not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set response
	w.Header().Set("Content-Type", "application/json")
}

type PostRequestParam struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

// CreatePost godoc
// @Summary      Create a Post
// @Description  create a post with title, content and description
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   body PostRequestParam  true  "Post Param"
// @Success      200  {object}	entities.Post
// @Failure      400  {object} 	string
// @Failure      404  {object}	string
// @Failure      500  {object}	string
// @Router       /posts/ [post]
func CreatePost(w http.ResponseWriter, r *http.Request) {
	req, err := parsePostBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var post entities.Post
	mapRequestToPost(req, &post)

	// create post
	tx := database.Instance.Create(&post)
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// set response
	w.Header().Set("Content-Type", "application/json")
}

// UpdatePost godoc
// @Summary      update one Post
// @Description  update post by id with title, content and description
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Param        id   body PostRequestParam  true  "Post Param"
// @Success      200  {object}	entities.Post
// @Failure      400  {object} 	string
// @Failure      404  {object}	string
// @Failure      500  {object}	string
// @Router       /posts/{id} [put]
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// check if post exists
	postId := parseReqParamID(r)
	if !isPostExists(postId) {
		handleNotFound(w, r)
		return
	}

	// read request body
	req, err := parsePostBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// save post
	var post entities.Post
	database.Instance.First(&post, postId)
	mapRequestToPost(req, &post)
	database.Instance.Save(&post)

	// set response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Post Updated")
}

// DeletePost godoc
// @Summary      Delete one Post
// @Description  delete post by id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}	string
// @Failure      400  {object} 	string
// @Failure      404  {object}	string
// @Failure      500  {object}	string
// @Router       /posts/{id} [delete]
func DeletePost(w http.ResponseWriter, r *http.Request) {
	var postId = mux.Vars(r)["id"]
	if !isPostExists(postId) {
		handleNotFound(w, r)
		return
	}
	database.Instance.Delete(&entities.Post{}, postId)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode("Post Deleted")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func isPostExists(postId string) bool {
	tx := database.Instance.First(&entities.Post{}, postId)
	return tx.Error == nil
}

func mapRequestToPost(req PostRequestParam, post *entities.Post) {
	post.Title = req.Title
	post.Content = req.Content
	post.Description = req.Description
}

func parsePostBody(r *http.Request) (PostRequestParam, error) {
	var req PostRequestParam
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, err
	}

	if req.Title == "" || req.Content == "" || req.Description == "" {
		return req, errors.New("invalid request, title, content and description are required")
	}

	return req, nil
}

func parseReqParamID(r *http.Request) string {
	var postId = mux.Vars(r)["id"]
	return postId
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := json.NewEncoder(w).Encode("Post Not Found")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
