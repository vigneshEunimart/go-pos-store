package models

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// DB schema - struct to store a data as a format in db
type (
	PosStoresSchema struct {
		mgm.DefaultModel
		Store_id            string `json:"store_id"`
		Store_name          string `json:"store_name" validate:"required"`
		Account_id          string `json:"account_id" validate:"required"`
		Inventory_type      string `json:"inventory_type" validate:"required"` //enum  default : "decenter_inventory"
		Store_phone_number  string `json:"store_phone_number" validate:"required"`
		Business_reg_number string `json:"business_reg_number" validate:"required"`
		Gst_number          string `json:"gst_number" validate:"required"`
		Iec_number          string `json:"iec_number" validate:"required"`
		Gst_file            string `json:"gst_file" validate:"required"`
		Iec_file            string `json:"iec_file" validate:"required"`
		Currency            string `json:"currency" validate:"required"`   //default : "inr"
		Transition          string `json:"transition" validate:"required"` //enum  default : "inclusive tax"
		Address             Addr   `json:"address" validate:"required"`
		Manager_name        string `json:"manager_name" validate:"required"`
		Manager_phone       string `json:"manager_phone" validate:"required"`
		Is_active           bool   `json:"is_active"`
		Is_deleted          bool   `json:"is_deleted"`
		Updated_by          string `json:"updated_by"`
		Created_by          string `json:"created_by"`
	}

	Addr struct {
		Zip     string `json:"zip" validate:"required"`
		City    string `json:"city" validate:"required"`
		State   string `json:"state" validate:"required"`
		Country string `json:"country" validate:"required"`
	}

	User_Info struct {
		UserId string `json:"user_id" validate:"required"`
	}

	// To achieve the input payload format
	Data struct {
		Data      PosStoresSchema `json:"data"`
		User_info User_Info       `json:"user_info"`
	}
)

func (store PosStoresSchema) CreateStore() map[string]interface{} {

	err := mgm.CollectionByName("stores").Create(&store)
	if err != nil {
		return map[string]interface{}{
			"status":     false,
			"message":    "error while adding the data",
			"error_code": "STORE_NOT_CREATED",
		}
	}

	return map[string]interface{}{
		"status":  true,
		"message": "data added successfully",
	}
}

func (store PosStoresSchema) FindStores(filter map[string]interface{}) map[string]interface{} {

	var stores []PosStoresSchema

	err := mgm.CollectionByName("stores").SimpleFind(&stores, filter)
	if err != nil {
		return map[string]interface{}{
			"status":     false,
			"message":    "error while retriving the data",
			"error_code": "NO_SOTRE_DATA_FOUND",
		}
	}

	return map[string]interface{}{
		"status":  true,
		"message": "data retrived successfully",
		"data":    stores,
	}
}

func (store PosStoresSchema) UpdateStore(filter map[string]interface{}, updateData ...map[string]interface{}) map[string]interface{} {

	if updateData != nil {

		res, err := mgm.CollectionByName("stores").UpdateOne(context.Background(), filter, bson.M{"$set": updateData[0]})
		if err != nil {
			return map[string]interface{}{
				"status":     false,
				"message":    "error while updating the data",
				"error_code": "ERROR_WHILE_DATA_UPDATE",
			}
		}

		return map[string]interface{}{
			"status":        true,
			"message":       "data retrived successfully",
			"matched_count": res.MatchedCount,
		}
	}

	res, err := mgm.CollectionByName("stores").UpdateOne(context.Background(), filter, bson.M{"$set": store})
	if err != nil {
		return map[string]interface{}{
			"status":     false,
			"message":    "error while updating the data",
			"error_code": "ERROR_WHILE_DATA_UPDATE",
		}
	}

	return map[string]interface{}{
		"status":        true,
		"message":       "data retrived successfully",
		"matched_count": res.MatchedCount,
	}
}
