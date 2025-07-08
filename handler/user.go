package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"log"
)

func GetUsers(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}
