package controllers

import (
	"letvagas/entities/dto"
	"letvagas/services"
	"letvagas/web"

	"github.com/gofiber/fiber/v2"
)

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
