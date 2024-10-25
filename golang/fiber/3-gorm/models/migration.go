package models

import (
	"log"

	"github.com/zeimedee/go-postgres/models"
	"gorm-postgres/database"
)

func Migrate() {
	db := database.DB().Db

	log.Println("running migrations")
	db.AutoMigrate(&models.Book{})
}
