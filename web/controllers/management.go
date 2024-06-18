package controllers

import (
	"letvagas/entities/dto"
	"letvagas/entities/models"
	"letvagas/services"
	"letvagas/web"
	"strconv"
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

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	users, total := services.ListAllUsers(page, pageSize)

	var pageRange []int
	for i := 0; i < (total / pageSize); i++ {
		pageRange = append(pageRange, i)
	}

	return c.Render("views/admin", fiber.Map{
		"title":            "Admin",
		"questions":        services.ListAllQuestions(),
		"users":            users,
		"users_total":      total,
		"users_page":       page,
		"users_page_size":  pageSize,
		"users_page_range": pageRange,
		"logged_in":        error == nil,
		"is_admin":         web.GetRole(c) == models.ADMIN,
		"user_id":          userId,
	})
}

func searchProfilesService(c *fiber.Ctx) ([]models.User, int, dto.SearchParams, string) {

	city := c.Query("city")
	state := c.Query("state")
	neighborhood := c.Query("neighborhood")

	q := c.Query("q")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	params := dto.SearchParams{
		Query:        q,
		City:         city,
		State:        state,
		Neighborhood: neighborhood,
	}

	profiles, total := services.SearchProfiles(page, pageSize, params)

	querySttring := "&q=" + params.Query + "&state=" + params.State + "&city=" + params.City

	return profiles, total, params, querySttring
}

func SearchProfiles(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)

	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	profiles, total, _, querySttring := searchProfilesService(c)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	var pageRange []int
	for i := 0; i < (total / pageSize); i++ {
		pageRange = append(pageRange, i)
	}

	return c.Render("views/admin/profile_search", fiber.Map{
		"title":        "Busca de Perfis",
		"logged_in":    error == nil,
		"is_admin":     web.GetRole(c) == models.ADMIN,
		"states":       services.ListUserStates(),
		"users":        profiles,
		"total":        total,
		"page":         page,
		"page_size":    pageSize,
		"page_range":   pageRange,
		"query_string": querySttring,
	})
}

func SearchResults(c *fiber.Ctx) error {
	_, error := web.GetUserID(c)
	role := web.GetRole(c)

	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	profiles, total, params, querySttring := searchProfilesService(c)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	var pageRange []int
	for i := 0; i < (total / pageSize); i++ {
		pageRange = append(pageRange, i)
	}

	return c.Render("views/partials/admin/search_results", fiber.Map{
		"users":        profiles,
		"total":        total,
		"page":         page,
		"page_size":    pageSize,
		"page_range":   pageRange,
		"query_params": params,
		"query_string": querySttring,
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

func CreatePosition(c *fiber.Ctx) error {
	user_id, error := web.GetUserID(c)
	role := web.GetRole(c)
	if error != nil || role != models.ADMIN {
		return c.Redirect("/login")
	}

	if c.Method() == "GET" {
		return c.Render("views/admin/add_position", fiber.Map{
			"errors": nil,
		})
	}

	created_by_id := user_id
	position := dto.CreatePositionRequest{
		Company:     c.FormValue("company"),
		Title:       c.FormValue("title"),
		Type:        models.PositionType(c.FormValue("type")),
		Allocation:  models.Allocation(c.FormValue("allocation")),
		Wage:        c.FormValue("wage"),
		Contract:    models.ContractType(c.FormValue("contract")),
		Location:    c.FormValue("location"),
		Description: c.FormValue("description"),
		PCD:         c.FormValue("pcd") == "on",
	}

	services.CreatePosition(position, created_by_id)

	return c.SendStatus(201)
}

func ListCitiesByState(c *fiber.Ctx) error {
	state := c.Query("state")
	return c.Render("views/partials/admin/cities", fiber.Map{
		"cities": services.ListUserCities(state),
	})
}

func ListAllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	users, total := services.ListAllUsers(page, pageSize)
	var pageRange []int

	for i := 0; i < (total / pageSize); i++ {
		pageRange = append(pageRange, i)
	}

	userId, _ := web.GetUserID(c)

	return c.Render("views/partials/admin/users", fiber.Map{
		"users":      users,
		"user_id":    userId,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"page_range": pageRange,
	})
}
