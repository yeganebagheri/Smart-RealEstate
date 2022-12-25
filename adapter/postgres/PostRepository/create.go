package PostRepository

import (
	"context"
	"time"

	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

func (repository repository) CreatePost(
	postRequest *dto.CreatePostRequest,
) (*domain.Post, error) {
	ctx := context.Background()
	post := domain.Post{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO post (registrationtime, price, location, image, title, description) VALUES ($1, $2, $3, $4, $5, $6) returning *",
		time.Now(),
		postRequest.Price,
		postRequest.Location,
		postRequest.Image,
		postRequest.Title,
		postRequest.Description,
	).Scan(
		&post.Id,
		&post.Date,
		&post.Price,
		&post.Location,
		&post.Image,
		&post.Title,
		&post.Description,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
