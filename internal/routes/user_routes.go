package routes

import (
	"net/http"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/handler"
)

func SetupUserRoutes(mux *http.ServeMux, handler *handler.UserHandler) {
	mux.HandleFunc("POST /api/v1/users/register", handler.Register)
}
