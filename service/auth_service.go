package service

import (
	"context"
	user "tes-synapsis/model/api/Auth"
)

type AuthService interface {
	Register(ctx context.Context, request user.RegisterRequest) user.AuthResponse
	Login(ctx context.Context, request user.LoginRequest) user.AuthResponse
}
