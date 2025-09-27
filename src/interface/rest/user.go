package rest

import (
	"fiber-gorm-channel-ecommerce/src/application/http/handlers"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router *fiber.App) {

	db := databaseCore.Connection()

	//router.Get("/", func(ctx *fiber.Ctx) error {
	//
	//	// var categories []entity.Category
	//	var products []entity.Product
	//	// var user []entity.User
	//
	//	_ = qb.Model(&entity.Product{}).
	//		With("User").
	//		With("User.Orders").
	//		With("Categories").
	//		WhereLikeTable("name", "arash", &entity.User{}). // or "users"
	//		Get(&products)
	//
	//	qb.Gorm().Clauses()
	//	ch := make(chan []entity.Product)
	//
	//	go func() {
	//		ch <- products
	//		close(ch)
	//	}()
	//
	//	result := <-ch
	//
	//	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
	//		"data": result,
	//	})
	//})

	userHandler := handlers.NewUserHandler(db)

	router.Post("/create", userHandler.Create)
	router.Get("/show/:id", userHandler.Show)
}
