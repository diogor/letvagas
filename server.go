package main

import (
	"letvagas/database"
	"letvagas/entities/models"
	"letvagas/web"
	"letvagas/web/controllers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.InitDB()
	template_engine := django.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: template_engine,
	})

	app.Static("/static", "./static")
	app.Get("/", controllers.Index)
	app.Get("/admin", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.Admin)))
	app.Get("/login", controllers.Login)
	app.Post("/login", controllers.Login)
	app.Get("/register", controllers.Register)
	app.Post("/register", controllers.Register)
	app.Get("/profile", web.LoginRequired(controllers.Curriculum))
	app.Get("/logout", controllers.Logout)
	app.Post("/formation", web.LoginRequired(controllers.CreateEducation))

	app.Get("/partials/educations", web.LoginRequired(controllers.ListEducations))

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
