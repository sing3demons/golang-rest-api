package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sing3demons/go-rest-api/entity"
	"sing3demons/go-rest-api/repository"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		ErrorToJson(w, http.StatusInternalServerError)("Error getting the posts")
		return
	}

	w.WriteHeader(http.StatusOK)
	// w.Write(result)
	json.NewEncoder(w).Encode(posts)

}

func addPost(w http.ResponseWriter, r *http.Request) {
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		ErrorToJson(w, http.StatusInternalServerError)("Error marshalling the request")
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
