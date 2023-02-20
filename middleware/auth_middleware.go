package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
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

	if request.RequestURI == "/api/login" || request.RequestURI == "/api/register" || request.RequestURI == "/api/transaction/notification" {
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
				claims := token.Claims.(jwt.MapClaims)
				ctx := context.WithValue(request.Context(), "user", claims["username"])
				middleware.Handler.ServeHTTP(writer, request.WithContext(ctx))
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
