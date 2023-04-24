package service

import (
	"forum/internal/models"
	"forum/internal/storage"
)

type CommentServiceIR interface {
	GetCommentsByIdPost(id int) ([]models.Comment, error)
	CreateComment(id int, username, commentText string) error
}

type CommentService struct {
	storage storage.CommentIR
}

func newCommentServ(storage storage.CommentIR) CommentServiceIR {
	return &CommentService{
		storage: storage,
	}
}

func (c *CommentService) GetCommentsByIdPost(id int) ([]models.Comment, error) {
	return c.storage.GetCommentsByIdPost(id)
}

func (c *CommentService) CreateComment(id int, username, commentText string) error {
	return c.storage.CreateComment(id, username, commentText)
}
