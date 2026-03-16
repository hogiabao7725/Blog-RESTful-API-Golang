package routes

import (
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/handler"
)

func SetupPostRoutes(mux *http.ServeMux, handler *handler.PostHandler) {
	mux.HandleFunc("POST /api/v1/posts", handler.Create)
	mux.HandleFunc("GET /api/v1/posts", handler.FindAll)
	mux.HandleFunc("GET /api/v1/posts/{id}", handler.FindByID)
	mux.HandleFunc("PUT /api/v1/posts/{id}", handler.Update)
	mux.HandleFunc("GET /api/v1/categories/{category_id}/posts", handler.FindByCategoryID)
	mux.HandleFunc("DELETE /api/v1/posts/{id}", handler.Delete)
}
