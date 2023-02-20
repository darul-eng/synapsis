package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PaymentController interface {
	GetNotification(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
