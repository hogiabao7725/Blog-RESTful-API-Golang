package routes

import (
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/handler"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/middleware"
)

func SetupPostRoutes(mux *http.ServeMux, handler *handler.PostHandler, auth *middleware.AuthMiddleware) {
	mux.Handle(
		"POST /api/v1/posts",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID, middleware.RoleUserID)(http.HandlerFunc(handler.Create))),
	)
	mux.HandleFunc("GET /api/v1/posts", handler.FindAll)
	mux.HandleFunc("GET /api/v1/posts/{id}", handler.FindByID)
	mux.Handle(
		"PUT /api/v1/posts/{id}",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID, middleware.RoleUserID)(http.HandlerFunc(handler.Update))),
	)
	mux.HandleFunc("GET /api/v1/categories/{category_id}/posts", handler.FindByCategoryID)
	mux.Handle(
		"DELETE /api/v1/posts/{id}",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID, middleware.RoleUserID)(http.HandlerFunc(handler.Delete))),
	)
}
