package controllers

import (
	"letvagas/entities/models"
	"letvagas/services"
	"letvagas/web"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Admin(c *fiber.Ctx) error {
	userId, error := web.GetUserID(c)
	role := web.GetRole(c)
	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}
	return c.Render("views/admin", fiber.Map{
		"title":     "Admin",
		"questions": services.ListAllQuestions(),
		"users":     services.ListAllUsers(),
		"logged_in": error == nil,
		"is_admin":  web.GetRole(c) == models.ADMIN,
		"user_id":   userId,
	})
}

func SearchProfiles(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)

	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	city := c.Query("city")
	state := c.Query("state")

	q := c.Query("q")

	profiles := services.SearchProfiles(q, city, state)

	return c.Render("views/admin/profile_search", fiber.Map{
		"title":     "Busca de Perfis",
		"logged_in": error == nil,
		"is_admin":  web.GetRole(c) == models.ADMIN,
		"states":    services.ListUserStates(),
		"users":     profiles,
	})
}

func SearchResults(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)

	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	city := c.FormValue("city")
	state := c.FormValue("state")

	q := c.FormValue("q")

	profiles := services.SearchProfiles(q, city, state)

	return c.Render("views/partials/admin/search_results", fiber.Map{
		"users": profiles,
	})
}

func CreateQuestion(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)
	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	options := []string{}

	if c.FormValue("options") != "" {
		options = strings.Split(c.FormValue("options"), ",")
	}

	err := services.CreateQuestion(c.FormValue("question"), options, models.QuestionType(c.FormValue("question_type")))

	if err != nil {
		return c.SendStatus(500)
	}

	return c.Render("views/partials/admin/questions", fiber.Map{
		"questions": services.ListAllQuestions(),
	})
}

func DeleteQuestion(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)
	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	question_id := uuid.MustParse(c.Params("question_id"))
	services.DeleteQuestion(question_id)

	return c.Render("views/partials/admin/questions", fiber.Map{
		"questions": services.ListAllQuestions(),
	})
}

func ChangeUserRole(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)
	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	user_id := uuid.MustParse(c.Params("user_id"))
	services.ChangeUserRole(user_id, models.UserRole(c.FormValue("role")))

	return c.SendStatus(204)
}

func ListCitiesByState(c *fiber.Ctx) error {
	state := c.Query("state")
	return c.Render("views/partials/admin/cities", fiber.Map{
		"cities": services.ListUserCities(state),
	})
}
