package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Qalifah/movieReviewApi/db"
	"github.com/Qalifah/movieReviewApi/handler"
	"github.com/Qalifah/movieReviewApi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
	gormHandler, err := gorm.Open("postgres", os.Getenv("DB_ADDRESS"))
	if err != nil {
		log.Fatalln(err)
	}

	requestHandler := handler.NewRequestHandler(
		db.NewMovieRepository(redis.NewClient(&redis.Options{Addr: "localhost:6379"})),
		db.NewCommentRepository(gormHandler),
	)

	router := gin.Default()
	router.GET("/movies", requestHandler.AllMovies)
	router.POST("/movies/:movie_id/comment", requestHandler.AddComment)
	router.GET("/movies/:movie_id/comment", requestHandler.ListComments)
	router.GET("/movies/:movie_id/characters", requestHandler.ListCharacters)

	log.Fatal(router.Run(":8080"))
}


func PopulateRedis(ctx context.Context, data models.Movies, client *redis.Client) (*redis.Client, error) {
	for _, item := range data {
		err := client.Set(ctx, item.Title, item, 0).Err()
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func GetDataFromAPI() (models.Movies, error) {
	resp, err := http.Get("https://swapi.dev/api/films/")
	if err != nil {
		return nil, err
	}
	var movies models.Movies
	if err = json.NewDecoder(resp.Body).Decode(&movies); err != nil {
		return nil, err
	}
	return movies, nil
}