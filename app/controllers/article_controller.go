package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/helpers"
	"gitlab.com/d6825/golang_template/app/interfaces"
	"gitlab.com/d6825/golang_template/app/models"
)

type ArticleController struct {
	ArticleService interfaces.ArticleService
}

func NewArticleController(articleService interfaces.ArticleService) interfaces.ArticleControlller {
	return &ArticleController{ArticleService: articleService}
}

func (service ArticleController) GetArticle(context *fiber.Ctx) error {
	article := &models.Article{}

	result, err := service.ArticleService.Get(context)

	fmt.Println(err);

	if err != nil { helpers.JsonResponse( context, "articles were not found", fiber.StatusNotFound, article, err)
	}

	return helpers.JsonResponse(context,"Success", fiber.StatusOK, result,nil)
}

