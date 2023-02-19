package helper

import (
	user "tes-synapsis/model/api/Auth"
	"tes-synapsis/model/domain"
)

func ToAuthResponse(data domain.User, token string) user.AuthResponse {
	return user.AuthResponse{
		Id:    data.Id,
		Name:  data.Username,
		Token: token,
	}
}
