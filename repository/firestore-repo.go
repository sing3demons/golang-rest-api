package repository

import (
	"context"
	"log"
	"sing3demons/go-rest-api/entity"

	"cloud.google.com/go/firestore"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/iterator"
)

type postRepository struct{}

// func NewFirestoreRepository() PostRepository {
// 	return &postRepository{}
// }

const (
	projectId      string = "progmatic-reviews-52f07"
	collectionName string = "posts"
)

func (r *postRepository) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed to adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*postRepository) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post

	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(primitive.ObjectID),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
