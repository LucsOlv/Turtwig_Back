// @title           Turtwig Back API
// @version         1.0
// @description     API para o projeto Turtwig Back
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
package main

import (
	"log"

	_ "github.com/LucsOlv/Turtwing_Back/docs"
	"github.com/LucsOlv/Turtwing_Back/internal/app"
)

func main() {
	container := app.BuildContainer()

	err := container.Invoke(func(application *app.Application) {
		if err := application.Start(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Error resolving dependencies: %v", err)
	}
}
