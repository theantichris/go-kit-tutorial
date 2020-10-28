package account

import "context"

// User holds the information for a user.
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Repository defines the interface for database interaction.``
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}
