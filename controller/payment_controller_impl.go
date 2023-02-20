package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tes-synapsis/helper"
	"tes-synapsis/model/api"
	"tes-synapsis/model/api/payment"
	"tes-synapsis/service"
)

type PaymentControllerImpl struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &PaymentControllerImpl{PaymentService: paymentService}
}

func (controller *PaymentControllerImpl) GetNotification(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentNotificationRequest := payment.PaymentNotifikastionRequest{}
	helper.ReadFromRequestBody(request, &paymentNotificationRequest)

	processPayment := controller.PaymentService.ProcessPayment(request.Context(), paymentNotificationRequest)

	apiResponse := api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   processPayment,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
