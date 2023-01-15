package repository

import (
	"context"

	"github.com/SakataAtsuki/e-architecture/pkg/entity"
)

type Database struct {
	User User
}

type User interface {
	Create(ctx context.Context, v *entity.User) (*entity.User, error)
	Get(ctx context.Context, id string) (*entity.User, error)
	// List(ctx context.Context, params *ListUsersParams) (entity.Users, error)
	// Update(ctx context.Context, id string, update func(*entity.User) bool) (*entity.User, error)
	// Delete(ctx context.Context, id string) error
}

// type ListUsersParams struct {
// 	Name    string
// 	OrderBy string
// 	Limit   int
// }
