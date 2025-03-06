package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/dig"

	"github.com/LucsOlv/Turtwing_Back/config"
	"github.com/LucsOlv/Turtwing_Back/internal/auth"
	"github.com/LucsOlv/Turtwing_Back/internal/database"
	"github.com/LucsOlv/Turtwing_Back/internal/handlers"
	authMiddleware "github.com/LucsOlv/Turtwing_Back/internal/middleware"
)

// Application contém todas as dependências necessárias para a aplicação
type Application struct {
	config           *config.Config             // Configurações da aplicação
	authService      *auth.Service              // Serviço de autenticação
	authHandler      *handlers.AuthHandler      // Handler de autenticação
	protectedHandler *handlers.ProtectedHandler // Handler para rotas protegidas
}

// NewApplication cria uma nova instância da aplicação usando dig
func NewApplication(
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
		r.Post("/auth/signup", app.authHandler.SignUp)
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

	log.Printf("API servidor iniciado na porta %s...\n", app.config.Port)
	log.Printf("Documentação Swagger disponível em: http://localhost:%s/swagger/index.html\n", app.config.Port)
	return http.ListenAndServe(":"+app.config.Port, router)
}

// provideJWTSecret retorna o JWT secret como []byte
func provideJWTSecret(cfg *config.Config) []byte {
	return []byte(cfg.JWTSecret)
}

// BuildContainer configura o container de injeção de dependência
func BuildContainer() *dig.Container {
	container := dig.New()

	// Registra os providers
	providers := []interface{}{
		config.NewConfig,
		database.NewClient,
		provideJWTSecret,
		auth.NewService,
		handlers.NewAuthHandler,
		handlers.NewProtectedHandler,
		NewApplication,
	}

	for _, provider := range providers {
		if err := container.Provide(provider); err != nil {
			log.Fatalf("Error providing %T: %v", provider, err)
		}
	}

	return container
}
