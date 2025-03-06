package handlers

import (
	"net/http"

	"github.com/LucsOlv/Turtwing_Back/internal/types"
)

type ProtectedHandler struct{}

func NewProtectedHandler() *ProtectedHandler {
	return &ProtectedHandler{}
}

// HandleProtected godoc
// @Summary      Rota protegida
// @Description  Exemplo de rota protegida que requer autenticação
// @Tags         protected
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200  {object}  types.Response
// @Failure      401  {object}  types.Response
// @Router       /protected [get]
func (h *ProtectedHandler) HandleProtected(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id")
	jsonResponse(w, http.StatusOK, types.Response{
		Status:  "success",
		Message: "Protected route accessed successfully",
		Data: map[string]interface{}{
			"user_id": userId,
		},
	})
}
