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

type PostHandler struct {
	service domain.PostService
}

func NewPostHandler(postService domain.PostService) *PostHandler {
	return &PostHandler{
		service: postService,
	}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	post := &domain.Post{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		UserID:      req.UserID,
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
	posts, err := h.service.FindAll(r.Context())
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	var postResponses []response.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, toPostResponse(post))
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "get posts successfully",
		Data:    postResponses,
	})
}

func (h *PostHandler) FindByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.ParsePathID(r, "category_id")
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	posts, err := h.service.FindByCategoryID(r.Context(), categoryID)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	var postResponses []response.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, toPostResponse(post))
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "get posts by category successfully",
		Data:    postResponses,
	})
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	var req request.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	post := &domain.Post{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		UserID:      req.UserID,
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
	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid post id")
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
