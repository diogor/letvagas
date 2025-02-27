package services

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func CreateExperience(profile_id uuid.UUID, experience *dto.CreateExperienceRequest) error {
	profile := models.Profile{ID: profile_id}
	database.DB.First(&profile)

	var err error

	start_date, err := time.Parse("2006-01-02", experience.StartDate)
	end_date, err := time.Parse("2006-01-02", experience.EndDate)

	if err != nil {
		return err
	}

	uuidv7, _ := uuid.NewV7()

	new_exp := models.Experience{
		ID:                uuidv7,
		Company:           experience.Company,
		StartDate:         datatypes.Date(start_date),
		EndDate:           datatypes.Date(end_date),
		LastWage:          &experience.LastWage,
		Role:              &experience.Role,
		Activities:        &experience.Activities,
		ReferenceContacts: &experience.Reference,
	}

	result := database.DB.Model(&profile).
		Association("Experiences").
		Append(&new_exp)

	return result
}

func ListExperiences(profile_id uuid.UUID) []models.Experience {
	experiences := []models.Experience{}

	profile := models.Profile{ID: profile_id}
	database.DB.First(&profile)

	database.DB.Model(&profile).Association("Experiences").Find(&experiences)

	return experiences
}

func RemoveExperience(experienceId uuid.UUID) error {
	experience := models.Experience{ID: experienceId}
	return database.DB.Delete(&experience).Error
}
