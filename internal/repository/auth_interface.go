package repository

import "context"

type AuthDB interface {
	CheckCredentials(ctx context.Context, username string) (int64, string, string, error)
}
