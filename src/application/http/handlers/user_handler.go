package handlers

import (
	"fiber-gorm-channel-ecommerce/src/application/http/requests"
	"fiber-gorm-channel-ecommerce/src/domain/model/entity"
	"fiber-gorm-channel-ecommerce/src/infrastructure/security"
	"fiber-gorm-channel-ecommerce/src/pkg/qrybldr"
	"net/http"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
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

	return u.db.Transaction(func(tx *gorm.DB) error {

		exists, err := qrybldr.Orm().Query().
			Where("email = ? OR name = ?", userReq.Email, userReq.Name).
			Exists(&entity.User{})

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "خطا در بررسی اطلاعات",
				"error":   err.Error(),
			})
		}

		if exists {
			return c.Status(409).JSON(fiber.Map{
				"message": "کاربری با این ایمیل یا نام قبلا ثبت نام کرده است",
			})
		}

		newUser := entity.User{
			Name:     userReq.Name,
			Email:    userReq.Email,
			Password: security.HashPassword(userReq.Password),
			Role:     userReq.Role,
		}

		if err := qrybldr.Orm().Query().Create(&newUser); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "خطا در ایجاد کاربر",
				"error":   err.Error(),
			})
		}

		userResponse := fiber.Map{
			"id":    newUser.ID,
			"name":  newUser.Name,
			"email": newUser.Email,
			"role":  newUser.Role,
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "کاربر با موفقیت ایجاد شد",
			"data":    userResponse,
		})
	})

}

func (u *UserHandler) Show(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "param id is required !",
		})
	}
	var user entity.User
	// err := u.db.First(&user, id).Error

	qrybldr.Orm().Query().Where("id = ?", id).First(&user)

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": user})

	// return application.RespondRecord(c, err, "User", user)
}
