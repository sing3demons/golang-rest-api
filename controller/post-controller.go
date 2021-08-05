package controller

import (
	"encoding/json"
	"net/http"
	"sing3demons/go-rest-api/entity"
	"sing3demons/go-rest-api/service"
	"sing3demons/go-rest-api/utils"

	"github.com/go-chi/chi/v5"
)

type ResponsePost struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
}

type postController struct {
	service service.PostService
}

func NewPostController(service service.PostService) PostController {
	return &postController{service: service}
}

func (sv *postController) GetPost(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["id"]
	id := chi.URLParam(r, "id") 

	post, err := sv.service.FindOne(id)
	if err != nil {
		utils.ErrorToJson(w, http.StatusInternalServerError)("Error getting the post")
		return
	}

	utils.Json(w, http.StatusOK)(post)
}

func (sv *postController) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := sv.service.FindAll()
	if err != nil {
		utils.ErrorToJson(w, http.StatusInternalServerError)("Error getting the posts")
		return
	}

	// w.Write(result)
	utils.Json(w, http.StatusOK)(posts)

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

	var resp ResponsePost
	resp.Title = p.Title
	resp.Text = p.Text

	utils.Json(w, http.StatusCreated)(resp)
}
