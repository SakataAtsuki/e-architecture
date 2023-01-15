package usecase

import (
	"context"

	"github.com/SakataAtsuki/e-architecture/pkg/entity"
	"github.com/SakataAtsuki/e-architecture/pkg/util/errcode"
)

type CreateUserRequest struct {
	User *entity.User
}

type CreateUserResponse struct {
	User *entity.User
}

func (u *UsecaseImpl) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	if err := u.validate.Struct(req); err != nil {
		return nil, errcode.New(err)
	}

	// database
	user, err := u.db.User.Create(ctx, req.User)
	if err != nil {
		return nil, errcode.New(err)
	}
	return &CreateUserResponse{User: user}, nil
}

type GetUserRequest struct {
	ID string
}

type GetUserResponse struct {
	User *entity.User
}

func (u *UsecaseImpl) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	if err := u.validate.Struct(req); err != nil {
		return nil, errcode.New(err)
	}

	// database
	user, err := u.db.User.Get(ctx, req.ID)
	if err != nil {
		return nil, errcode.New(err)
	}
	return &GetUserResponse{User: user}, nil
}
