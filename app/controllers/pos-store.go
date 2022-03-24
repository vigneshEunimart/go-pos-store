package controllers

import (
	"fmt"
	"strconv"

	"go-pos-store/app/controllers/errorCodes"
	"go-pos-store/app/models"
	"go-pos-store/app/routes/validationschema"

	"github.com/gofiber/fiber/v2"
	uuidv4 "github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// To Create a Pos Store
func CreatePosStore(c *fiber.Ctx) error {

	var data models.Data

	err := c.BodyParser(&data)
	fmt.Printf("%T", data)
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while parsing the data")
		c.Locals("error_obj", errorCodes.Error_codes("INVALID_INPUT"))
		return c.Next()
	}

	data.Data.Store_id = uuidv4.New().String()
	data.Data.Created_by = data.User_info.UserId

	if data.Data.Inventory_type == "" {
		data.Data.Inventory_type = "decenter_inventory"
	}

	if data.Data.Currency == "" {
		data.Data.Currency = "INR"
	}

	if data.Data.Transition == "" {
		data.Data.Transition = "Exclusive tax"
	}

	filter := fiber.Map{
		"address.state": fiber.Map{"$ne": data.Data.Address.State},
		//"address.state": data.Data.Address.State,
		"gst_number": data.Data.Gst_number,
	}

	result := data.Data.FindStores(filter)
	if !result["status"].(bool) {
		c.Locals("status", result["status"])
		c.Locals("message", result["message"])
		c.Locals("error_obj", errorCodes.Error_codes(result["error_code"].(string)))
		return c.Next()
	}

	if len(result["data"].([]models.PosStoresSchema)) > 0 {
		c.Locals("status", false)
		c.Locals("message", "gst number already exists")
		c.Locals("error_obj", errorCodes.Error_codes("GST_EXISTS_ALREADY"))
		return c.Next()
	}

	res := data.Data.CreateStore()
	if !res["status"].(bool) {
		c.Locals("status", false)
		c.Locals("message", "error while adding the data")
		c.Locals("error_obj", errorCodes.Error_codes(res["error_code"].(string)))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "successfully added")
	c.Locals("data", data.Data)
	return c.Next()
}

// To Update a Pos Store
func UpdatePosStore(c *fiber.Ctx) error {

	var data models.Data

	if err := c.BodyParser(&data); err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while parsing the data")
		c.Locals("error_obj", errorCodes.Error_codes("INVALID_INPUT"))
		return c.Next()
	}

	data.Data.Updated_by = data.User_info.UserId

	gstFilter := fiber.Map{
		"address.state": fiber.Map{
			"$ne": data.Data.Address.State,
		},
		"gst_number": data.Data.Gst_number,
		"store_id": fiber.Map{
			"$ne": data.Data.Store_id,
		},
	}

	filter := fiber.Map{
		"account_id": data.Data.Account_id,
		"store_id":   data.Data.Store_id,
	}

	result := data.Data.FindStores(gstFilter)

	if len(result["data"].([]models.PosStoresSchema)) > 0 {
		c.Locals("status", false)
		c.Locals("message", "gst number found")
		c.Locals("error_obj", errorCodes.Error_codes("GST_EXISTS_ALREADY"))
		return c.Next()
	}

	fmt.Println("....", result["data"].([]models.PosStoresSchema))

	update_res := data.Data.UpdateStore(filter)
	if !update_res["status"].(bool) {
		c.Locals("status", update_res["status"])
		c.Locals("message", update_res["message"])
		c.Locals("error_obj", errorCodes.Error_codes(update_res["error_code"].(string)))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "update successfully")
	c.Locals("data", data.Data)

	return c.Next()
}

// To list the Pos Store Based on the demands or requirements
func ListPosStores(c *fiber.Ctx) error {

	filter := fiber.Map{
		"account_id": c.Query("account_id"),
	}

	if is_active := c.Query("is_active"); is_active == "" {
		filter["is_active"] = bson.M{"$exists": true}
	} else {

		filter["is_active"], _ = strconv.ParseBool(c.Query("is_active"))
	}

	if store_id := c.Query("store_id"); store_id != "" {
		filter["store_id"] = store_id
	}

	if is_deleted := c.Query("is_deleted"); is_deleted != "" {
		filter["is_deleted"] = false
	}

	var res models.PosStoresSchema

	result := res.FindStores(filter)
	if !result["status"].(bool) {
		c.Locals("status", result["status"])
		c.Locals("message", result["message"])
		c.Locals("error_obj", errorCodes.Error_codes(result["error_code"].(string)))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "data retrived successfully")
	c.Locals("data", result["data"])

	return c.Next()
}

// To Delete a pos store
func DeletePosStore(c *fiber.Ctx) error {

	var data validationschema.MainDeletedata

	err := c.BodyParser(&data)
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while parsing the data")
		c.Locals("error_obj", errorCodes.Error_codes("INVALID_INPUT"))
		return c.Next()
	}

	filter := fiber.Map{
		"account_id": data.Delete.Account_id,
		"store_id":   data.Delete.Store_id,
	}

	update := bson.M{
		"is_deleted": false}

	var s models.PosStoresSchema

	res := s.UpdateStore(filter, update)
	if !res["status"].(bool) {
		c.Locals("status", res["status"])
		c.Locals("message", res["message"])
		c.Locals("error_obj", errorCodes.Error_codes(res["error_code"].(string)))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "Deletion successfully done")
	c.Locals("data", data)

	return c.Next()
}

// To get the pos store details based on account id
func GetStoresByAccountId(c *fiber.Ctx) error {

	type StoreDetails struct {
		Store_id   string `json:"store_id"`
		Store_name string `json:"store_name"`
	}

	var storeData []StoreDetails

	filter := bson.M{
		"account_id": c.Query("account_id"),
		"is_deleted": false,
	}

	err := mgm.CollectionByName("stores").SimpleFind(&storeData, filter)
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while fetching the data")
		c.Locals("error_obj", errorCodes.Error_codes("NO_SOTRE_DATA_FOUND"))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "store details successfully fetched")
	c.Locals("data", storeData)

	return c.Next()
}
