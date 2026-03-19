package routes

import (
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/handler"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/middleware"
)

func SetupCategoryRoutes(mux *http.ServeMux, handler *handler.CategoryHandler, auth *middleware.AuthMiddleware) {
	mux.Handle(
		"POST /api/v1/categories",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID)(http.HandlerFunc(handler.Create))),
	)
	mux.HandleFunc("GET /api/v1/categories/{id}", handler.FindByID)
	mux.HandleFunc("GET /api/v1/categories", handler.FindAll)
	mux.Handle(
		"PUT /api/v1/categories/{id}",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID)(http.HandlerFunc(handler.Update))),
	)
	mux.Handle(
		"DELETE /api/v1/categories/{id}",
		auth.RequireAuth(auth.RequireRoles(middleware.RoleAdminID)(http.HandlerFunc(handler.Delete))),
	)
}
