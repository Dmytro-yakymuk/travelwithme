package service

import (
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
)

type CommentService struct {
	commentRepository repository.Comment
}

func NewCommentService(commentRepository repository.Comment) *CommentService {
	return &CommentService{commentRepository: commentRepository}
}

func (s *CommentService) Get(urlQuery map[string][]string) ([]models.Comment, error) {
	return s.commentRepository.Get(urlQuery)
}
func (s *CommentService) GetAllCommentsByTourId(tourId int) ([]models.Comment, error) {
	return s.commentRepository.GetAllCommentsByTourId(tourId)
}
func (s *CommentService) GetOneByID(id int) ([]models.Comment, error) {
	return s.commentRepository.GetOneByID(id)
}
func (s *CommentService) Create(comment models.Comment) error {
	return s.commentRepository.Create(comment)
}
func (s *CommentService) Update(id int, comment models.Comment) error {
	return s.commentRepository.Update(id, comment)
}
func (s *CommentService) Delete(id int) error {
	return s.commentRepository.Delete(id)
}
