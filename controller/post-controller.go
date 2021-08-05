package controller

import (
	"encoding/json"
	"net/http"
	"sing3demons/go-rest-api/entity"
	"sing3demons/go-rest-api/service"
	"sing3demons/go-rest-api/utils"
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

type postController struct {
	service service.PostService
}

func NewPostController(service service.PostService) PostController {
	return &postController{service: service}
}

func (sv *postController) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	posts, err := sv.service.FindAll()
	if err != nil {
		utils.ErrorToJson(w, http.StatusInternalServerError)("Error getting the posts")
		return
	}

	w.WriteHeader(http.StatusOK)
	// w.Write(result)
	json.NewEncoder(w).Encode(posts)

}

func (sv *postController) AddPost(w http.ResponseWriter, r *http.Request) {
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		utils.ErrorToJson(w, http.StatusInternalServerError)("Error marshalling the request")
		return
	}

	if err := sv.service.Validate(&post); err != nil {
		utils.ErrorToJson(w, http.StatusUnprocessableEntity)(err)
	}

	p, err := sv.service.Create(&post)
	if err != nil {
		utils.ErrorToJson(w, http.StatusUnprocessableEntity)(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
