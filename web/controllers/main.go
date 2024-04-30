package controllers

import (
	"letvagas/entities/models"
	"letvagas/services"
	"letvagas/web"
	"strings"

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

func Admin(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)
	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}
	return c.Render("views/admin", fiber.Map{
		"title":     "Admin",
		"questions": services.ListAllQuestions(),
		"logged_in": error == nil,
		"is_admin":  web.GetRole(c) == models.ADMIN,
	})
}

func CreateQuestion(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)
	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	err := services.CreateQuestion(c.FormValue("question"), strings.Split(c.FormValue("options"), ","), models.QuestionType(c.FormValue("question_type")))

	if err != nil {
		return c.SendStatus(500)
	}

	return c.Render("views/partials/admin/questions", fiber.Map{
		"questions": services.ListAllQuestions(),
	})
}
