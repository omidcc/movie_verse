package handlers

import (
	"fmt"
	"movie-verse/db"
	"movie-verse/helpers"
	"movie-verse/models"
	"net/http"

	"github.com/go-chi/render"
	"github.com/spf13/viper"
)

func GetTenant(w http.ResponseWriter, r *http.Request) {

	conf := &models.Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	fmt.Printf("data, %v", conf.Db.Host)

	tenant, err := db.GetTenantById(1, conf)

	if err := render.Render(w, r, tenant); err != nil {
		render.Render(w, r, helpers.ErrRender(err))
		return
	}
}
