package ClientService

import (
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
)

type service struct {
	usecase domain.UserUseCase
}

// New returns contract implementation of ProductService
func New(usecase domain.UserUseCase) domain.UserService {
	return &service{
		usecase: usecase,
	}
}
