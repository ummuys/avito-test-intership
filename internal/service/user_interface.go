package service

import "context"

type UserService interface {
	SetState(ctx context.Context, userID string, state bool) (string, string, error)
	Get()
}
