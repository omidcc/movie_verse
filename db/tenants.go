package db

import (
	"fmt"
	"movie-verse/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetTenantById(tenantId int64, config *models.Config) (*models.Tenant, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	var tenant models.Tenant
	result := db.Table("Tenant").First(&tenant, tenantId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &tenant, nil
}
