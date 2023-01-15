package entity

type User struct {
	ID   string
	Name string
	// Gender    api.Gender
	// UpdatedAt int64
}

// type Users []*User

// func (u *User) Proto() (*api.User, error) {
// 	return &api.User{
// 		Id:        u.ID,
// 		Name:      u.Name,
// 		Gender:    u.Gender,
// 		UpdatedAt: u.UpdatedAt,
// 	}, nil
// }
