package repository

import (
	"context"
	"sing3demons/go-rest-api/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindOne(id string) (*entity.Post, error)
}

type mongoRepository struct{ db *mongo.Database }

func NewMongoRepository(db *mongo.Database) PostRepository {
	return &mongoRepository{db: db}
}

func (tx *mongoRepository) collection() *mongo.Collection {
	return tx.db.Collection("posts")
}

func (tx *mongoRepository) FindOne(id string) (*entity.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var post entity.Post

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}

	if err := tx.collection().FindOne(ctx, filter).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (tx *mongoRepository) Save(post *entity.Post) (*entity.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := tx.collection().InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (tx *mongoRepository) FindAll() ([]entity.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	posts := []entity.Post{}

	cursor, err := tx.collection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &posts); err != nil {
		panic(err)
	}

	return posts, nil
}
