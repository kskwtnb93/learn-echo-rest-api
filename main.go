package main

import (
	"learn-echo-rest-api/controller"
	"learn-echo-rest-api/db"
	"learn-echo-rest-api/repository"
	"learn-echo-rest-api/router"
	"learn-echo-rest-api/usecase"
	"learn-echo-rest-api/validator"
)

func main() {
	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)

	// サーバーを起動
	e.Logger.Fatal(e.Start(":8080"))
}
