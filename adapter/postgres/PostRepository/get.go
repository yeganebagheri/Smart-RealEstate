package PostRepository

import (
	"context"
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
	
)

func (repository repository) Get(
	getPostRequest *dto.GetPostRequest,
) (*domain.Post, error) {
	ctx := context.Background()
	post := domain.Post{}

	err := repository.db.QueryRow(
		ctx,
		"SELECT * FROM public.post where (price = $1 OR $1 IS NULL) and (location = $2 OR $2 IS NULL) "+
			"and (title = $3 OR $3 IS NULL)",
		getPostRequest.Price,
		getPostRequest.Location,
		getPostRequest.Title,
	).Scan(
		&post.Id,
		&post.RegistrationTime,
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
