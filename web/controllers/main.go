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
	})
}

func Admin(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)
	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}
	return c.Render("views/admin", fiber.Map{
		"title": "Admin",
	})
}
