package models

import (
	"log"

	"github.com/zeimedee/go-postgres/models"
	"gorm-postgres/database"
)

func Migrate(db database.Dbinstance) {
	log.Println("running migrations")
	db.Db.AutoMigrate(&models.Book{})
}
