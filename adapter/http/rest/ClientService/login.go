package ClientService

import (
	"encoding/json"
	"net/http"

	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
)

// @Summary Fetch products with server pagination
// @Description Fetch products with server pagination
// @Tags product
// @Accept  json
// @Produce  json
// @Param sort query string true "1,2"
// @Param descending query string true "true,false"
// @Param page query integer true "1"
// @Param itemsPerPage query integer true "10"
// @Param search query string false "product name or any column"
// @Success 200 {object} domain.Pagination
// @Router /product [get]
func (service service) Login(response http.ResponseWriter, request *http.Request) {
	loginRequest, err := dto.FromJSONLoginUserRequest(request.Body)

	if err != nil {
		response.WriteHeader(250)
		response.Write([]byte(err.Error()))
		return
	}

	user, err := service.usecase.Login(loginRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(user)
}
