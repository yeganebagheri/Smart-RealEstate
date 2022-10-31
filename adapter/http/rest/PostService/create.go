package PostService

import (
	"encoding/json"
	"net/http"

	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

// @Summary Create new product
// @Description Create new product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductRequest true "product"
// @Success 200 {object} domain.Product
// @Router /product [post]
func (service service) CreatePost(response http.ResponseWriter, request *http.Request) {
	postRequest, err := dto.FromJSONCreatePostRequest(request.Body)

	if err != nil {
		response.WriteHeader(250)
		response.Write([]byte(err.Error()))
		return
	}

	post, err := service.usecase.CreatePost(postRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(post)
}
