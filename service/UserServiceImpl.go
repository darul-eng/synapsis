package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"tes-synapsis/helper"
	user "tes-synapsis/model/api/Auth"
	"tes-synapsis/model/domain"
	"tes-synapsis/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	JWT            JWTService
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate, JWT JWTService) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate, JWT: JWT}
}

func (service *UserServiceImpl) Register(ctx context.Context, request user.RegisterRequest) user.AuthResponse {
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

func (service *UserServiceImpl) Login(ctx context.Context, request user.LoginRequest) user.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.Find(ctx, tx, request.Username)
	helper.PanicIfError(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	helper.PanicIfError(err)

	token, err := service.JWT.GenerateToken(user.Username)
	helper.PanicIfError(err)

	return helper.ToAuthResponse(user, token)
}