package di

import (
	"github.com/yeganebagheri/Smart-RealEstate/adapter/http/rest/ClientService"
	"github.com/yeganebagheri/Smart-RealEstate/adapter/postgres"
	"github.com/yeganebagheri/Smart-RealEstate/adapter/postgres/ClientRepository"
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/usecase/userusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigUserDI(conn postgres.PoolInterface) domain.UserService {
	userRepository := ClientRepository.New(conn)
	userUseCase := userusecase.New(userRepository)
	userService := ClientService.New(userUseCase)

	return userService
}

// ConfigProductGraphQLDI return a PoductGraphQLService abstraction with dependency injection configuration
