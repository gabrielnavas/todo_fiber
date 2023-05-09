package controller

import (
	"backend/task/models"
	service "backend/task/services"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (tc *TaskController) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var task = &models.Task{}
		err := ctx.BodyParser(task)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		task, err = tc.taskService.Create(task)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
			"task": fiber.Map{
				"id":      task.Id,
				"message": task.Description,
			},
		})
		return nil
	}
}

func (tc *TaskController) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (tc *TaskController) UpdateStatus() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (tc *TaskController) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (tc *TaskController) Get() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (tc *TaskController) GetAll() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	}
}
