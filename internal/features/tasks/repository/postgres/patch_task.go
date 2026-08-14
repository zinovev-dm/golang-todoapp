package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/zinovev-dm/golang-todoapp/internal/core/domain"
	core_errors "github.com/zinovev-dm/golang-todoapp/internal/core/errors"
	core_repository_postgres_pool "github.com/zinovev-dm/golang-todoapp/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) PathTask(ctx context.Context, id int, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `
		UPDATE todoapp.tasks
		SET
			title = $2,
			description = $3,
			completed = $4,
			completed_at = $5
			version = version + 1
		WHERE id = $1
		AND version = $6
		RETURNING
			id,
			title,
			description,
			completed,
			created_at,
			completed_at,
			author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		id,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.Version,
	)

	var taskModel TaskModel

	err := row.Scan(
		&taskModel.ID,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserID,
	)

	if err != nil {
		if errors.Is(err, core_repository_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id=%d concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return taskDomainFromModel(taskModel), nil
}
