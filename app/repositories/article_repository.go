package repositories

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app"
	"gitlab.com/d6825/golang_template/app/interfaces"
	"gitlab.com/d6825/golang_template/app/models"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(app *app.App) interfaces.ArticleRepository {
	return &ArticleRepository{
		db: app.GetDbGorm(),
	}
}

func (repository *ArticleRepository) Get(context *fiber.Ctx) ([]models.Article, error) {
	var article []models.Article
	result := repository.db.Preload("Category").Find(&article)
	if err := result.Error; err != nil {
		return nil, err
	}
	return article, nil
}