package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

// ErrRepo is returned when there is an error communicating with the database.
var ErrRepo = errors.New("unable to handle repo request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepo creates and returns a new implementation of a Repository.
func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	if user.Email == "" || user.Password == "" {
		return ErrRepo
	}

	sql := "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)"

	if _, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password); err != nil {
		return err
	}

	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string

	if err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email); err != nil {
		return "", ErrRepo
	}

	return email, nil
}
