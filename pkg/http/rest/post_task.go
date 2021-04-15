package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wmattei/gotodo/pkg/services"
	"github.com/wmattei/gotodo/pkg/task"
)

func PostTask(service *services.TaskService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req PostTaskReq
		err := c.BodyParser(&req)
		if err != nil {
			// TODO set http error codes and messages
			c.Status(500)
			return err
		}

		task := req.ToTask()
		fmt.Printf("task %v", task)

		err = service.Save(c.Context(), task)

		if err != nil {
			// TODO set http error codes and messages
			c.Status(500)
			return err
		}

		c.Status(fiber.StatusCreated)
		return nil
	}

}

func (req PostTaskReq) ToTask() *task.Task {
	title := req.Title
	description := req.Description

	return &task.Task{
		Title:       title,
		Description: description,
	}
}

type PostTaskReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
