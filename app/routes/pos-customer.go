package routes

import (
	pos_customer "go-pos-store/app/controllers"
	"go-pos-store/app/response"

	"github.com/gofiber/fiber/v2"
)

func PosCustomer(Api fiber.Router) {

	// To List the customers list based on account id
	Api.Get("/pos_customer/list", pos_customer.ListPosCustomers, response.Response)

	// To Create a Pos Customer
	Api.Post("/pos_customer/create", pos_customer.CreatePosCustomer)
}
