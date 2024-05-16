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
	app.Get("/admin/profile-search", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.SearchProfiles)))
	app.Get("/login", controllers.Login)
	app.Post("/login", controllers.Login)
	app.Get("/register", controllers.Register)
	app.Post("/register", controllers.Register)
	app.Get("/profile", web.LoginRequired(controllers.Curriculum))
	app.Get("/logout", controllers.Logout)
	app.Post("/formation", web.LoginRequired(controllers.CreateEducation))
	app.Post("/course", web.LoginRequired(controllers.CreateCourse))
	app.Post("/experience", web.LoginRequired(controllers.CreateExperience))
	app.Post("/answers", web.LoginRequired(controllers.CreateAnswer))

	app.Post("/question", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.CreateQuestion)))

	app.Get("/partials/educations", web.LoginRequired(controllers.ListEducations))
	app.Get("/partials/courses", web.LoginRequired(controllers.ListCourses))
	app.Get("/partials/experiences", web.LoginRequired(controllers.ListExperiences))
	app.Get("/partials/questions/:question_type", web.LoginRequired(controllers.ListQuestions))

	app.Get("/partials/admin/cities", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.ListCitiesByState)))
	app.Post("/partials/admin/search-results", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.SearchResults)))

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
