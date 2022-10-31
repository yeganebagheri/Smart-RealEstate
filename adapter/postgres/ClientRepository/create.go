package ClientRepository

import (
	"context"

	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

func (repository repository) Create(
	userRequest *dto.CreateUserRequest,
) (*domain.User, error) {
	ctx := context.Background()
	user := domain.User{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO client (username, phone, password) VALUES ($1, $2, $3) returning id,username, phone, password",
		userRequest.Username,
		userRequest.PhonNo,
		userRequest.Password,
	).Scan(
		&user.Id,
		&user.Username,
		&user.PhonNo,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
