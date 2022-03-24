package models

import "github.com/kamva/mgm/v3"

type (
	// DB schema - struct to store a data as a format in db
	PosCustomer struct {
		mgm.DefaultModel
		Name          string          `json:"name"`
		Company_name  string          `json:"company_name"`
		Customer_id   string          `json:"customer_id"`
		Email         string          `json:"email"`
		Mobile        string          `json:"mobile"`
		Address       Customeraddress `json:"address"`
		Order_count   string          `json:"order_count"`
		Account_id    string          `json:"account_id"`
		Created_by    string          `json:"created_by"`
		Customer_type string          `json:"customer_type"`
		Pan           string          `json:"pan"`
		Gstin         string          `json:"gstin"`
		Notes         string          `json:"notes"`
		Is_deleted    bool            `json:"is_deleted"`
	}

	Customeraddress struct {
		Zip     string `json:"zip"`
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
	}

	Customer_User_info struct {
		User_id string `json:"user_id"`
	}

	//  To achieve the input payload format
	PosCostomerSchema struct {
		Data      PosCustomer        `json:"data"`
		User_Info Customer_User_info `json:"user_info"`
	}
)
