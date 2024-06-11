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
		phone1 := c.FormValue("phone")
		phone2 := c.FormValue("phone2")
		phone3 := c.FormValue("phone3")

		user := dto.CreateUserRequest{
			Name:         c.FormValue("name"),
			SocialName:   c.FormValue("social_name"),
			Password:     c.FormValue("password"),
			Email:        c.FormValue("email"),
			Cpf:          c.FormValue("cpf"),
			Role:         models.APPLICANT,
			BirthDate:    c.FormValue("birth_date"),
			Phone1:       phone1,
			AreaCode1:    c.FormValue("area_code"),
			Phone2:       phone2,
			AreaCode2:    c.FormValue("area_code2"),
			Phone3:       phone3,
			AreaCode3:    c.FormValue("area_code3"),
			Linkedin:     c.FormValue("linkedin"),
			City:         c.FormValue("city"),
			Neighborhood: c.FormValue("neighborhood"),
			State:        c.FormValue("state"),
			Street:       c.FormValue("street"),
			Number:       c.FormValue("number"),
			Complement:   c.FormValue("complement"),
			Cep:          c.FormValue("cep"),
			PCD:          c.FormValue("pcd") == "on",
			PCDInfo:      c.FormValue("pcd_info"),
		}

		var errors []string

		existing_user := services.GetUserByEmail(user.Email)

		if existing_user != nil {
			errors = append(errors, "O email informado já está em uso")
		}

		if phone1 == "" && phone2 == "" && phone3 == "" {
			errors = append(errors, "Preencha pelo menos um telefone")
		}

		if len(errors) > 0 {
			return c.Render("views/register", fiber.Map{
				"errors": errors,
				"user":   user,
			})
		}

		if err := c.BodyParser(&user); err != nil {
			return err
		}

		uid, err := services.CreateUser(&user)

		if err != nil {
			return err
		}

		session, err := web.GetSession(c)
		if err != nil {
			return err
		}

		role_name, _ := user.Role.Value()

		session.Set("user_id", hex.EncodeToString(uid[:]))
		session.Set("user_role", role_name)

		err = session.Save()
		if err != nil {
			return err
		}
		return c.Redirect("/profile")
	}
	return c.Render("views/register", fiber.Map{})
}
