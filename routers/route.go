package routers

import (
	"movie-verse/handlers"

	"github.com/go-chi/chi"
)

func TenantRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.GetTenant)
	return r
}
