package database

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var internalDB Dbinstance

var setup sync.Once

// connectDb
func ConnectDb() {
	setup.Do(func() {
		dsn := "postgres://postgres:postgres@127.0.0.1:5432/pgsql?search_path=public&sslmode=disable"

		println(dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatal("Failed to connect to database. \n", err)
		}

		log.Println("connected")
		db.Logger = logger.Default.LogMode(logger.Info)

		internalDB = Dbinstance{
			Db: db,
		}
	})
}

func DB() Dbinstance {
	ConnectDb()

	return internalDB
}
