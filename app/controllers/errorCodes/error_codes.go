package errorCodes

import (
	"github.com/tidwall/gjson"
)

const Errors = `{
    
    "INVALID_INPUT":{
        "error_code":"INPUT_ERR_0001",
        "description": "error while validating the data or while parsing the data"
    },

    "GST_EXISTS_ALREADY":{
        "error_code":"GST_ERR_0001",
        "description": "gst number not found or already exists"
    },

    "INVALID_ID":{
        "error_code":"SUB_ID",
        "description": "Object with following id doesn't exist."
    },

    "DATABASE_VALIDATION_ERROR":{
        "error_code":"POS_STORE_0001",
        "description": "Validation error(Duplicate/Required filed empty) in Database"
    }, 

    "ERROR_WHILE_DATA_UPDATE": {
        "error_code":"POS_CUSTOMER_0001",
        "description": "Customer Data not found"
    },


    "CUSTOMER_NOT_CREATED" : {
        "error_code":"CUSTOMER_0001",
        "description": "customer details not created"
    },

    "STORE_NOT_CREATED" : {
        "error_code":"POS_STORE_0001",
        "description": "Pos store not created"
    },
    "STORE_NOT_UPDATED":{
        "error_code":"POS_STORE_0001",
        "description": "Pos store not updated"
    },
    "NO_DATA_FOUND" : {
        "error_code":"POS_CUSTOMER_0001",
        "description": "Customer Data not found"
    },
    "NO_SOTRE_DATA_FOUND" : {
        "error_code":"POS_STORE_0001",
        "description": "Store Data not found"
    },
    "DATABASE_ERROR":{
        "error_code":"POS_DATABASE_0001",
        "description": "Invalid data"
    }
}`

// convert the json type error codes into map data type

func Error_codes(s string) map[string]string {
	value := gjson.Get(Errors, s)
	res := value.Map()

	var result = map[string]string{
		"error_code":  res["error_code"].Str,
		"description": res["description"].Str,
	}

	return result
}
