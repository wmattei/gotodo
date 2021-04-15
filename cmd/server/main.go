package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/wmattei/gotodo/pkg/http/rest"
	"github.com/wmattei/gotodo/pkg/postgres"
	"github.com/wmattei/gotodo/pkg/services"
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background() // TODO research about context
	service, err := createTaskService(ctx)

	if err != nil {
		fmt.Print("Failed to create service")
		return err
	}

	app := rest.Handler(service)

	// TODO get from config
	err = app.Listen(":3000")

	if err != nil {
		fmt.Print("Failed to start server")
		return err
	}

	return nil
}

func createTaskService(ctx context.Context) (*services.TaskService, error) {
	storage, err := createTaskStorage(ctx)

	if err != nil {
		fmt.Print("Failed to create storage")
		return nil, err
	}

	return services.NewTaskService(storage), nil
}

func createTaskStorage(ctx context.Context) (*postgres.TaskStorage, error) {
	// TODO use environment variables
	pgPool, err := pgxpool.Connect(ctx, "postgresql://postgres@localhost:5432/gotodo")

	if err != nil {
		fmt.Print("Failed to connect to the database")
		return nil, err
	}

	return postgres.NewTaskStorage(pgPool), nil

}
