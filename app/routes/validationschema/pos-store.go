package validationschema

import (

	//validate "gopkg.in/dealancer/validate.v2"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Ierror struct {
	Field string
	Tag   string
	Value interface{}
}

// for create Pos store Validate
func PosStoreValidate(c *fiber.Ctx) error {

	var data CreatePosData

	c.BodyParser(&data)

	err := validate.Struct(&data)
	if err != nil {
		var errors []Ierror

		for _, err := range err.(validator.ValidationErrors) {

			var el Ierror

			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, el)
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status":         false,
			"Message":        "Input vallidation error",
			"Error_response": errors,
		})

	}

	return c.Next()
}

// for update Pos store Validate
func UpdatePosStoreValidate(c *fiber.Ctx) error {

	var data UpdatePosData
	c.BodyParser(&data)

	err := validate.Struct(&data)
	if err != nil {
		var errors []Ierror

		for _, err := range err.(validator.ValidationErrors) {

			var el Ierror

			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, el)
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status":         false,
			"Message":        "Input vallidation error",
			"Error_response": errors,
		})

	}

	return c.Next()
}

// for delete Pos store Validate
func DeletePosStoreValidate(c *fiber.Ctx) error {

	var data MainDeletedata

	c.BodyParser(&data)

	err := validate.Struct(&data)
	if err != nil {
		var errors []Ierror

		for _, err := range err.(validator.ValidationErrors) {

			var el Ierror

			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, el)
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status":         false,
			"Message":        "Input vallidation error",
			"Error_response": errors,
		})
	}

	return c.Next()
}
