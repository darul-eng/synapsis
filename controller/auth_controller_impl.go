package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tes-synapsis/helper"
	"tes-synapsis/model/api"
	user "tes-synapsis/model/api/Auth"
	"tes-synapsis/service"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{AuthService: authService}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := user.RegisterRequest{}
	helper.ReadFromRequestBody(request, &registerRequest)

	registerResponse := controller.AuthService.Register(request.Context(), registerRequest)

	apiResponse := api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   registerResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := user.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse := controller.AuthService.Login(request.Context(), loginRequest)

	apiResponse := api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
