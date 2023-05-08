package main

import (
	"backend/task/controller"
	repository "backend/task/repositories"
	service "backend/task/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func makeDatabase() *sql.DB {
	var err error

	connStr := "postgres://postgres:password123@localhost:5435/task?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func main() {
	app := fiber.New()

	db := makeDatabase()
	defer db.Close()

	taskStatusRepository := repository.NewTaskStatusRepository(db)
	taskRepository := repository.NewTaskRepository(db, taskStatusRepository)
	taskService := service.NewTaskService(taskRepository, taskStatusRepository)
	taskController := controller.NewTaskController(taskService)

	app.Post("/task", taskController.Create())
	app.Put("/task", taskController.Update())
	app.Patch("/task/status", taskController.UpdateStatus())
	app.Delete("/task/status", taskController.Delete())
	app.Get("/task", taskController.GetAll())
	app.Get("/task/:id", taskController.Get())

	app.Listen(":8080")
}