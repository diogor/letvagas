package controllers

import (
	"encoding/hex"
	"letvagas/entities/dto"
	"letvagas/entities/models"
	"letvagas/services"
	"letvagas/web"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	if c.Method() == "POST" {
		user := services.GetUserByEmail(c.FormValue("email"))
		password := c.FormValue("password")

		if user == nil || !services.CheckPasswordHash(password, user.Password) {
			return c.Render("views/login", fiber.Map{
				"errors": []string{"Usuário ou senha errados."},
			})
		}

		session, err := web.GetSession(c)
		if err != nil {
			return err
		}

		role_name, _ := user.Role.Value()

		session.Set("user_id", hex.EncodeToString(user.ID[:]))
		session.Set("user_role", role_name)

		err = session.Save()
		if err != nil {
			return err
		}

		return c.Redirect("/")
	}

	return c.Render("views/login", fiber.Map{})
}

func Logout(c *fiber.Ctx) error {
	session, err := web.GetSession(c)
	if err != nil {
		return err
	}

	session.Destroy()

	return c.Redirect("/login")
}

func Register(c *fiber.Ctx) error {
	if c.Method() == "POST" {
		user := dto.CreateUserRequest{
			Name:         c.FormValue("name"),
			SocialName:   c.FormValue("social_name"),
			Password:     c.FormValue("password"),
			Email:        c.FormValue("email"),
			Cpf:          c.FormValue("cpf"),
			Role:         models.APPLICANT,
			BirthDate:    c.FormValue("birth_date"),
			Phone1:       c.FormValue("phone"),
			AreaCode1:    c.FormValue("area_code"),
			Phone2:       c.FormValue("phone2"),
			AreaCode2:    c.FormValue("area_code2"),
			Linkedin:     c.FormValue("linkedin"),
			City:         c.FormValue("city"),
			Neighborhood: c.FormValue("neighborhood"),
			State:        c.FormValue("state"),
			Street:       c.FormValue("street"),
			Number:       c.FormValue("number"),
			Complement:   c.FormValue("complement"),
			Cep:          c.FormValue("cep"),
			PCD:          c.FormValue("pcd") == "on",
		}

		if err := c.BodyParser(&user); err != nil {
			return err
		}

		err := services.CreateUser(&user)

		if err != nil {
			return err
		}

		return c.Redirect("/login")
	}
	return c.Render("views/register", fiber.Map{})
}
