package main

import (
	"log"

	"github.com/LucsOlv/Turtwing_Back/internal/app"
)

func main() {
	container := app.BuildContainer()

	// Invoca a aplicação
	err := container.Invoke(func(application *app.Application) {
		if err := application.Start(); err != nil {
			log.Fatalf("Error starting application: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Error resolving dependencies: %v", err)
	}
}
