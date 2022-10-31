package userusecase

import "github.com/yeganebagheri/Smart-RealEstate/core/domain"

type usecase struct {
	repository domain.UserRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.UserRepository) domain.UserUseCase {
	return &usecase{
		repository: repository,
	}
}
