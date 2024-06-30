package controllers

import (
	"letvagas/entities/models"
	"letvagas/services"
	"letvagas/web"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Index(c *fiber.Ctx) error {
	uid, error := web.GetUserID(c)

	profile := services.GetProfile(uid)
	applications := services.FindApplicationsByProfile(profile.ID)

	return c.Render("views/index", fiber.Map{
		"title":        "Index",
		"logged_in":    error == nil,
		"is_admin":     web.GetRole(c) == models.ADMIN,
		"applications": applications,
	})
}

func Apply(c *fiber.Ctx) error {
	position_id, _ := uuid.Parse(c.Params("position_id"))

	url := "/positions/" + position_id.String()
	uid, error := web.GetUserID(c)

	if err := error; err != nil {
		return c.Redirect("/login?next=" + url)
	}

	profile_id := services.GetProfile(uid).ID

	_, err := services.CreateApplication(profile_id, position_id)

	if err != nil {
		return c.SendStatus(500)
	}

	return_html := `<div class="pico-background-azure-500"
                     style="padding: 10px;
                            margin-bottom: 20px">
					  <p>Candidatura efetuada!</p>
					</div>`

	return c.SendString(return_html)
}

func MyApplications(c *fiber.Ctx) error {
	uid, error := web.GetUserID(c)
	if err := error; err != nil {
		return c.Redirect("/login")
	}
	profile := services.GetProfile(uid)
	applications := services.FindApplicationsByProfile(profile.ID)
	return c.Render("views/partials/my_applications", fiber.Map{
		"applications": applications,
	})
}
