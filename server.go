package main

import (
	"log"
	"os"
	"rifa/database"
	"rifa/web"
	"rifa/web/controllers"

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
	app.Get("/", web.LoginRequired(controllers.Index))
	app.Get("/login", controllers.Login)
	app.Post("/login", controllers.Login)
	app.Get("/register", controllers.Register)
	app.Post("/register", controllers.Register)
	app.Get("/logout", controllers.Logout)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
