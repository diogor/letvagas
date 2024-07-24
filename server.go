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
		Views:        template_engine,
		Network:      fiber.NetworkTCP,
		AppName:      "Let vagas 0.1.0",
		ServerHeader: "Let vagas",
	})

	app.Static("/static", "./static")
	app.Static("/media", "./uploads")
	app.Get("/", controllers.Index)
	app.Get("/admin", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.Admin)))
	app.Get("/admin/profile-search", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.SearchProfiles)))
	app.Get("/login", controllers.Login)
	app.Post("/login", controllers.Login)
	app.Get("/register", controllers.Register)
	app.Post("/register", controllers.Register)
	app.Get("/profile", web.LoginRequired(controllers.Curriculum))
	app.Get("/admin/profile/:profile_id", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.Curriculum)))
	app.Get("/logout", controllers.Logout)
	app.Post("/formation", web.LoginRequired(controllers.CreateEducation))
	app.Delete("/formation/:education_id", web.LoginRequired(controllers.DeleteEducation))
	app.Post("/course", web.LoginRequired(controllers.CreateCourse))
	app.Post("/experience", web.LoginRequired(controllers.CreateExperience))
	app.Delete("/experience/:experience_id", web.LoginRequired(controllers.DeleteExperience))
	app.Post("/answers", web.LoginRequired(controllers.CreateAnswer))

	app.Post("/upload-file", web.LoginRequired(controllers.UploadProfileFile))

	app.Post("/questions", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.CreateQuestion)))
	app.Delete("/questions/:question_id", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.DeleteQuestion)))
	app.Put("/users/:user_id", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.ChangeUserRole)))

	app.Get("/partials/educations", web.LoginRequired(controllers.ListEducations))
	app.Get("/partials/courses", web.LoginRequired(controllers.ListCourses))
	app.Get("/partials/experiences", web.LoginRequired(controllers.ListExperiences))
	app.Get("/partials/questions/:question_type", web.LoginRequired(controllers.ListQuestions))
	app.Get("/admin/partials/questions/:question_type/:profile_id", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.ListQuestions)))

	app.Get("/partials/admin/cities", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.ListCitiesByState)))
	app.Get("/partials/admin/users", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.ListAllUsers)))
	app.Get("/partials/admin/search-results", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.SearchResults)))

	app.Post("/profile", web.LoginRequired(controllers.UpdateProfile))

	app.Post("/positions", web.LoginRequired(controllers.CreatePosition))
	app.Get("/positions", web.LoginRequired(controllers.CreatePosition))
	app.Get("/positions/:position_id", controllers.GetPosition)
	app.Delete("/positions/:position_id", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.DeletePosition)))
	app.Get("/positions/:position_id/applications", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.ListUsersForPosition)))
	app.Patch("/positions/:position_id/deactivate", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.DeactivatePosition)))
	app.Patch("/positions/:position_id/activate", web.LoginRequired(web.RoleRequired(models.ADMIN, controllers.ActivatePosition)))
	app.Put("/positions/:position_id/apply", controllers.Apply)

	app.Get("/partials/jobs", controllers.ListPositionsPartial)
	app.Get("/jobs", controllers.ListPositions)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
