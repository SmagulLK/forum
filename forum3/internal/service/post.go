package service

import (
	"forum/internal/models"
	"forum/internal/storage"
)

type ServicePostIR interface {
	CreatePost(post models.Post) error
	GetPostId(id int) (models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetCategories() ([]string, error)
	GetAllPostsByCategories(category string) ([]models.Post, error)
	GetMyPost(int) ([]models.Post, error)
	GetMyLikePost(int) ([]models.Post, error)
}

type PostService struct {
	storage storage.PostIR
}

func NewPostService(postIR storage.PostIR) ServicePostIR {
	return &PostService{
		storage: postIR,
	}
}

func (p *PostService) CreatePost(post models.Post) error {
	return p.storage.CreatePost(post)
}

func (p *PostService) GetPostId(id int) (models.Post, error) {
	return p.storage.GetPostByID(id)
}

func (p *PostService) GetAllPosts() ([]models.Post, error) {
	return p.storage.GetAllPosts()
}

func (p *PostService) GetCategories() ([]string, error) {
	return p.storage.Category()
}

func (p *PostService) GetAllPostsByCategories(category string) ([]models.Post, error) {
	return p.storage.GetAllPostsByCategories(category)
}

func (p *PostService) GetMyPost(id int) ([]models.Post, error) {
	return p.storage.GetMyPost(id)
}
func (p *PostService) GetMyLikePost(id int) ([]models.Post, error) {
	return p.storage.GetMyLikedPost(id)
}
