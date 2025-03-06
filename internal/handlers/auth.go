package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LucsOlv/Turtwing_Back/internal/auth"
	"github.com/LucsOlv/Turtwing_Back/internal/models"
	"github.com/LucsOlv/Turtwing_Back/internal/types"
)

type AuthHandler struct {
	authService *auth.Service
}

func NewAuthHandler(authService *auth.Service) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// SignIn godoc
// @Summary      Autenticar usuário
// @Description  Autentica um usuário e retorna um token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.SignInInput true "Credenciais do usuário"
// @Success      200  {object}  types.Response
// @Failure      400  {object}  types.Response
// @Failure      401  {object}  types.Response
// @Router       /auth/signin [post]
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input models.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		jsonResponse(w, http.StatusBadRequest, types.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	token, err := h.authService.SignIn(input)
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, types.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	jsonResponse(w, http.StatusOK, types.Response{
		Status: "success",
		Data: map[string]string{
			"token": token,
		},
	})
}

// SignUp godoc
// @Summary      Cadastrar usuário
// @Description  Cadastra um novo usuário
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.SignUpInput true "Dados do usuário"
// @Success      200  {object}  types.Response
// @Failure      400  {object}  types.Response
// @Failure      401  {object}  types.Response
// @Router       /auth/signup [post]
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input models.SignUpInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		jsonResponse(w, http.StatusBadRequest, types.Response{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	token, err := h.authService.SignUp(input)
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, types.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	jsonResponse(w, http.StatusOK, types.Response{
		Status: "success",
		Data: map[string]string{
			"token": token,
		},
	})
}

func jsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
