package tasks_postgres_repository

import core_repository_postgres_pool "github.com/zinovev-dm/golang-todoapp/internal/core/repository/postgres/pool"

type TasksRepository struct {
	pool core_repository_postgres_pool.Pool
}

func NewTasksRepository(pool core_repository_postgres_pool.Pool) *TasksRepository {
	return &TasksRepository{
		pool: pool,
	}
}
