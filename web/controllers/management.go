package controllers

import (
	"letvagas/services"

	"github.com/gofiber/fiber/v2"
)

func ListCitiesByState(c *fiber.Ctx) error {
	state := c.Query("state")
	return c.Render("views/partials/admin/cities", fiber.Map{
		"cities": services.ListUserCities(state),
	})
}
