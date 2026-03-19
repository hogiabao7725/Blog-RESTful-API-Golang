package routes

import (
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/handler"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/middleware"
)

func SetupCommentRoutes(mux *http.ServeMux, handler *handler.CommentHandler, auth *middleware.AuthMiddleware) {
	mux.Handle(
		"POST /api/v1/posts/{post_id}/comments",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID, middleware.RoleUserID)(http.HandlerFunc(handler.Create))),
	)
	mux.HandleFunc("GET /api/v1/posts/{post_id}/comments", handler.FindByPostID)
	mux.HandleFunc("GET /api/v1/comments/{id}", handler.FindByID)
	mux.Handle(
		"PUT /api/v1/comments/{id}",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID, middleware.RoleUserID)(http.HandlerFunc(handler.Update))),
	)
	mux.Handle(
		"DELETE /api/v1/comments/{id}",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID, middleware.RoleUserID)(http.HandlerFunc(handler.Delete))),
	)
}
