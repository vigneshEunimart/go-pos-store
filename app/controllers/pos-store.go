package controllers

import (
	"context"
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

	filter := bson.M{
		"address.state": data.Data.Address.State,
		"gst_number":    data.Data.Gst_number,
	}

	var check []models.PosStoresSchema

	err = mgm.CollectionByName("stores").SimpleFind(&check, filter)
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while fetching the data")
		c.Locals("error_obj", errorCodes.Error_codes("NO_SOTRE_DATA_FOUND"))
		return c.Next()
	}

	if check != nil {
		c.Locals("status", false)
		c.Locals("message", "gst number already exists")
		c.Locals("error_obj", errorCodes.Error_codes("GST_EXISTS_ALREADY"))
		return c.Next()
	}
	var add *models.PosStoresSchema = &data.Data

	err = mgm.CollectionByName("stores").Create(add)
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while adding the data")
		c.Locals("error_obj", errorCodes.Error_codes("STORE_NOT_CREATED"))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "successfully added")
	c.Locals("data", add)
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

	gstFliter := bson.M{
		"address.state": data.Data.Address.State,
		"gst_number":    data.Data.Gst_number,
		"store_id":      data.Data.Store_id,
	}

	filter := bson.M{
		"account_id": data.Data.Account_id,
		"store_id":   data.Data.Store_id,
	}

	var check []models.PosStoresSchema

	_ = mgm.CollectionByName("stores").SimpleFind(&check, gstFliter)

	//fmt.Println("....", check)
	if len(check) <= 0 {
		c.Locals("status", false)
		c.Locals("message", "gst number not found")
		c.Locals("error_obj", errorCodes.Error_codes("GST_EXISTS_ALREADY"))
		return c.Next()
	}

	data.Data.Created_by = check[0].Created_by

	res, err := mgm.CollectionByName("stores").UpdateOne(context.Background(), filter, bson.M{"$set": data.Data})
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while updating the data")
		c.Locals("error_obj", errorCodes.Error_codes("ERROR_WHILE_DATA_UPDATE"))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "update successfully")
	c.Locals("data", res)

	return c.Next()
}

// To list the Pos Store Based on the demands or requirements
func ListPosStores(c *fiber.Ctx) error {

	filter := bson.M{
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

	var result []models.PosStoresSchema

	err := mgm.CollectionByName("stores").SimpleFind(&result, filter)
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while listing the data")
		c.Locals("error_obj", errorCodes.Error_codes("NO_SOTRE_DATA_FOUND"))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "data retrived successfully")
	c.Locals("data", result)

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

	filter := bson.M{
		"account_id": data.Delete.Account_id,
		"store_id":   data.Delete.Store_id,
	}

	update := bson.M{
		"is_deleted": true}

	res, err := mgm.CollectionByName("stores").UpdateOne(context.Background(), filter, bson.M{"$set": update})
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while Deleting the data")
		c.Locals("error_obj", errorCodes.Error_codes("ERROR_WHILE_DATA_UPDATE"))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "Deletion successfully done")
	c.Locals("data", res)

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
