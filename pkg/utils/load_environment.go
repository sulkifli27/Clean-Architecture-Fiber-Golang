package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func boot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Config Cant load", err)
	}
}

func LoadEnvironment() {
	boot()
}
