package entity

import "github.com/SakataAtsuki/e-architecture/pkg/proto/api"

type User struct {
	ID   string
	Name string
}

func (u *User) Proto() *api.User {
	return &api.User{
		Id:   u.ID,
		Name: u.Name,
	}
}
