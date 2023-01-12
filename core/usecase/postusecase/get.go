package postusecase

import (
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

func (usecase usecase) Get(getPostRequest *dto.GetPostRequest) (*domain.Post, error) {
	posts, err := usecase.repository.Get(getPostRequest)

	if err != nil {
		return nil, err 
	}

	return posts, nil
}
