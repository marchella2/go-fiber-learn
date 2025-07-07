package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/database"
	"go-fiber-gorm/migration"
	"go-fiber-gorm/route"
)

func main() {
	// initial database
	database.DatabaseInit()

	// run migration
	migration.RunMigrations()

	// init fiber
	app := fiber.New()

	// initial route
	route.RouteInit(app)

	app.Listen(":3030")
}
