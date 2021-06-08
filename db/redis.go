package db

import (
	"context"
	"net/http"
	"encoding/json"

	"github.com/go-redis/redis/v8"

	"github.com/Qalifah/movieReviewApi/models"
)

type MovieRepository struct {
	Client *redis.Client
}

func NewMovieRepository(client *redis.Client) *MovieRepository {
	return &MovieRepository{Client: client}
}

func(r *MovieRepository) Get(ctx context.Context, episodeID string) (*models.Movie, error) {
	var movie *models.Movie
	iter := r.Client.Scan(ctx, 0, episodeID, 1).Iterator()
	err := r.Client.Get(ctx, iter.Val()).Scan(&movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func(r *MovieRepository) GetAll(ctx context.Context)  (models.Movies, error) {
	var movies models.Movies
	iter := r.Client.Scan(ctx, 0, "*", 0).Iterator()

	for iter.Next(ctx) {
		var movie *models.Movie
		err := r.Client.Get(ctx, iter.Val()).Scan(&movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository)GetCharactersInMovie(ctx context.Context, characterUrls []string) (models.Characters, error) {
	// declare a character and a collection of characters entity
	var characters models.Characters
	var character *models.Character

	for _, item := range characterUrls {

		// check if the key exist in the redis store
		result := r.Client.Get(ctx, item)

		// if key doesn't exist, fetches it value from the api and stores it
		if result.Err() == redis.Nil {
			resp, err := http.Get(item)
			if err != nil {
				return nil, err
			}

			var character *models.Character
			json.NewDecoder(resp.Body).Decode(&character)

			err = r.Client.Set(ctx, item, character, 0).Err()
			if err != nil {
				return nil, err
			}
			characters = append(characters, character)
		}
		result.Scan(&character)
		characters = append(characters, character)
	}
	return characters, nil
}

