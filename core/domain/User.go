package domain

import (
	"net/http"

	"github.com/boooscaaa/clean-go/core/dto"
	"github.com/graphql-go/graphql"
)

// Product is entity of table product database column
type User struct {
	ID          int32   `json:"id"`
	Phone       string  `json:"phone"`
	Image       string `json:"image"`
	Password 	string  `json:"password"`
}

type PoductGraphQLService interface {
	Fetch(params graphql.ResolveParams) (interface{}, error)
}

// ProductService is a contract of http adapter layer
type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type UserUseCase interface {
	Create(productRequest *dto.CreateProductRequest) (*User, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
}

// ProductRepository is a contract of database connection adapter layer
type UserRepository interface {
	Create(productRequest *dto.CreateProductRequest) (*User, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
}
