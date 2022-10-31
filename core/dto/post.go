package dto

import (
	"encoding/json"
	"io"
)

// CreateProductRequest is an representation request body to create a new Product
type CreatePostRequest struct {
	Price       string `json:"Price"`
	Location    string `json:"Location"`
	Image       string `json:"Image"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type GetPostRequest struct {
	Location string `json:"Location"`
	Price    string `json:"Price"`
	Title    string `json:"Title"`
}

// FromJSONCreateProductRequest converts json body request to a CreateProductRequest struct
func FromJSONCreatePostRequest(body io.Reader) (*CreatePostRequest, error) {
	createPostRequest := CreatePostRequest{}
	if err := json.NewDecoder(body).Decode(&createPostRequest); err != nil {
		return nil, err
	}

	return &createPostRequest, nil
}

func FromJSONGetPostRequest(body io.Reader) (*GetPostRequest, error) {
	getPostRequest := GetPostRequest{}
	if err := json.NewDecoder(body).Decode(&getPostRequest); err != nil {
		return nil, err
	}

	return &getPostRequest, nil
}

// func Create(
// 	string price, string location
// ) (*domain.Post, error) {

// }
