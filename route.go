package main

import (
	"sing3demons/go-rest-api/database"
	"sing3demons/go-rest-api/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	db   *mongo.Database           = database.ConnectMongoDB()
	repo repository.PostRepository = repository.NewMongoRepository(db)
)
