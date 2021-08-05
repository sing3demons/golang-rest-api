package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		ErrorToJson(w, http.StatusInternalServerError)("Error marshalling the posts array")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func addPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		ErrorToJson(w, http.StatusInternalServerError)("Error marshalling the request")
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	result, err := json.Marshal(posts)
	if err != nil {
		ErrorToJson(w, http.StatusInternalServerError)("Error marshalling the posts array")
		return
	}
	w.Write(result)
}
