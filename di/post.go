package di

import (
	"github.com/yeganebagheri/Smart-RealEstate/adapter/http/rest/PostService"
	"github.com/yeganebagheri/Smart-RealEstate/adapter/postgres"
	"github.com/yeganebagheri/Smart-RealEstate/adapter/postgres/PostRepository"
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/usecase/postusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigPostDI(conn postgres.PoolInterface) domain.PostService {
	postRepository := PostRepository.New(conn)
	postUseCase := postusecase.New(postRepository)
	postService := PostService.New(postUseCase)

	return postService
}

// ConfigProductGraphQLDI return a PoductGraphQLService abstraction with dependency injection configuration
