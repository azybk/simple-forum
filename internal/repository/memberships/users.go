package memberships

import (
	"context"

	"github.com/azybk/simple-forum/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := "SELECT id, email, username FROM users WHERE email=$1 OR username=$2"
	row := r.db.QueryRowContext(ctx, query, email, username)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Username)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := "INSERT INTO users(email, password, created_at, updated_at, created_by, updated_by, username)VALUES($1, $2, $3, $4, $5, $6, $7)"

	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy, model.Username)

	if err != nil {
		return err
	}

	return nil
}
