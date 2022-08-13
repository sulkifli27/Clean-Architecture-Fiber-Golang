package migrations

import (
	"gitlab.com/d6825/golang_template/app/models"
	"gorm.io/gorm"
)

func AutoMigration(db gorm.DB) {
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Category{})
}
