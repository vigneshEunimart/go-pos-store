package validationschema

//create pos store validate schema
type (
	CreatePosStoresSchema struct {
		Store_id            string `json:"store_id"`
		Store_name          string `json:"store_name" validate:"required"`
		Account_id          string `json:"account_id" validate:"required"`
		Inventory_type      string `json:"inventory_type" validate:"oneof='center_inventory' 'decenter_inventory' ''"` //enum  default : "decenter_inventory"
		Store_phone_number  string `json:"store_phone_number" validate:"required"`
		Business_reg_number string `json:"business_reg_number" validate:"required"`
		Gst_number          string `json:"gst_number" validate:"required"`
		Iec_number          string `json:"iec_number" validate:"required"`
		Gst_file            string `json:"gst_file" validate:"required"`
		Iec_file            string `json:"iec_file" validate:"required"`
		Currency            string `json:"currency" validate:"gte=0"`                                      //default : "inr"
		Transition          string `json:"transition" validate:"oneof='Inclusive tax' 'Exclusive tax' ''"` //enum  default : "inclusive tax"
		Address             Addr   `json:"address" validate:"required"`
		Manager_name        string `json:"manager_name,omitempty"`
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

	CreatePosData struct {
		Data      CreatePosStoresSchema `json:"data"`
		User_info User_Info             `json:"user_info"`
	}
)

// update pos store validtae schema
type (
	UpdatePosStoresSchema struct {
		Store_id            string `json:"store_id" validate:"required"`
		Store_name          string `json:"store_name" validate:"required"`
		Account_id          string `json:"account_id" validate:"required"`
		Inventory_type      string `json:"inventory_type" validate:"oneof='center_inventory' 'decenter_inventory'"` //enum  default : "decenter_inventory"
		Store_phone_number  string `json:"store_phone_number" validate:"required"`
		Business_reg_number string `json:"business_reg_number" validate:"required"`
		Gst_number          string `json:"gst_number" validate:"required"`
		Iec_number          string `json:"iec_number" validate:"required"`
		Gst_file            string `json:"gst_file" validate:"required"`
		Iec_file            string `json:"iec_file" validate:"required"`
		Currency            string `json:"currency" validate:"required"`                                //default : "inr"
		Transition          string `json:"transition" validate:"oneof='Inclusive tax' 'Exclusive tax'"` //enum  default : "inclusive tax"
		Address             Addre  `json:"address" validate:"required"`

		Manager_name string `json:"manager_name" validate:"required"`

		Manager_phone string `json:"manager_phone" validate:"required"`
		Is_active     bool   `json:"is_active" validate:"required"`
		Is_deleted    bool   `json:"is_deleted"`
		Updated_by    string `json:"updated_by"`
		Created_by    string `json:"created_by"`
	}
	Addre struct {
		Zip     string `json:"zip" validate:"required"`
		City    string `json:"city" validate:"required"`
		State   string `json:"state" validate:"required"`
		Country string `json:"country" validate:"required"`
	}

	UpdateUser_Info struct {
		UserId string `json:"user_id" validate:"required"`
	}

	UpdatePosData struct {
		Data      UpdatePosStoresSchema `json:"data"`
		User_info User_Info             `json:"user_info"`
	}
)

// Delete pos store validate schema
type (
	MainDeletedata struct {
		Delete Deletedata `json:"data"`
	}

	Deletedata struct {
		Account_id string `json:"account_id" validate:"required"`
		Store_id   string `json:"store_id" validate:"required"`
	}
)
