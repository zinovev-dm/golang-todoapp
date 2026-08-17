package tasks_postgres_repository

import (
	"time"

	"github.com/zinovev-dm/golang-todoapp/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	AuthorUserID int
	CreatedAt    time.Time
	CompletedAt  *time.Time
}

func taskDomainsFromModels(models []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(models))
	for i, t := range models {
		domains[i] = taskDomainFromModel(t)
	}
	return domains
}

func taskDomainFromModel(model TaskModel) domain.Task {
	return domain.NewTask(
		model.ID,
		model.Version,
		model.Title,
		model.Description,
		model.Completed,
		model.AuthorUserID,
		model.CreatedAt,
		model.CompletedAt,
	)
}
