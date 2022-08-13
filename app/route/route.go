// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package route

import (
	// "github.com/google/wire"
	"gitlab.com/d6825/golang_template/app/controllers"
	// "gitlab.com/d6825/golang_template/app/interfaces"
	"gitlab.com/d6825/golang_template/app/repositories"
	"gitlab.com/d6825/golang_template/app/services"
	"gitlab.com/d6825/golang_template/app"
	"github.com/gofiber/fiber/v2"

)


func Inizialize(a *fiber.App)  {
	db := app.Init()
	api := a.Group("/api/v1")
	// article
	articleRepository := repositories.NewArticleRepository(db)
	articleService := services.NewArticleService(articleRepository)
	articleController := controllers.NewArticleController(articleService)

	// route article
	api.Get("/articles", 	articleController.GetArticle)

}