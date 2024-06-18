package main

import (
	"fmt"
	"injection_testing/config"
	"injection_testing/controller"
	"net/http"

	driver "injection_testing/repository"
	driverservice "injection_testing/service"

	"github.com/go-playground/validator/v10"
	zerolog "github.com/rs/zerolog/log"
)

// @title Injection demo
// @version 1.0
// @description This is a sample server for the IoT Project API.
// @host localhost:8080
// @BasePath /
// @schemes http https

func main() {
	fmt.Println("Hello, World!")
	db := config.ConnectDatabase()
	
	validate := validator.New()

	driverRepository := driver.NewDriverRepository(db)
	driverService := driverservice.NewDriverService(driverRepository, validate)
	driverController := controller.NewDriverController(driverService, )

	controllers := &Controllers{
		DriverController:  driverController,
	}
	routes := NewRouter(controllers)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		zerolog.Fatal().Err(err).Msg("Server Stopped")
	}
}