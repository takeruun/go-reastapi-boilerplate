package main

import (
	"app/config"
	"app/controller"
	"app/database"
	"app/router"
	"app/usecase"
	"net/http"
	"time"
)

func main() {
	time.Local = time.FixedZone("JST", 9*60*60)

	db := config.NewDB()

	userRepository := database.NewUserRepository(db)
	todoRepository := database.NewTodoRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	authUsecase := usecase.NewAuthUsecase(userRepository)

	appController := controller.NewAppController()
	userController := controller.NewUserController(userUsecase)
	authController := controller.NewAuthController(authUsecase)
	todoController := controller.NewTodoController(todoUsecase)

	appRoute := router.NewAppRouter(appController)
	userRoute := router.NewUserRouter(userController)
	authRoute := router.NewAuthRouter(authController)
	todoRoute := router.NewTodoRouter(todoController)
	r := router.NewRouter(appRoute, userRoute, authRoute, todoRoute)
	r.SetRouting()

	http.ListenAndServe(":3000", nil)
}
