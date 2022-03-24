package routes

import (
	//"go-pos-store/app/routes"

	"github.com/gofiber/fiber/v2"
)

func PosRoute(app *fiber.App) {

	Store_management_api := app.Group("/pos_store_management")

	// file-path : Pos-store route
	PosStore(Store_management_api)

	// file-path : Pos-customer route
	PosCustomer(Store_management_api)

}
