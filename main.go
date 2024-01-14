package main

import (
	"movie-verse/routers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	viper.SetConfigFile("conf.yaml")
	viper.ReadInConfig()

	r.Mount("/tenants", routers.TenantRouter())

	http.ListenAndServe(":3333", r)
}
