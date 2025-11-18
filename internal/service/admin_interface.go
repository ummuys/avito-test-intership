package service

import "context"

type AdminService interface {
	CreateUser(ctx context.Context, username, password, role string) error
}
