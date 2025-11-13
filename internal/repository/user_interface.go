package repository

import (
	"context"
)

type UserDB interface {
	SetUserState(ctx context.Context, userID string, state bool) (string, string, error)
}
