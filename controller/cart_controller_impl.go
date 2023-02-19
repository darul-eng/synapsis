package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"tes-synapsis/helper"
	"tes-synapsis/model/api"
	"tes-synapsis/model/api/cart"
	"tes-synapsis/service"
)

type CartControllerImpl struct {
	CartService service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return &CartControllerImpl{CartService: cartService}
}

func (controller *CartControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartCreateRequest := cart.CartCreateRequest{}
	helper.ReadFromRequestBody(request, &cartCreateRequest)

	controller.CartService.Create(request.Context(), cartCreateRequest)

	apiResponse := api.ApiResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CartControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartResponses := controller.CartService.FindAll(request.Context())

	apiResponse := api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponses,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CartControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartId := params.ByName("cartId")
	id, err := strconv.Atoi(cartId)
	helper.PanicIfError(err)

	controller.CartService.Delete(request.Context(), id)

	apiResponse := api.ApiResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
