package PostService

import (
	"encoding/json"
	"log"
	"net/http"
	"bytes"
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
func (service service) Get(response http.ResponseWriter, request *http.Request) {
	price := request.URL.Query().Get("Price")
	location := request.URL.Query().Get("Location")
	title := request.URL.Query().Get("Title")
	//post := domain.Post{}
	// u := &post{
	// 	Price:    price,
	// 	Location: location,
	// 	Title:    title,
	// }
	var getPostRequest = new(dto.GetPostRequest)
	getPostRequest.Price = price
	getPostRequest.Location = location
	getPostRequest.Title = title

	data, err := json.Marshal(getPostRequest)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(data)


	getPostRequest1, err := dto.FromJSONGetPostRequest(reader)

	if err != nil {
		response.WriteHeader(250)
		response.Write([]byte(err.Error()))
		return
	}

	user, err := service.usecase.Get(getPostRequest1)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(user)
}
