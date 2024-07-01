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

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)
	educations := services.ListEducations(profile.ID)
	courses := services.ListCourses(profile.ID)
	experiences := services.ListExperiences(profile.ID)

	files := services.ListProfileFiles(profile.ID)

	return c.Render("views/curriculum", fiber.Map{
		"educations":  educations,
		"courses":     courses,
		"experiences": experiences,
		"goal":        profile.Goal,
		"files":       files,
		"logged_in":   true,
		"is_admin":    web.GetRole(c) == models.ADMIN,
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

func UpdateProfileGoal(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	services.UpdateProfileGoal(profile.ID, c.FormValue("goal"))
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
