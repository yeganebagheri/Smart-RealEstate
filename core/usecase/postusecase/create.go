package postusecase

import (
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

func (usecase usecase) CreatePost(postRequest *dto.CreatePostRequest) (*domain.Post, error) {
	post, err := usecase.repository.CreatePost(postRequest)

	if err != nil {
		return nil, err
	}

	return post, nil
}
