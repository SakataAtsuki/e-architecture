package usecase

import (
	"context"

	"github.com/SakataAtsuki/e-architecture/pkg/repository"
	"github.com/go-playground/validator"
)

type Usecase interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error)
}

type UsecaseImpl struct {
	validate *validator.Validate
	db       *repository.Database
}

type Config struct {
	DB *repository.Database
}

func New(cfg *Config) *UsecaseImpl {
	return &UsecaseImpl{
		validate: validator.New(),
		db:       cfg.DB,
	}
}
