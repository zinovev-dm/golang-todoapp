package tasks_service

import (
	"context"
	"fmt"

	"github.com/zinovev-dm/golang-todoapp/internal/core/domain"
)

func (s *TasksService) PathTask(ctx context.Context, id int, path domain.TaskPath) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}
	if err := task.ApplyPath(path); err != nil {
		return domain.Task{}, fmt.Errorf("apply path: %w", err)
	}
	patchedTask, err := s.tasksRepository.PathTask(ctx, id, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}
	return patchedTask, nil
}
