package domain

import (
	"context"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserUsecase interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, user User) error
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, user User) error
}
