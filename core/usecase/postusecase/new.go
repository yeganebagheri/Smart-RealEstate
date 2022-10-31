package postusecase

import "github.com/yeganebagheri/Smart-RealEstate/core/domain"

type usecase struct {
	repository domain.PostRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.PostRepository) domain.PostUseCase {
	return &usecase{
		repository: repository,
	}
}
