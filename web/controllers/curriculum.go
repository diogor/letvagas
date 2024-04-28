package controllers

import (
	"letvagas/entities/dto"
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

	return c.Render("views/curriculum", fiber.Map{
		"educations": profile.Educations,
		"logged_in":  true,
	})
}
func CreateEducation(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	new_education := dto.CreateEducationRequest{
		EducationType: models.EducationType(c.FormValue("education_type")),
		Institution:   c.FormValue("institution"),
		Graduation:    c.FormValue("graduation"),
		StartYear:     c.FormValue("start_year"),
		EndYear:       c.FormValue("end_year"),
		StartMonth:    c.FormValue("start_month"),
		EndMonth:      c.FormValue("end_month"),
		Description:   c.FormValue("description"),
		Ongoing:       c.FormValue("ongoing") == "on",
	}

	return services.CreateEducation(profile.ID, &new_education)
}
