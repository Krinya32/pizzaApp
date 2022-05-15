package main

import (
	"github.com/krinya32/pizzaApp"
	"github.com/krinya32/pizzaApp/pkg/handlers"
	"log"
)

func main() {
	handler := new(handlers.Handler)
	srv := new(pizzaApp.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Problems while running server: %s", err.Error())
	}
}
