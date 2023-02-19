package middleware

import (
	"net/http"
	"tes-synapsis/helper"
	"tes-synapsis/model/api"
	"tes-synapsis/service"
)

type AuthMiddleware struct {
	Handler    http.Handler
	jwtService service.JWTService
}

func NewAuthMiddleware(handler http.Handler, jwtService service.JWTService) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, jwtService: jwtService}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if request.RequestURI == "/api/login" || request.RequestURI == "/api/register" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		token, err := middleware.jwtService.VerifyToken(request.Header.Get("API-KEY"))
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			apiResponse := api.ApiResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}
			helper.WriteToResponseBody(writer, apiResponse)
		} else {
			if token.Valid {
				middleware.Handler.ServeHTTP(writer, request)
			} else {
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusUnauthorized)

				apiResponse := api.ApiResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
				}
				helper.WriteToResponseBody(writer, apiResponse)
			}
		}

	}
}
