package main

import (
	"app/controller"
	"app/router"
	"net/http"
	"time"
)

func main() {
	time.Local = time.FixedZone("JST", 9*60*60)

	appController := controller.NewAppController()
	userController := controller.NewUserController()
	authController := controller.NewAuthController()
	todoController := controller.NewTodoController()

	appRoute := router.NewAppRouter(appController)
	userRoute := router.NewUserRouter(userController)
	authRoute := router.NewAuthRouter(authController)
	todoRoute := router.NewTodoRouter(todoController)
	r := router.NewRouter(appRoute, userRoute, authRoute, todoRoute)
	r.SetRouting()

	http.ListenAndServe(":3000", nil)
}
