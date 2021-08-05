package service

import (
	"errors"
	"sing3demons/go-rest-api/entity"
	"sing3demons/go-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindOne(id string) (*entity.Post, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}

	if post.Text == "" {
		err := errors.New("The post text is empty")
		return err
	}

	return nil
}

func (s *postService) Create(post *entity.Post) (*entity.Post, error) {
	return s.repo.Save(post)
}

func (s *postService) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}

func (s *postService) FindOne(id string) (*entity.Post, error) {

	return s.repo.FindOne(id)
}
