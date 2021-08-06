package service

import (
	"sing3demons/go-rest-api/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) FindOne(id string) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	id, _ := primitive.ObjectIDFromHex("610c2ee55fc8892a19599aaf")
	post := entity.Post{ID: id, Title: "A", Text: "B"}

	//Setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()

	//Mock Assertion Behavior
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, id, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	id, err := primitive.ObjectIDFromHex("610c2ee55fc8892a19599aaf")
	// var id primitive.ObjectID
	post := entity.Post{ID: id, Title: "A", Text: "B"}

	//Setup expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	result, _ := testService.Create(&post)

	//Mock Assertion Behavior
	mockRepo.AssertExpectations(t)

	assert.NotNil(t, id, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)

}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("")
	post := entity.Post{ID: id, Title: "", Text: ""}
	testService := NewPostService(nil)

	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())
}
