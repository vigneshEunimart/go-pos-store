package routes

import (
	pos_store "go-pos-store/app/controllers"
	"go-pos-store/app/response"
	validationschema "go-pos-store/app/routes/validationschema"

	"github.com/gofiber/fiber/v2"
)

func PosStore(Api fiber.Router) {

	// To Create a Pos Store
	Api.Post("/pos_store/create", validationschema.PosStoreValidate, pos_store.CreatePosStore, response.Response)

	// To Update Pos Store
	Api.Post("/pos_store/update", validationschema.UpdatePosStoreValidate, pos_store.UpdatePosStore, response.Response)

	// To List All The Stores
	Api.Get("/pos_store/list", pos_store.ListPosStores, response.Response)

	// To get the Pos Store details
	Api.Get("/pos_store/get", pos_store.ListPosStores, response.Response)

	// To Delete a Pos Store
	Api.Post("/pos_store/delete", validationschema.DeletePosStoreValidate, pos_store.DeletePosStore, response.Response)

	// To List Pos Store By Account id
	Api.Get("/pos_stores/by_account_id/list", pos_store.GetStoresByAccountId, response.Response)
}
