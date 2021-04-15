package services

import (
	"context"

	"github.com/wmattei/gotodo/pkg/postgres"
	"github.com/wmattei/gotodo/pkg/task"
)

type TaskService struct {
	storage *postgres.TaskStorage
}

func NewTaskService(storage *postgres.TaskStorage) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) Save(ctx context.Context, task *task.Task) error {
	return s.storage.Save(ctx, task)
}
