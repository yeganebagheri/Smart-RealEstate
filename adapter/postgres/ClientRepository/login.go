package ClientRepository

import (
	"context"

	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

func (repository repository) Login(
	userRequest *dto.LoginUserRequest,
) (*domain.User, error) {
	ctx := context.Background()
	user := domain.User{}

	err := repository.db.QueryRow(
		ctx,
		"SELECT * FROM public.client where phone = $1 and password = $2",
		userRequest.Phone,
		userRequest.Password,
	).Scan(
		&user.Username,
		&user.Phone,
		&user.Password,
		&user.Id,
	)

	if err != nil {
		
		return nil, err
	}

	return &user, nil
}
