package userusecase

import (
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

func (usecase usecase) Create(userRequest *dto.CreateUserRequest) (*domain.User, error) {
	user, err := usecase.repository.Create(userRequest)

	if err != nil {
		return nil, err
	}

	return user, nil
}
