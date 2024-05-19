package controllers

import (
	"letvagas/entities/models"
	"letvagas/web"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	return c.Render("views/index", fiber.Map{
		"title":     "Index",
		"logged_in": error == nil,
		"is_admin":  web.GetRole(c) == models.ADMIN,
	})
}
