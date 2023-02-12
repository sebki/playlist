package data

import "context"

// User holds all relevant userinformation
type User struct {
	Uid      string `json:"uid,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserService interface {
	FindUserByID(ctx context.Context, id string) (*User, error)
	FindUsers(ctx context.Context, filter UserFilter) (*[]User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, id string, upd UserUpdate) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

type UserFilter struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
