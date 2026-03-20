package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/request"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/response"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/middleware"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/utils"
)

type PostHandler struct {
	service domain.PostService
}

func NewPostHandler(postService domain.PostService) *PostHandler {
	return &PostHandler{
		service: postService,
	}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	authUserID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		errorx.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req request.PostRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Normalize()

	if err := req.Validate(); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	post := &domain.Post{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		UserID:      authUserID,
		CategoryID:  req.CategoryID,
	}

	createdPost, err := h.service.Create(r.Context(), post)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Message: "create post successfully",
		Data:    toPostResponse(createdPost),
	})
}

func (h *PostHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	post, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "get post successfully",
		Data:    toPostResponse(post),
	})
}

func (h *PostHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	pagination, err := utils.ParsePagination(r)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	// Check for search query parameter
	query := r.URL.Query().Get("query")

	var paginated *domain.PaginatedPosts

	if query != "" {
		// If query parameter is provided, perform search with pagination
		paginated, err = h.service.SearchPaginated(r.Context(), query, pagination.GetOffset(), pagination.Limit)
	} else {
		// Otherwise, return all posts with pagination
		paginated, err = h.service.FindAllPaginated(r.Context(), pagination.GetOffset(), pagination.Limit)
	}

	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	var postResponses []response.PostResponse
	for _, post := range paginated.Posts {
		postResponses = append(postResponses, toPostResponse(post))
	}

	paginationMeta := dto.NewPaginationMeta(pagination.Page, pagination.Limit, paginated.Total)

	response.WriteJSON(w, http.StatusOK, response.PaginatedResponse{
		Success: true,
		Message: "get posts successfully",
		Data:    postResponses,
		Meta:    paginationMeta,
	})
}

func (h *PostHandler) FindByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.ParsePathID(r, "category_id")
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	// Parse pagination parameters
	pagination, err := utils.ParsePagination(r)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	paginated, err := h.service.FindByCategoryIDPaginated(r.Context(), categoryID, pagination.GetOffset(), pagination.Limit)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	var postResponses []response.PostResponse
	for _, post := range paginated.Posts {
		postResponses = append(postResponses, toPostResponse(post))
	}

	paginationMeta := dto.NewPaginationMeta(pagination.Page, pagination.Limit, paginated.Total)

	response.WriteJSON(w, http.StatusOK, response.PaginatedResponse{
		Success: true,
		Message: "get posts by category successfully",
		Data:    postResponses,
		Meta:    paginationMeta,
	})
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	authUserID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		errorx.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	authRoleID, ok := middleware.RoleIDFromContext(r.Context())
	if !ok {
		errorx.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	existingPost, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	if authRoleID != middleware.RoleAdminID && existingPost.UserID != authUserID {
		errorx.WriteError(w, http.StatusForbidden, "forbidden")
		return
	}

	var req request.PostRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Normalize()

	if err := req.Validate(); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	post := &domain.Post{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		UserID:      existingPost.UserID,
		CategoryID:  req.CategoryID,
	}

	updatedPost, err := h.service.Update(r.Context(), post)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "update post successfully",
		Data:    toPostResponse(updatedPost),
	})
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	authUserID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		errorx.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	authRoleID, ok := middleware.RoleIDFromContext(r.Context())
	if !ok {
		errorx.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	existingPost, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	if authRoleID != middleware.RoleAdminID && existingPost.UserID != authUserID {
		errorx.WriteError(w, http.StatusForbidden, "forbidden")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "delete post successfully",
	})
}

func toPostResponse(post *domain.Post) response.PostResponse {
	return response.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		UserID:      post.UserID,
		CategoryID:  post.CategoryID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}
