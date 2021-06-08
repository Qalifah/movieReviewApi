package handler

import (
	"net/http" 

	"github.com/Qalifah/movieReviewApi/models"

	"github.com/gin-gonic/gin"
)


type RequestHandler struct {
	MovieRepo models.MovieRepository
	CommentRepo models.CommentRepository
}

func NewRequestHandler(movieRepo models.MovieRepository, commentRepo models.CommentRepository) *RequestHandler {
	return &RequestHandler{
		MovieRepo: movieRepo,
		CommentRepo: commentRepo,
	}
}

func(h *RequestHandler) AllMovies(c *gin.Context) {
	movies, err := h.MovieRepo.GetAll(c.Copy())
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, movies)
}

func(h *RequestHandler) AddComment(c *gin.Context) {
	movieID := c.Param("movie_id")
	var comment *models.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusUnprocessableEntity, NewErrorResponse("invalid json object"))
     	return
	}
	comment, err := h.CommentRepo.Create(models.NewComment(movieID, comment.Content,  comment.CommenterIPAddress))
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, comment)
}

func(h *RequestHandler) ListComments(c *gin.Context) {
	movieID := c.Param("movie_id")
	
	comments, err := h.CommentRepo.GetAll(movieID)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, comments)
}

func(h *RequestHandler) ListCharacters(c *gin.Context) {
	movieID := c.Param("movie_id")
	movie, err := h.MovieRepo.Get(c.Copy(), movieID)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}
	characters, err := h.MovieRepo.GetCharactersInMovie(c.Copy(), movie.Characters)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, characters)
}
