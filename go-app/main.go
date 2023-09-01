package main

import (

	// Framework V4
	"fmt"
	"project/config"
	"project/handlers"
	"project/repos"
	"project/services"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	// load config
	config := config.Load()
	userRepo := repos.NewUserRepository(config.DbConn)
	userService := services.NewUserService(userRepo)

	healthHandler := handlers.NewHealthHandler()
	server.GET("/live", healthHandler.IsAlive)

	// REMOVE THAT ENDPOINT
	userHandler := handlers.NewUserHandler(userService)
	server.GET("/users/:id", userHandler.Get)

	// TODO: Register a new endpoint for POST user

	if err := server.Start(":1323"); err != nil {
		fmt.Println(err)
	}
}
