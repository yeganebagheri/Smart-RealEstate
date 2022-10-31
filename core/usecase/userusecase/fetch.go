package userusecase

import (
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	users, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return users, nil
}
