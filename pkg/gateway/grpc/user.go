package grpc

import (
	"context"

	"github.com/SakataAtsuki/e-architecture/pkg/entity"
	"github.com/SakataAtsuki/e-architecture/pkg/proto/api"
	"github.com/SakataAtsuki/e-architecture/pkg/usecase"
	"github.com/SakataAtsuki/e-architecture/pkg/util/errcode"
)

func (s *Service) CreateUser(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	req := &usecase.CreateUserRequest{
		User: &entity.User{
			ID:   in.User.Id,
			Name: in.User.Name,
		},
	}
	resp, err := s.uc.CreateUser(ctx, req)
	if err != nil {
		return nil, errcode.New(err)
	}
	return &api.CreateUserResponse{User: resp.User.Proto()}, nil
}

func (s *Service) GetUser(ctx context.Context, in *api.GetUserRequest) (*api.GetUserResponse, error) {
	req := &usecase.GetUserRequest{
		ID: in.Id,
	}
	resp, err := s.uc.GetUser(ctx, req)
	if err != nil {
		return nil, errcode.New(err)
	}
	return &api.GetUserResponse{User: resp.User.Proto()}, nil
}
