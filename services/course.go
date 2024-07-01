package services

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"

	"github.com/google/uuid"
)

func CreateCourse(profile_id uuid.UUID, course *dto.CreateCourseRequest) error {
	profile := models.Profile{ID: profile_id}
	database.DB.First(&profile)

	uuidv7, _ := uuid.NewV7()

	new_course := models.Course{
		ID:          uuidv7,
		Name:        course.Name,
		Year:        &course.Year,
		Month:       &course.Month,
		Description: &course.Description,
		Ongoing:     course.Ongoing,
	}

	result := database.DB.Model(&profile).
		Association("Courses").
		Append(&new_course)

	return result
}

func ListCourses(profile_id uuid.UUID) []models.Course {
	courses := []models.Course{}

	profile := models.Profile{ID: profile_id}
	database.DB.First(&profile)

	database.DB.Model(&profile).Association("Courses").Find(&courses)

	return courses
}
