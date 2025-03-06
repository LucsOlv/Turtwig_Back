//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	"github.com/LucsOlv/Turtwing_Back/config"
	"github.com/LucsOlv/Turtwing_Back/internal/auth"
	"github.com/LucsOlv/Turtwing_Back/internal/handlers"
)

// provideApplication é um provider para criar uma nova instância da Application
func provideApplication(
	config *config.Config,
	authService *auth.Service,
	authHandler *handlers.AuthHandler,
	protectedHandler *handlers.ProtectedHandler,
) *Application {
	return &Application{
		config:           config,
		authService:      authService,
		authHandler:      authHandler,
		protectedHandler: protectedHandler,
	}
}

// InitializeApplication é a função que o Wire usará para gerar o código de injeção de dependência
func InitializeApplication() (*Application, error) {
	wire.Build(
		// Providers de configuração
		config.NewConfig,

		// Providers de serviços
		auth.NewService,

		// Providers de handlers
		handlers.NewAuthHandler,
		handlers.NewProtectedHandler,

		// Provider da aplicação
		provideApplication,
	)
	return nil, nil // Será substituído pelo Wire
}
