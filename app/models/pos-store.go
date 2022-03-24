package models

import "github.com/kamva/mgm/v3"

// DB schema - struct to store a data as a format in db
type PosStoresSchema struct {
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

type Addr struct {
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
}

type User_Info struct {
	UserId string `json:"user_id" validate:"required"`
}

// To achieve the input payload format
type Data struct {
	Data      PosStoresSchema `json:"data"`
	User_info User_Info       `json:"user_info"`
}
