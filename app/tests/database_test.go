package tests

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"gitlab.com/d6825/golang_template/platform/database"
	"gitlab.com/d6825/golang_template/platform/migrations"
)

func boot() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Config Cant load", err)
	}
}

func TestConnection(t *testing.T) {
	boot()
	db := database.GormMysqlConnection()
	if db.Error != nil {
		panic(db.Error)
	}
	print(db)
}

func TestRunMigration(t *testing.T) {
	boot()
	db := database.GormMysqlConnection()
	if db.Error != nil {
		panic(db.Error)
	}
	migrations.AutoMigration(*db)

	print("migration success")
}


