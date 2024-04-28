package database

import (
	"letvagas/entities/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	godotenv.Load()
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&models.User{}, &models.Profile{}, &models.Education{},
		&models.Course{}, &models.ComputingSkill{}, &models.ComputingSkillAnswer{},
		&models.ComputingQuestion{}, &models.ComputingAnswer{},
		&models.LanguageQuestion{}, &models.LanguageAnswer{}, &models.Course{},
		&models.Experience{},
	)
}
