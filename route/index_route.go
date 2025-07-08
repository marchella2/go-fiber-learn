package route

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/handler"
)

func RouteInit(r *fiber.App) {
	// routing
	r.Get("/user", handler.GetUsers)
}
