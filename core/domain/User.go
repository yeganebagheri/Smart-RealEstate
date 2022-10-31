package domain

import (
	"net/http"

	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

// Product is entity of table product database column
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	PhonNo   string `json:"PhonNo"`
	Password string `json:"Password"`
}

// ProductService is a contract of http adapter layer
type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Login(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
	//GET(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type UserUseCase interface {
	Create(userRequest *dto.CreateUserRequest) (*User, error)
	Login(userRequest *dto.LoginUserRequest) (*User, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
	//GET(userRequest *dto.CreateUserRequest) (*User, error)
}

// ProductRepository is a contract of database connection adapter layer
type UserRepository interface {
	Create(userRequest *dto.CreateUserRequest) (*User, error)
	Login(userRequest *dto.LoginUserRequest) (*User, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
	//GET(response http.ResponseWriter, request *http.Request)
}
