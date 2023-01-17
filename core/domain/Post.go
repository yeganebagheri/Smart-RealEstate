package domain

import (
	"net/http"

	"github.com/jackc/pgtype"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

// Product is entity of table product database column
type Post struct {
	Id          int         `json:"id"`
	Date        pgtype.Date `json:"date"`
	Price       int         `json:"Price"`
	Location    string      `json:"Location"`
	Image       string      `json:"Image"`
	Title       string      `json:"Title"`
	Description string      `json:"Description"`
}

// ProductService is a contract of http adapter layer
type PostService interface {
	CreatePost(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	//Fetch(response http.ResponseWriter, request *http.Request)
	//GET(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type PostUseCase interface {
	CreatePost(postRequest *dto.CreatePostRequest) (*Post, error)
	Get(postRequest *dto.GetPostRequest) ([]*Post, error)
	//Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
	//GET(userRequest *dto.CreateUserRequest) (*User, error)
}

// ProductRepository is a contract of database connection adapter layer
type PostRepository interface {
	CreatePost(postRequest *dto.CreatePostRequest) (*Post, error)
	Get(postRequest *dto.GetPostRequest) ([]*Post, error)
	//Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
	//GET(response http.ResponseWriter, request *http.Request)
}
