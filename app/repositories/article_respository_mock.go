package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"gitlab.com/d6825/golang_template/app/models"
)


type ArticleMockRepository struct {
	mock.Mock
}

func (a *ArticleMockRepository)  Get(context *fiber.Ctx) ([]models.Article, error) {
    var mock = a.Mock.Called(context)
    return mock.Get(0).([]models.Article), mock.Error(1)
}


