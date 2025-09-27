package handlers

import (
	"fiber-gorm-channel-ecommerce/src/application/http/requests"
	"fiber-gorm-channel-ecommerce/src/domain/model/entity"
	"fiber-gorm-channel-ecommerce/src/infrastructure/security"
	"fiber-gorm-channel-ecommerce/src/pkg/qrybldr"
	"fiber-gorm-channel-ecommerce/src/pkg/validator"
	"gorm.io/gorm"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	db *gorm.DB
	qb *qrybldr.Qrybldr
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	qb := qrybldr.Instance(db)
	return &UserHandler{
		db: db,
		qb: qb,
	}
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
		"name.required":     "نام الزامی است",
		"name.min":          "نام باید حداقل ۳ کاراکتر باشد",
		"email.required":    "ایمیل الزامی است",
		"email.email":       "ایمیل معتبر نیست",
		"password.required": "رمز عبور الزامی است",
		"password.min":      "رمز عبور باید حداقل ۶ کاراکتر باشد",
		"role.required":     "نقش کاربری الزامی است",
		"role.min":          "نقش کاربری حداقل 15 کاراکتر باشد",
	}

	if errs := validator.ValidateStruct(userReq, messages); errs != nil {
		return c.Status(422).JSON(fiber.Map{
			"validate_error": errs,
		})
	}

	newUser := entity.User{
		Name:     c.Params("name"),
		Email:    c.Params("email"),
		Password: security.HashPassword(c.Params("password")),
		Role:     c.Params("role"),
	}

	_ = u.qb.Create(&newUser)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create user successful",
		"data":    newUser,
	})
}
