package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"tes-synapsis/exception"
	"tes-synapsis/helper"
	"tes-synapsis/model/api/auth"
	"tes-synapsis/model/domain"
	"tes-synapsis/repository"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	JWT            JWTService
}

func NewAuthService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate, JWT JWTService) AuthService {
	return &AuthServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate, JWT: JWT}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request auth.RegisterRequest) auth.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	User := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(password),
	}

	user := service.UserRepository.Save(ctx, tx, User)

	token, err := service.JWT.GenerateToken(user.Username)
	helper.PanicIfError(err)

	return helper.ToAuthResponse(user, token)
}

func (service *AuthServiceImpl) Login(ctx context.Context, request auth.LoginRequest) auth.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.Find(ctx, tx, request.Username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewCredentialNotMatchError(err.Error()))
	}

	token, err := service.JWT.GenerateToken(user.Username)
	helper.PanicIfError(err)

	return helper.ToAuthResponse(user, token)
}
