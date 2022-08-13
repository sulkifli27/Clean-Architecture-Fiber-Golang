package main

import (
	"github.com/gofiber/fiber/v2"

	"gitlab.com/d6825/golang_template/app"
	"gitlab.com/d6825/golang_template/pkg/routes"
	"gitlab.com/d6825/golang_template/pkg/utils"
)

func main() {
	//Initialization new app
	apps := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Midtrans",
		AppName:       "Midtrans v1",
	})

	app.Init()

	//Declare Middleware
	// middleware.FiberMiddleware(app)
	utils.LoadEnvironment()

	//Declare versioning or grouping
	routes.V1Routes(apps)
	//Start Server
	utils.StartServerWithGracefulShutdown(apps)
}
