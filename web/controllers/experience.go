package controllers

import (
	"letvagas/entities/dto"
	"letvagas/services"
	"letvagas/web"

	"github.com/gofiber/fiber/v2"
)

func CreateExperience(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	new_experience := dto.CreateExperienceRequest{
		Company:    c.FormValue("company"),
		StartDate:  c.FormValue("start_date"),
		EndDate:    c.FormValue("end_date"),
		LastWage:   c.FormValue("last_wage"),
		Role:       c.FormValue("role"),
		Activities: c.FormValue("activities"),
		Reference:  c.FormValue("reference_contacts"),
	}

	err = services.CreateExperience(profile.ID, &new_experience)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.Render("views/partials/experiences", fiber.Map{
		"experiences": services.ListExperiences(profile.ID),
	})
}

func ListExperiences(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	experiences := services.ListExperiences(profile.ID)

	return c.Render("views/partials/experiences", fiber.Map{
		"experiences": experiences,
	})
}
