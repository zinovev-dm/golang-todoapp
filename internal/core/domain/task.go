package domain

import (
	"fmt"
	"time"

	core_errors "github.com/zinovev-dm/golang-todoapp/internal/core/errors"
)

type Task struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	AuthorUserID int
	CreatedAt    time.Time
	CompletedAt  *time.Time
}

type TaskPath struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTask(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	authorUserID int,
	createdAt time.Time,
	completedAt *time.Time,
) Task {
	return Task{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		Completed:    completed,
		AuthorUserID: authorUserID,
		CreatedAt:    createdAt,
		CompletedAt:  completedAt,
	}
}

func NewTaskUninitialized(
	title string,
	description *string,
	authorUserID int,
) Task {
	return NewTask(
		UninitializedID,
		UninitializedVersion,
		title,
		description,
		false,
		authorUserID,
		time.Now(),
		nil,
	)
}

func (t Task) Validate() error {
	titleLength := len([]rune(t.Title))

	if titleLength < 1 || titleLength > 100 {
		return fmt.Errorf(
			"invalid title length: %d: %w",
			titleLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if t.Description != nil {
		descriptionLength := len([]rune(*t.Description))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf(
				"invalid description length: %d: %w",
				descriptionLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf(
				"completed_at is empty, but task completed: %w",
				core_errors.ErrInvalidArgument,
			)
		}
		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf(
				"completed_at before created_at: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf(
				"completed_at is not empty, but task not completed: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}
	return nil
}

func NewTaskPath(title Nullable[string], desctiption Nullable[string], completed Nullable[bool]) TaskPath {
	return TaskPath{
		Title:       title,
		Description: desctiption,
		Completed:   completed,
	}
}

func (p *TaskPath) Validate() error {
	if p.Title.Set {
		if p.Title.Value == nil {
			return fmt.Errorf("Title can't be NULL")
		}

		titleLen := len([]rune(*p.Title.Value))

		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("Title lenght must been between 1 and 100")
		}
	}
	if p.Description.Set {
		if p.Description.Value != nil {
			descriptionLen := len([]rune(*p.Description.Value))

			if descriptionLen < 1 || descriptionLen > 1000 {
				return fmt.Errorf("Description lenght must been between 1 and 1000")
			}
		}
	}
	if p.Completed.Set {
		if p.Completed.Value == nil {
			return fmt.Errorf("Completed can't be NULL")
		}
	}

	return nil
}

func (t *Task) ApplyPath(path TaskPath) error {
	if err := path.Validate(); err != nil {
		return fmt.Errorf("validate path task: %w", err)
	}

	tmp := *t

	if path.Title.Set {
		tmp.Title = *path.Title.Value
	}
	if path.Description.Set {
		tmp.Description = path.Description.Value
	}
	if path.Completed.Set {
		tmp.Completed = *path.Completed.Value
		if tmp.Completed {
			compleatedAt := time.Now()
			tmp.CompletedAt = &compleatedAt
		} else {
			tmp.CompletedAt = nil
		}
	}
	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched task: %w", err)
	}
	*t = tmp

	return nil
}
