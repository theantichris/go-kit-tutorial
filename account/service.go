package account

import "context"

// Service defines the interface for the account service.
type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
