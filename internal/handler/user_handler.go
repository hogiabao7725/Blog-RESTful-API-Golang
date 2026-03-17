package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/request"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/response"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
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

	var req request.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Normalize()

	if err := req.Validate(); err != nil {
		errorx.WriteDomainError(w, err)
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
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.Response{
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

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req request.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Normalize()

	if err := req.Validate(); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	user, err := h.service.Login(r.Context(), req.UsernameOrEmail, req.Password)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "login successfully",
		Data: response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			RoleID:    user.RoleID,
			CreatedAt: user.CreatedAt,
		},
	})
}
