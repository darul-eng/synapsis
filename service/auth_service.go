package service

import (
	"context"
	"tes-synapsis/model/api/auth"
)

type AuthService interface {
	Register(ctx context.Context, request auth.RegisterRequest) auth.AuthResponse
	Login(ctx context.Context, request auth.LoginRequest) auth.AuthResponse
}
