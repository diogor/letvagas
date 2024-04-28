package services

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"

	"github.com/google/uuid"
)

func CreateEducation(profile_id uuid.UUID, education *dto.CreateEducationRequest) error {
	new_ed := models.Education{
		Type:        education.EducationType,
		Institution: education.Institution,
		Graduation:  &education.Graduation,
		StartYear:   &education.StartYear,
		EndYear:     &education.EndYear,
		StartMonth:  &education.StartMonth,
		EndMonth:    &education.EndMonth,
		Description: &education.Description,
		Ongoing:     education.Ongoing,
	}

	result := database.DB.Create(&new_ed).Association("Profile").Append(profile_id)

	return result
}
