package web

import (
	"fmt"
	"letvagas/entities/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
)

var store *session.Store = session.New()

func GetSession(c *fiber.Ctx) (*session.Session, error) {
	sess, err := store.Get(c)

	if err != nil {
		return nil, err
	}

	return sess, err
}

func LoginRequired(handler func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := GetSession(c)
		if sess.Get("user_id") == nil {
			return c.Redirect("/login")
		}
		handler(c)
		return err
	}
}

func GetUserID(c *fiber.Ctx) (uuid.UUID, error) {
	session, _ := GetSession(c)
	user_id := fmt.Sprintf("%v", session.Get("user_id"))
	user_uuid, error := uuid.Parse(user_id)

	return user_uuid, error
}

func GetRole(c *fiber.Ctx) models.UserRole {
	session, _ := GetSession(c)
	role := fmt.Sprintf("%v", session.Get("user_role"))
	user_role := models.UserRole(role)

	return user_role
}

func RoleRequired(role models.UserRole, handler func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, _ := GetSession(c)
		role_name, _ := role.Value()
		if sess.Get("user_role") != role_name {
			return c.Redirect("/login")
		}
		handler(c)
		return nil
	}
}
