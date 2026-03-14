package routes

import (
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/handler"
)

func SetupCategoryRoutes(mux *http.ServeMux, handler *handler.CategoryHandler) {
	mux.HandleFunc("POST /api/v1/categories", handler.Create)
	mux.HandleFunc("GET /api/v1/categories/{id}", handler.FindByID)
	mux.HandleFunc("GET /api/v1/categories", handler.FindAll)
	mux.HandleFunc("PUT /api/v1/categories/{id}", handler.Update)
	mux.HandleFunc("DELETE /api/v1/categories/{id}", handler.Delete)
}
