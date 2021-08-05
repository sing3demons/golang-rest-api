package main

import (
	"net/http"
	"sing3demons/go-rest-api/controller"
	"sing3demons/go-rest-api/database"
	router "sing3demons/go-rest-api/http"
	"sing3demons/go-rest-api/repository"
	"sing3demons/go-rest-api/service"
	"sing3demons/go-rest-api/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

const port = ":8080"

type Map map[string]interface{}

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	db             *mongo.Database           = database.ConnectMongoDB()
	postRepository repository.PostRepository = repository.NewMongoRepository(db)
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	// router := mux.NewRouter()

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	utils.Json(w, http.StatusOK)(Map{"msg": "Hello"})
	// })

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		utils.Json(w, http.StatusOK)("Hello")
	})

	httpRouter.GET("/posts/{id}", postController.GetPost)
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)

	// routes.Serve(router)
	// log.Println("Server listening on port: ", port)
	// log.Fatalln(http.ListenAndServe(":"+port, router))
}
