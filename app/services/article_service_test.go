package services

import (
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/models"
	"gitlab.com/d6825/golang_template/app/repositories"
)


func TestArticleGet(t *testing.T){ 
    t.Run("Failed Get Data", func(t *testing.T) {
        var ctx  *fiber.Ctx
        var mocks = new(repositories.ArticleMockRepository)
        var test = NewArticleService(mocks)

        mocks.On("Get", ctx).Return([]models.Article{}, errors.New("data filed get article"))

        var _ , err = test.Get(ctx) 
        if err == nil {
            panic("failed test get data")
        }
    })
    
    t.Run("Succes Get Data", func(t *testing.T) {
        var ctx  *fiber.Ctx
        var mocks = new(repositories.ArticleMockRepository)
        var test = NewArticleService(mocks)

        mocks.On("Get", ctx).Return([]models.Article{models.Article{Title: "halo"}}, nil)

        var _ , err = test.Get(ctx) 
        if err != nil {
            panic("succes get data")
        }
    })
    
}