package service

import (
	"context"

	"github.com/kahuri1/birthday-notifier/internal/entity"
)

type UserRepository interface {
	Find(ctx context.Context, userID int) (entity.User, error)
	Create(ctx context.Context, user entity.User) error
}

type User struct {
	repo UserRepository
}

func New(repo UserRepository) *User {
	return &User{repo: repo}
}

func (u *User) Create(ctx context.Context, user entity.User) error {
	return u.repo.Create(ctx, user)
}
