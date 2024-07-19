package service

import (
	"context"

	"github.com/kahuri1/birthday-notifier/internal/entity"
)

type SubscriptionRepository interface {
	Find(ctx context.Context, userID int) ([]entity.Subscription, error)
}

type Subscription struct {
	repo SubscriptionRepository
}

func NewSubscription(repo SubscriptionRepository) *Subscription {
	return &Subscription{repo: repo}
}

func (s *Subscription) Find(ctx context.Context, userID int) ([]entity.Subscription, error) {
	return s.repo.Find(ctx, userID)
}
