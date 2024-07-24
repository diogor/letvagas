package controllers

import (
	"github.com/google/uuid"
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"
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

func DeleteCourse(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	course_id := c.Params("course_id")
	course := models.Course{ID: uuid.MustParse(course_id)}

	database.DB.First(&course)

	if course.ProfileId != profile.ID {
		return c.Redirect("/login")
	}

	services.DeleteCourse(uuid.MustParse(course_id))

	return c.Render("views/partials/courses", fiber.Map{
		"courses": services.ListCourses(profile.ID),
	})

}
