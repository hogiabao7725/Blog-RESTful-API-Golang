package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/request"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/dto/response"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/middleware"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/utils"
)

type CommentHandler struct {
	service domain.CommentService
}

func NewCommentHandler(service domain.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	authUserID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		errorx.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	postID, err := utils.ParsePathID(r, "post_id")
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	var req request.CreateCommentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Normalize()

	if err := req.Validate(); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	comment := &domain.Comment{
		Body:   req.Body,
		UserID: authUserID,
		PostID: postID,
	}

	createdComment, err := h.service.Create(r.Context(), comment)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Message: "create comment successfully",
		Data:    toCommentResponse(createdComment),
	})
}

func (h *CommentHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromURL(r)
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid comment id")
		return
	}

	comment, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "get comment successfully",
		Data:    toCommentResponse(comment),
	})
}

func (h *CommentHandler) FindByPostID(w http.ResponseWriter, r *http.Request) {
	postID, err := utils.ParsePathID(r, "post_id")
	if err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	comments, err := h.service.FindByPostID(r.Context(), postID)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	var commentResponses []response.CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, toCommentResponse(comment))
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "get comments by post successfully",
		Data:    commentResponses,
	})
}

func (h *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
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
		errorx.WriteError(w, http.StatusBadRequest, "invalid comment id")
		return
	}

	existingComment, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	if authRoleID != middleware.RoleAdminID && existingComment.UserID != authUserID {
		errorx.WriteError(w, http.StatusForbidden, "forbidden")
		return
	}

	var req request.UpdateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.Normalize()

	if err := req.Validate(); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	comment := &domain.Comment{
		ID:     id,
		Body:   req.Body,
		UserID: existingComment.UserID,
		PostID: existingComment.PostID,
	}

	updatedComment, err := h.service.Update(r.Context(), comment)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "update comment successfully",
		Data:    toCommentResponse(updatedComment),
	})
}

func (h *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
		errorx.WriteError(w, http.StatusBadRequest, "invalid comment id")
		return
	}

	existingComment, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	if authRoleID != middleware.RoleAdminID && existingComment.UserID != authUserID {
		errorx.WriteError(w, http.StatusForbidden, "forbidden")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		errorx.WriteDomainError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Success: true,
		Message: "delete comment successfully",
	})
}

func toCommentResponse(comment *domain.Comment) response.CommentResponse {
	return response.CommentResponse{
		ID:        comment.ID,
		Body:      comment.Body,
		UserID:    comment.UserID,
		PostID:    comment.PostID,
		CreatedAt: comment.CreatedAt,
	}
}
