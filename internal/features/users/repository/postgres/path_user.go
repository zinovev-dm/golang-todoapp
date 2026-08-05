package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/zinovev-dm/golang-todoapp/internal/core/domain"
	core_errors "github.com/zinovev-dm/golang-todoapp/internal/core/errors"
)

func (r *UsersRepository) PathUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE todoapp.users
		SET
			full_name = $2,
			phone_number = $3,
			version = version + 1
		WHERE id = $1
		AND version = $4
		RETURNING
			id,
			version,
			full_name,
			phone_number
	`

	row := r.pool.QueryRow(ctx, query, id, user.FullName, user.PhoneNumber, user.Version)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id=%d concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return domain.NewUser(userModel.ID, userModel.Version, userModel.FullName, userModel.PhoneNumber), nil
}
