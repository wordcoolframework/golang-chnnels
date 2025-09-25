package handlers

import (
	"fiber-gorm-channel-ecommerce/src/application/http/requests"
	"fiber-gorm-channel-ecommerce/src/pkg/validator"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Create(c *fiber.Ctx) error {
	var userReq requests.UserRequest

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"error":   err.Error(),
		})
	}

	messages := map[string]string{
		"Name.required":     "نام الزامی است",
		"Name.min":          "نام باید حداقل ۳ کاراکتر باشد",
		"Email.required":    "ایمیل الزامی است",
		"Email.email":       "ایمیل معتبر نیست",
		"Password.required": "رمز عبور الزامی است",
		"Password.min":      "رمز عبور باید حداقل ۶ کاراکتر باشد",
		"Role.required":     "نقش کاربری الزامی است",
		"Role.min":          "نقش کاربری حداقل 15 کاراکتر باشد",
	}

	if errs := validator.ValidateStruct(userReq, messages); errs != nil {
		return c.Status(422).JSON(fiber.Map{
			"validate_error": errs,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create user successful",
	})
}
