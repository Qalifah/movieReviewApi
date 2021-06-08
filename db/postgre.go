package db

import (
	"github.com/jinzhu/gorm"

	"github.com/Qalifah/movieReviewApi/models"
)

type CommentRepository struct {
	Store *gorm.DB
}

func NewCommentRepository(store *gorm.DB) *CommentRepository {
	return &CommentRepository{
		Store: store,
	}
}

func(r *CommentRepository) Create(newComment *models.Comment) (*models.Comment, error) {
	err := r.Store.Create(newComment).Error
	if err != nil {
		return nil, err
	}
	return newComment, nil
}

func(r *CommentRepository) GetAll(movieID string) (models.Comments, error) {
	var comments models.Comments
	err := r.Store.Where("movie_id = ?", movieID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
