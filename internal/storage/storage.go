package user

import "context"

type Storage interface {
	Create(ctx context.Context, user User) (string, error)
	FindOne(ctx context.Context, customFilter FindUserByFilter) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}

type FindUserByFilter struct {
	Email  string
	UserID string
}
