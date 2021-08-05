package routes

import (
	"net/http"
	"sing3demons/go-rest-api/controller"
	"sing3demons/go-rest-api/database"
	"sing3demons/go-rest-api/repository"
	"sing3demons/go-rest-api/service"

	"github.com/gorilla/mux"
)

func Serve(router *mux.Router) {
	db := database.ConnectMongoDB()

	postRepository := repository.NewMongoRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)

	router.HandleFunc("/posts", postController.GetPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts", postController.AddPost).Methods(http.MethodPost)
}
