package service

import "context"

type AuthService interface {
	CheckCredentials(ctx context.Context, username, password string) (int64, string, error)
}
