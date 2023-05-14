package main

import (
	"backend/task/controller"
	repository "backend/task/repositories"
	service "backend/task/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	// app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "*",
	}))

	db := makeDatabase()
	defer db.Close()

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	taskStatusRepository := repository.NewTaskStatusRepository(db)
	taskRepository := repository.NewTaskRepository(db, taskStatusRepository)
	taskService := service.NewTaskService(taskRepository, taskStatusRepository)
	taskController := controller.NewTaskController(taskService)

	app.Post("/task", taskController.Create())
	app.Put("/task/:task_id", taskController.Update())
	app.Patch("/task/status/:task_id", taskController.UpdateStatus())
	app.Delete("/task/:task_id", taskController.Delete())
	app.Get("/task", taskController.GetAll())
	app.Get("/task/:id", taskController.Get())

	app.Listen(":8080")
}
