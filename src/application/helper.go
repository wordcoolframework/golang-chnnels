package application

import (
	"errors"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func EmptyData() map[string]interface{} {
	return map[string]interface{}{}
}

func RespondRecord(c *fiber.Ctx, err error, entityName string, record interface{}) error {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": entityName + " not found",
				"data":    EmptyData(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "database error",
			"error":   err.Error(),
		})
	}

	val := reflect.ValueOf(record)
	if val.Kind() == reflect.Slice && val.Len() == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": entityName + " not found",
			"data":    EmptyData(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": record,
	})
}
