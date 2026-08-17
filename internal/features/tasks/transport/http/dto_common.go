package tasks_transport_http

import (
	"time"

	"github.com/zinovev-dm/golang-todoapp/internal/core/domain"
)

type TasksDTOResponse struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	AuthorUserID int        `json:"author_user_id"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
}

func taskDTOFromDomain(task domain.Task) TasksDTOResponse {
	return TasksDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		AuthorUserID: task.AuthorUserID,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
	}
}

func tasksDTOFromDomain(tasks []domain.Task) []TasksDTOResponse {
	dtos := make([]TasksDTOResponse, len(tasks))
	for i, t := range tasks {
		dtos[i] = taskDTOFromDomain(t)
	}
	return dtos
}
