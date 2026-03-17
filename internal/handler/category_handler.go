package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/request"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/response"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/utils"
)

type CategoryHandler struct {
	service domain.CategoryService
}

func NewCategoryHandler(s domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: s,
	}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Normalize()

	if err := req.Validate(); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	category := &domain.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	createdCategory, err := h.service.Create(r.Context(), category)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Message: "create category successfully",
		Data: response.CategoryResponse{
			ID:          createdCategory.ID,
			Name:        createdCategory.Name,
			Description: createdCategory.Description,
		},
	})
}

func (h *CategoryHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	category, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "get category successfully",
		Data: response.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	})
}

func (h *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.FindAll(r.Context())
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, response.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "get categories successfully",
		Data:    categoryResponses,
	})
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	var req request.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Normalize()

	if err := req.Validate(); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	category := &domain.Category{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	updatedCategory, err := h.service.Update(r.Context(), category)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "update category successfully",
		Data: response.CategoryResponse{
			ID:          updatedCategory.ID,
			Name:        updatedCategory.Name,
			Description: updatedCategory.Description,
		},
	})
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "delete category successfully",
	})
}
