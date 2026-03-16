package routes

import (
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/handler"
)

func SetupCommentRoutes(mux *http.ServeMux, handler *handler.CommentHandler) {
	mux.HandleFunc("POST /api/v1/posts/{post_id}/comments", handler.Create)
	mux.HandleFunc("GET /api/v1/posts/{post_id}/comments", handler.FindByPostID)
	mux.HandleFunc("GET /api/v1/comments/{id}", handler.FindByID)
	mux.HandleFunc("PUT /api/v1/comments/{id}", handler.Update)
	mux.HandleFunc("DELETE /api/v1/comments/{id}", handler.Delete)
}
