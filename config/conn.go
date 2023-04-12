package config

import (
	"fmt"
	"log"
	"project-1-chapter-2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "365442"
	port     = 5432
	dbname   = "db-gorm-api"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	conn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	db.AutoMigrate(models.Books{})
}

func GetDB() *gorm.DB {
	return db
}
