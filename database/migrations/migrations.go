package migrations

import (
	"github.com/PedroPaiao/rest-api-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
