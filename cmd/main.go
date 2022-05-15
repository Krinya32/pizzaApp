package main

import (
	"github.com/krinya32/pizzaApp"
	"github.com/krinya32/pizzaApp/pkg/handlers"
	"github.com/krinya32/pizzaApp/pkg/repository"
	"github.com/krinya32/pizzaApp/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(pizzaApp.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Problems while running server: %s", err.Error())
	}
}
