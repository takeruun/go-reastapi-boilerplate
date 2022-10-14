package main

import (
	"app/router"
	"net/http"
	"time"
)

func main() {
	time.Local = time.FixedZone("JST", 9*60*60)

	appRoute := router.NewAppRouter()
	userRoute := router.NewUserRouter()
	authRoute := router.NewAuthRouter()
	todoRoute := router.NewTodoRouter()
	r := router.NewRouter(appRoute, userRoute, authRoute, todoRoute)
	r.SetRouting()

	http.ListenAndServe(":3000", nil)
}
