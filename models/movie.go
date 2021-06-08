package models

import (
	"context"
	"time"
)

// Movie --
type Movie struct {
	EpisodeID string `json:"-"`
	Title string `json:"title"`
	OpeningCrawl string `json:"opening_crawl"`
	ReleaseDate		string	`json:"-"`
	Characters		[]string	`json:"-"`
	URL				string 		`json:"-"`
	NumberOfComments int `json:"number_of_comments"`
}

// Movies --
type Movies []*Movie 

// MovieRepository --
type MovieRepository interface {
	Get(ctx context.Context, episodeID string) (*Movie, error)
	GetAll(ctx context.Context) (Movies, error)
	GetCharactersInMovie(ctx context.Context, characterUrls []string) (Characters, error)
}

type Character struct {
	Name string		`json:"name"`
	Gender	string	`json:"gender"`
	Height	string	`json:"height"`
	Films 	[]string	`json:"-"`
}

type Characters []*Character

// Comment --
type Comment struct {
	ID   string  `gorm:"size:500;primary_key;" json:"id"`
	MovieID	string`gorm:"size:500;not null;" json:"movie_id"`
	Content   string `gorm:"size:500;not null;" json:"content"`
	CommenterIPAddress   string	`gorm:"size:25;" json:"commenter_ip_address"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

// Comments --
type Comments []*Comment

func NewComment( movieID string, content, commenterIPAddress string) *Comment {
	return &Comment{
		MovieID: movieID,
		Content: content,
		CommenterIPAddress: commenterIPAddress,
		CreatedAt: time.Now(),
	}
}

type CommentRepository interface {
	Create(newComment *Comment) (*Comment, error)
	GetAll(movieID string) (Comments, error)
}