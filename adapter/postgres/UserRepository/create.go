package productrepository

import (
	"context"

	"github.com/boooscaaa/clean-go/core/domain"
	"github.com/boooscaaa/clean-go/core/dto"
)

func (repository repository) Create(
	productRequest *dto.CreateProductRequest,
) (*domain.User, error) {
	ctx := context.Background()
	product := domain.User{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO product (name, price, description) VALUES ($1, $2, $3) returning *",
		productRequest.Name,
		productRequest.Price,
		productRequest.Description,
	).Scan(
		&product.ID,
		&product.Phone,
		&product.Image,
		&product.Password,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
