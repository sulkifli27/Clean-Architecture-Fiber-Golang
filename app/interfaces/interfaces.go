package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/models"
)

type (

	// article interface
	ArticleRepository interface {
		Get(context *fiber.Ctx) ([]models.Article, error)
	}

	ArticleService interface {
		Get(context *fiber.Ctx) ([]models.Article, error)
	}

	ArticleControlller interface {
		GetArticle(context *fiber.Ctx) error
	}

)
