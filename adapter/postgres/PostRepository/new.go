package PostRepository

import (
	"github.com/yeganebagheri/Smart-RealEstate/adapter/postgres"
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.PostRepository {
	return &repository{
		db: db,
	}
}
