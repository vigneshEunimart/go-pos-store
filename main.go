package main

import (
	"go-pos-store/app/routes"
	"go-pos-store/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())

	utils.Init()

	routes.PosRoute(app)

	app.Listen(":3000")

}
