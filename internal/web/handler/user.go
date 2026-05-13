package handler

import (
	"net/http"

	"github.com/m-bromo/go-auth-template/internal/service"
	"github.com/m-bromo/go-auth-template/internal/web/models"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, err := h.userService.GetProfile(r.Context(), id)
	if err != nil {
		HandleError(w, err)
		return
	}

	HandleJSON(w, http.StatusOK, models.GetProfilePayload{
		Email:    user.Email,
		Username: user.Username,
	})

}
