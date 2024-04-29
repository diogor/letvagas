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
	educations := services.ListEducations(profile.ID)
	courses := services.ListCourses(profile.ID)

	return c.Render("views/curriculum", fiber.Map{
		"educations": educations,
		"courses":    courses,
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

	services.CreateEducation(profile.ID, &new_education)

	return c.Render("views/partials/educations", fiber.Map{
		"educations": services.ListEducations(profile.ID),
	})

}

func ListEducations(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	educations := services.ListEducations(profile.ID)

	return c.Render("views/partials/educations", fiber.Map{
		"educations": educations,
	})
}

func CreateCourse(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	new_course := dto.CreateCourseRequest{
		Name:        c.FormValue("name"),
		Year:        c.FormValue("year"),
		Month:       c.FormValue("month"),
		Description: c.FormValue("description"),
		Ongoing:     c.FormValue("ongoing") == "on",
	}

	services.CreateCourse(profile.ID, &new_course)

	return c.Render("views/partials/courses", fiber.Map{
		"courses": services.ListCourses(profile.ID),
	})
}

func ListCourses(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	courses := services.ListCourses(profile.ID)

	return c.Render("views/partials/courses", fiber.Map{
		"courses": courses,
	})
}
