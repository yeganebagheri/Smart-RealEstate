package userusecase

import (
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

func (usecase usecase) Login(userRequest *dto.LoginUserRequest) (*domain.User, error) {
	user, err := usecase.repository.Login(userRequest)

	if err != nil {
		return nil, err
	}

	return user, nil
}
