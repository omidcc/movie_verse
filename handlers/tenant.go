package handlers

import (
	"fmt"
	"movie-verse/helpers"
	"movie-verse/models"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/spf13/viper"
)

func GetTenant(w http.ResponseWriter, r *http.Request) {

	var timeNow = time.Now()
	var tenant = &models.Tenant{
		Id:        1,
		Name:      "Test tenant",
		Email:     "testemail@gmail.com",
		Address:   "H#1, R#1, DHaka-1215",
		LogoUrl:   "http://image.com/name.jpeg",
		CreatedAt: timeNow,
		UpdatedAt: &timeNow,
	}
	conf := &models.Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	fmt.Println("new config", conf.Db.Username)

	if err := render.Render(w, r, tenant); err != nil {
		render.Render(w, r, helpers.ErrRender(err))
		return
	}
}
