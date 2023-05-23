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

		return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
			"task": fiber.Map{
				"id":          task.Id,
				"description": task.Description,
			},
		})
	}
}

func (tc *TaskController) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		taskId := ctx.Params("task_id", "0")

		var task = &models.Task{}
		err := ctx.BodyParser(task)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		err = tc.taskService.Update(taskId, task)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	}
}

func (tc *TaskController) UpdateStatus() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		taskId := ctx.Params("task_id", "0")

		var task = &models.Task{}
		err := ctx.BodyParser(task)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		err = tc.taskService.UpdateStatus(taskId, task.Status.Name)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	}
}

func (tc *TaskController) GetAllStatus() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tasksStatus, err := tc.taskService.GetAllStatus()
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		tasksStatusFiltred := []fiber.Map{}
		for _, taskStatus := range tasksStatus {
			tasksStatusFiltred = append(tasksStatusFiltred, fiber.Map{
				"id":   taskStatus.Id,
				"name": taskStatus.Name,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"tasks": tasksStatusFiltred,
		})
	}
}

func (tc *TaskController) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		taskId := ctx.Params("task_id", "0")

		err := tc.taskService.Delete(taskId)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	}
}

func (tc *TaskController) Get() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (tc *TaskController) GetAll() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tasks, err := tc.taskService.GetAll()
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err.Error(),
			})
			return nil
		}

		tasksFiltred := []fiber.Map{}
		for _, task := range tasks {
			tasksFiltred = append(tasksFiltred, fiber.Map{
				"id":          task.Id,
				"description": task.Description,
				"createdAt":   task.CreatedAt,
				"updatedAt":   task.UpdateAt,
				"status": fiber.Map{
					"id":   task.Status.Id,
					"name": task.Status.Name,
				},
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"tasks": tasksFiltred,
		})
	}
}
