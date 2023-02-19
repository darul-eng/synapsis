package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tes-synapsis/helper"
	"tes-synapsis/model/api"
	"tes-synapsis/model/api/transaction"
	"tes-synapsis/service"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{TransactionService: transactionService}
}

func (controller *TransactionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transactionCreateRequest := transaction.TransactionCreateRequest{}
	helper.ReadFromRequestBody(request, &transactionCreateRequest)

	transactionResponse := controller.TransactionService.Create(request.Context(), transactionCreateRequest)

	apiResponse := api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
