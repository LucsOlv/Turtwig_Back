package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/LucsOlv/Turtwing_Back/config"
	"github.com/LucsOlv/Turtwing_Back/internal/auth"
	"github.com/LucsOlv/Turtwing_Back/internal/handlers"
	authMiddleware "github.com/LucsOlv/Turtwing_Back/internal/middleware"
)

// Application contém todas as dependências necessárias para a aplicação
type Application struct {
	config           *config.Config        // Configurações da aplicação
	authService      *auth.Service         // Serviço de autenticação
	authHandler      *handlers.AuthHandler // Handler de autenticação
	protectedHandler *handlers.ProtectedHandler // Handler para rotas protegidas
}

func (app *Application) SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Middleware globais
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Rotas públicas
	r.Group(func(r chi.Router) {
		r.Post("/auth/signin", app.authHandler.SignIn)
	})

	// Rotas protegidas
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.AuthMiddleware(app.authService))
		r.Get("/protected", app.protectedHandler.HandleProtected)
	})

	return r
}

func (app *Application) Start() error {
	router := app.SetupRoutes()
	
	log.Printf("API servidor iniciado na porta %s...\n", app.config.ServerPort)
	log.Printf("Documentação Swagger disponível em: http://localhost%s/swagger/index.html\n", app.config.ServerPort)
	return http.ListenAndServe(app.config.ServerPort, router)
}
