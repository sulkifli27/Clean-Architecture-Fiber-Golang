package repositories

import (
	"log"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gitlab.com/d6825/golang_template/app"
	"gitlab.com/d6825/golang_template/platform/database"
)
func boot() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Config Cant load", err)
	}
}

func TestGetRepository(t *testing.T) {
   boot()
   main :=app.Init()
   db := database.GormMysqlConnection()  
   main.SetDbGorm(db)   
   ok :=NewArticleRepository(main)
    var ctx *fiber.Ctx
   _,err :=ok.Get(ctx)
   assert.NoError(t,err)

}