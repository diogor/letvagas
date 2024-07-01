package controllers

import (
	"letvagas/entities/dto"
	"letvagas/entities/models"
	"letvagas/services"
	"letvagas/web"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Curriculum(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)
	var profile_id uuid.UUID
	var profile models.Profile

	if err != nil {
		return c.Redirect("/login")
	}

	if c.Params("profile_id") != "" {
		profile = *services.GetProfileById(uuid.MustParse(c.Params("profile_id")))
		profile_id = uuid.MustParse(c.Params("profile_id"))
	} else {
		profile = *services.GetProfile(user_id)
		profile_id = profile.ID
	}

	educations := services.ListEducations(profile_id)
	courses := services.ListCourses(profile_id)
	experiences := services.ListExperiences(profile_id)

	files := services.ListProfileFiles(profile_id)

	return c.Render("views/curriculum", fiber.Map{
		"educations":       educations,
		"courses":          courses,
		"experiences":      experiences,
		"goal":             profile.Goal,
		"wage_expectation": profile.WageExpectation,
		"files":            files,
		"logged_in":        true,
		"is_admin":         web.GetRole(c) == models.ADMIN,
		"profile_id":       profile_id,
	})
}

func CreateAnswer(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	new_answer := dto.CreateAnswerRequest{
		QuestionId: uuid.MustParse(c.FormValue("question_id")),
		Answer:     c.FormValue("answer"),
	}

	err = services.CreateAnswer(profile.ID, &new_answer)

	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(204)

}

func UpdateProfile(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	goal := c.FormValue("goal")
	wage_expectation := c.FormValue("wage_expectation")

	profile := services.GetProfile(user_id)

	services.UpdateProfile(profile.ID, &goal, &wage_expectation)
	return c.SendStatus(204)

}

func UploadProfileFile(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	name := c.FormValue("name")
	file, err := c.FormFile("file")

	if err != nil {
		return c.SendStatus(500)
	}

	filename := uuid.NewString() + "." + strings.Split(file.Filename, ".")[1]

	c.SaveFile(file, "./uploads/"+filename)

	services.SaveProfileFile(profile.ID, name, &filename, nil)
	return c.SendStatus(204)
}
