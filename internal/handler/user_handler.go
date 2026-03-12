package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/request"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/response"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/utils"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(s domain.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	// Check METHOD
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req request.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user := &domain.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := h.service.Register(r.Context(), user)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrAlreadyExists):
			utils.WriteError(w, http.StatusConflict, err.Error())
		default:
			utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	utils.WriteJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Message: "register user successfully",
		Data: response.UserResponse{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Username:  createdUser.Username,
			Email:     createdUser.Email,
			RoleID:    createdUser.RoleID,
			CreatedAt: createdUser.CreatedAt,
		},
	})
}
