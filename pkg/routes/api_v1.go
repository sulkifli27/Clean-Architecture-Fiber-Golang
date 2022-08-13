package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/route"
)

func V1Routes(a *fiber.App) {
	route.Inizialize(a)
}
