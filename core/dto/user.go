package dto

import (
	"encoding/json"
	"io"
)

// CreateProductRequest is an representation request body to create a new Product
type CreateUserRequest struct {
	Username string `json:"username"`
	Phone    string `json:"Phone"`
	Password string `json:"Password"`
}

type LoginUserRequest struct {
	Phone    string `json:"Phone"`
	Password string `json:"Password"`
}

// FromJSONCreateProductRequest converts json body request to a CreateProductRequest struct
func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequest, error) {
	createUserRequest := CreateUserRequest{}
	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}

	return &createUserRequest, nil
}

func FromJSONLoginUserRequest(body io.Reader) (*LoginUserRequest, error) {
	loginUserRequest := LoginUserRequest{}
	if err := json.NewDecoder(body).Decode(&loginUserRequest); err != nil {
		return nil, err
	}

	return &loginUserRequest, nil
}
