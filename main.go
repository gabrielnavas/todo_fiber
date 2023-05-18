package main

import (
	"backend/task/controller"
	repository "backend/task/repositories"
	service "backend/task/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"
)

func main() {

	// env configs
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// env database
	var DATABASE_USERNAME_DEV string = os.Getenv("DATABASE_USERNAME_DEV")
	var DATABASE_PASSWORD_DEV string = os.Getenv("DATABASE_PASSWORD_DEV")
	var DATABASE_HOST_DEV string = os.Getenv("DATABASE_HOST_DEV")
	var DATABASE_PORT_DEV string = os.Getenv("DATABASE_PORT_DEV")
	var DATABASE_DATABASENAME_DEV string = os.Getenv("DATABASE_DATABASENAME_DEV")
	var DATABASE_SSLMODE_DEV string = os.Getenv("DATABASE_SSLMODE_DEV")

	// database config
	db := repository.MakeDatabase(
		DATABASE_USERNAME_DEV,
		DATABASE_PASSWORD_DEV,
		DATABASE_HOST_DEV,
		DATABASE_PORT_DEV,
		DATABASE_DATABASENAME_DEV,
		DATABASE_SSLMODE_DEV,
	)
	defer db.Close()

	// controllers, services, and repositories instances
	taskStatusRepository := repository.NewTaskStatusRepository(db)
	taskRepository := repository.NewTaskRepository(db, taskStatusRepository)
	taskService := service.NewTaskService(taskRepository, taskStatusRepository)
	taskController := controller.NewTaskController(taskService)

	// app configs
	app := fiber.New()

	// app middlewares
	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "*",
	}))
	// logger cli
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "25-12-2000",
		TimeZone:   "UTC",
	}))

	app.Post("/task", taskController.Create())
	app.Put("/task/:task_id", taskController.Update())
	app.Patch("/task/status/:task_id", taskController.UpdateStatus())
	app.Delete("/task/:task_id", taskController.Delete())
	app.Get("/task", taskController.GetAll())
	app.Get("/task/:id", taskController.Get())

	app.Listen(":8080")
}
