package repository

import "context"

type AdminDB interface {
	CreateUser(ctx context.Context, username string, hashPassword string, role string) (err error)
	ValidateRole(ctx context.Context, role string) error
}
