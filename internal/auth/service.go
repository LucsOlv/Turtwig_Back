package auth

import (
	"errors"
	"time"

	"github.com/LucsOlv/Turtwing_Back/config"
	"github.com/LucsOlv/Turtwing_Back/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	config *config.Config
	// Aqui você adicionaria seu repositório de usuários
}

func NewService(config *config.Config) *Service {
	return &Service{
		config: config,
	}
}

func (s *Service) GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(s.config.JWTExpiresIn).Unix(),
	})

	return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *Service) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(s.config.JWTSecret), nil
	})
}

func (s *Service) SignIn(input models.SignInInput) (string, error) {
	// Aqui você buscaria o usuário no banco de dados
	// Por enquanto, vamos simular um usuário fixo para teste
	storedUser := &models.User{
		ID:       "1",
		Email:    "test@example.com",
		Password: "123",
	}

	if input.Email != storedUser.Email {
		return "", errors.New("usuário não encontrado")
	}

	return s.GenerateToken(storedUser)
}
