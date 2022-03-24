package controllers

import (
	"context"
	"go-pos-store/app/controllers/errorCodes"
	"go-pos-store/app/models"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	uuidv4 "github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// To  list the pos customer details based on account id
func ListPosCustomers(c *fiber.Ctx) error {

	if c.Query("account_id") == "" {
		c.Locals("status", false)
		c.Locals("message", "account_id must be present")
		c.Locals("error_obj", errorCodes.Error_codes("INVALID_INPUT"))
		return c.Next()
	}

	filter := bson.M{
		"account_id": c.Query("account_id"),
	}

	if mobile := c.Query("mobile"); mobile == "" {
		filter["mobile"] = bson.M{"$exists": true}
	} else {
		filter["mobile"] = c.Query("mobile")
	}

	if customer_type := c.Query("customer_type"); customer_type != "" {
		filter["customer_type"] = customer_type
	}

	if company_name := c.Query("company_name"); company_name != "" {
		filter["customer_name"] = company_name
	}

	filter["is_deleted"] = false

	findOptions := options.Find()

	var perPage, pageNo int

	if c.Query("perpage") != "" {
		perPage, _ = strconv.Atoi(c.Query("perpage"))
	} else {
		perPage = 10
	}

	if c.Query("pageno") != "" {
		pageNo, _ = strconv.Atoi(c.Query("pageno"))
	} else {
		pageNo = 1
	}

	findOptions.SetSkip((int64(pageNo) - 1) * int64(perPage))
	findOptions.SetLimit(int64(perPage))

	var result []models.PosCustomer

	err := mgm.CollectionByName("customer").SimpleFind(&result, filter, findOptions)
	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while fetching the data")
		c.Locals("error_obj", errorCodes.Error_codes("NO_DATA_FOUND"))
		return c.Next()
	}

	count, _ := mgm.CollectionByName("customer").CountDocuments(context.Background(), filter)

	//total_pages := int(count / int64(perPage))
	total_pages := int(math.Round(float64(count) / float64(perPage)))
	pagination := bson.M{
		"total_pages": total_pages,
	}
	pagination["perPage"] = perPage
	pagination["current_pageNo"] = pageNo
	if count == 0 || pageNo == total_pages {
		pagination["next_page"] = 0
	} else {
		pagination["next_page"] = pageNo + 1
	}
	if pageNo == 1 {
		pagination["previous_page"] = 0
	} else {
		pagination["previous_page"] = pageNo - 1
	}

	c.Locals("status", true)
	c.Locals("message", "data successfully retrived")
	c.Locals("data", result)
	c.Locals("pagination", pagination)
	return c.Next()
}

// To Create a Pos Customer
func CreatePosCustomer(c *fiber.Ctx) error {
	var data models.PosCostomerSchema
	err := c.BodyParser(&data)

	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while Parsing the data")
		c.Locals("error_obj", errorCodes.Error_codes("INVALID_INPUT"))
		return c.Next()
	}
	data.Data.Created_by = data.User_Info.User_id
	data.Data.Customer_id = uuidv4.New().String()

	if data.Data.Customer_type == "" {
		data.Data.Customer_type = "business"
	}

	var customer *models.PosCustomer = &data.Data

	err = mgm.CollectionByName("customer").Create(customer)

	if err != nil {
		c.Locals("status", false)
		c.Locals("message", "error while adding the data")
		c.Locals("error_obj", errorCodes.Error_codes("CUSTOMER_NOT_CREATED"))
		return c.Next()
	}

	c.Locals("status", true)
	c.Locals("message", "data successfully added")
	c.Locals("data", customer)

	return c.Next()

}
