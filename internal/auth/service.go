package auth

import (
	"context"
	"errors"
	"time"

	"github.com/LucsOlv/Turtwing_Back/ent"
	"github.com/LucsOlv/Turtwing_Back/ent/user"
	"github.com/LucsOlv/Turtwing_Back/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	jwtSecret []byte
	client    *ent.Client
}

func NewService(jwtSecret []byte, client *ent.Client) *Service {
	return &Service{
		jwtSecret: jwtSecret,
		client:    client,
	}
}

func (s *Service) GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})

	return token.SignedString(s.jwtSecret)
}

func (s *Service) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return s.jwtSecret, nil
	})
}

func (s *Service) SignIn(input models.SignInInput) (string, error) {
	ctx := context.Background()

	// Buscar usuário pelo email
	user, err := s.client.User.Query().
		Where(user.EmailEQ(input.Email)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}

	// Verificar se o usuário está ativo
	if user.Status != "active" {
		return "", errors.New("user account is not active")
	}

	// Verificar a senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Atualizar último login
	_, err = s.client.User.UpdateOne(user).
		SetLastLogin(time.Now()).
		Save(ctx)
	if err != nil {
		return "", err
	}

	// Gerar o token JWT
	return s.GenerateToken(&models.User{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	})
}

func (s *Service) SignUp(input models.SignUpInput) (string, error) {
	ctx := context.Background()

	// Verificar se o email já existe
	exists, err := s.client.User.Query().
		Where(user.EmailEQ(input.Email)).
		Exist(ctx)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("email already exists")
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Criar o usuário
	user, err := s.client.User.Create().
		SetEmail(input.Email).
		SetPassword(string(hashedPassword)).
		SetName(input.Name).
		SetRole("user").
		SetStatus("active").
		Save(ctx)
	if err != nil {
		return "", err
	}

	// Gerar o token JWT
	return s.GenerateToken(&models.User{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	})
}
