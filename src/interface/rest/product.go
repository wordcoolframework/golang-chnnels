package rest

import (
	"fiber-gorm-channel-ecommerce/src/domain/model/entity"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
	"fiber-gorm-channel-ecommerce/src/pkg/qrybldr"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router *fiber.App) {
	router.Get("/", func(ctx *fiber.Ctx) error {

		db := databaseCore.Connection()
		qb := qrybldr.Instance(db)

		// var categories []entity.Category
		var products []entity.Product
		// var user []entity.User

		_ = qb.Model(&entity.Product{}).
			With("User.Orders").
			With("Categories").
			Get(&products)

		qb.Gorm().Clauses()
		ch := make(chan []entity.Product)

		go func() {
			ch <- products
			close(ch)
		}()

		result := <-ch

		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"data": result,
		})
	})
}
