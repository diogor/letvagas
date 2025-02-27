package controllers

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"
	"letvagas/services"
	"letvagas/web"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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

func DeleteEducation(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	education_id := c.Params("education_id")
	education := models.Education{ID: uuid.MustParse(education_id)}

	database.DB.First(&education)

	if education.ProfileId != profile.ID {
		return c.Redirect("/login")
	}

	services.DeleteEducation(uuid.MustParse(education_id))

	return c.Render("views/partials/educations", fiber.Map{
		"educations": services.ListEducations(profile.ID),
	})
}

func DeleteExperience(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)

	experience_id := c.Params("experience_id")
	experience := models.Experience{ID: uuid.MustParse(experience_id)}

	database.DB.First(&experience)

	if experience.ProfileId != profile.ID {
		return c.Redirect("/login")
	}

	services.RemoveExperience(uuid.MustParse(experience_id))

	return c.Render("views/partials/experiences", fiber.Map{
		"experiences": services.ListExperiences(profile.ID),
	})
}

func ListQuestions(c *fiber.Ctx) error {
	user_id, err := web.GetUserID(c)

	if err != nil {
		return c.Redirect("/login")
	}

	profile := services.GetProfile(user_id)
	profile_id := profile.ID

	if c.Params("profile_id") != "" {
		profile_id = uuid.MustParse(c.Params("profile_id"))
	}

	question_type := c.Params("question_type")
	return c.Render("views/partials/questions", fiber.Map{
		"questions": services.ListQuestions(models.QuestionType(question_type), profile_id),
		"answers":   services.ListAnswers(profile_id),
	})
}
