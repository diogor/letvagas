package controllers

import (
	"encoding/hex"
	"letvagas/entities/dto"
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
			Name:     c.FormValue("name"),
			Password: c.FormValue("password"),
			Email:    c.FormValue("email"),
		}

		if err := c.BodyParser(&user); err != nil {
			return err
		}

		go services.CreateUser(&user)

		return c.Redirect("/login")
	}
	return c.Render("views/register", fiber.Map{})
}
