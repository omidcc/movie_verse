package models

import (
	"net/http"
	"time"
)

type Tenant struct {
	TenantId  int64      `gorm:"primarykey" json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Address   string     `json:"address"`
	LogoUrl   string     `json:"logo_url"`
	CreatedAt time.Time  `gorm:"type:time" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:time" json:"updated_at"`
}

// Render implements render.Renderer.
func (tnt *Tenant) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
