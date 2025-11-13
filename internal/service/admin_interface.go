package service

import "context"

type AdminService interface {
	CreateUser(ctx context.Context, username, password, role string) error
	UpdateUser(ctx context.Context, userID int64, username, password, role string) error
	DeleteUser(ctx context.Context, username string) error
	GetUser(ctx context.Context, userID int64) error
}
