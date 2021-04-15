package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/wmattei/gotodo/pkg/task"
)

const insertTask = `
	INSERT INTO tasks (title, description, done) VALUES ($1, $2, false);
`

type TaskStorage struct {
	pgPool *pgxpool.Pool
}

func NewTaskStorage(pgPool *pgxpool.Pool) *TaskStorage {
	return &TaskStorage{
		pgPool: pgPool,
	}
}

func (s TaskStorage) Save(ctx context.Context, task *task.Task) error {

	params := []interface{}{
		task.Title,
		task.Description,
	} // TODO search about interface

	fmt.Print(task.Title)

	_, err := s.pgPool.Exec(ctx, insertTask, params...)

	if err != nil {
		fmt.Print("Error inserting task")
		return err
	}

	return nil
}
