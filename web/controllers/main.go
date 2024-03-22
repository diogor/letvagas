package controllers

import (
	"letvagas/web"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)

	if error != nil {
		return error
	}

	return c.Render("views/index", fiber.Map{
		"title": "Index",
	})
}
