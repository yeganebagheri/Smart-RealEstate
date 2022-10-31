package PostService

import (
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
)

type service struct {
	usecase domain.PostUseCase
}

// New returns contract implementation of ProductService
func New(usecase domain.PostUseCase) domain.PostService {
	return &service{
		usecase: usecase,
	}
}
