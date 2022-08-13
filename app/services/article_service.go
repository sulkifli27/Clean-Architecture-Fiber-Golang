package services

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/interfaces"
	"gitlab.com/d6825/golang_template/app/models"
)


type ArticleService struct {
	ArticleRepository interfaces.ArticleRepository
}

func NewArticleService(articleRepository interfaces.ArticleRepository) interfaces.ArticleRepository{
	return &ArticleService{
		ArticleRepository: articleRepository,
	}
}

func (repository *ArticleService) Get(context *fiber.Ctx) ([]models.Article, error) {
	var article []models.Article
	article, err := repository.ArticleRepository.Get(context)
	if err != nil {
		return article, err
	}
	return article, nil
}

