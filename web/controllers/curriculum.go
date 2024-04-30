package controllers

import (
	"letvagas/entities/models"
	"letvagas/services"
	"letvagas/web"

	"github.com/gofiber/fiber/v2"
)

func Curriculum(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)
	educations := services.ListEducations(profile.ID)
	courses := services.ListCourses(profile.ID)
	experiences := services.ListExperiences(profile.ID)

	return c.Render("views/curriculum", fiber.Map{
		"educations":  educations,
		"courses":     courses,
		"experiences": experiences,
		"logged_in":   true,
		"is_admin":    web.GetRole(c) == models.ADMIN,
	})
}
