package main

import (
	"log"
	"net/http"

	"github.com/tanvirislam06/go-api-test/internal/user"
	"github.com/tanvirislam06/go-api-test/pkg/routes"
)

func main() {
	// init the repository, service and handler
	userRepository := user.NewRepository()
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(*userService)

	// setup the router and start the server
	r := routes.SetupRouter(userHandler)
	log.Println("Server is running on: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
