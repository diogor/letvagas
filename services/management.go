package services

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"

	"github.com/google/uuid"
)

func CreateQuestion(question string, options []string, question_type models.QuestionType) error {

	new_question := models.Question{
		Question: question,
		Options:  options,
		Type:     question_type,
	}

	result := database.DB.Create(&new_question)

	return result.Error
}

func DeleteQuestion(question_id uuid.UUID) error {
	question := models.Question{ID: question_id}

	result := database.DB.Delete(&question)

	return result.Error
}

func ListAllQuestions() []dto.QuestionList {
	questions := []models.Question{}
	result := []dto.QuestionList{}

	database.DB.Find(&questions)

	for _, question := range questions {
		result = append(result, dto.QuestionList{
			QuestionId: question.ID,
			Question:   question.Question,
			Type:       question.Type,
			Options:    question.Options,
		})
	}

	return result
}

func ListAllUsers(page, pageSize int) ([]models.User, int) {
	users := []models.User{}
	var total int64

	database.DB.Model(&models.User{}).Count(&total)
	database.DB.Scopes(Paginate(page, pageSize)).Find(&users)

	return users, int(total)

}

func ChangeUserRole(user_id uuid.UUID, role models.UserRole) {
	database.DB.Model(&models.User{}).Where("id = ?", user_id).Update("role", role)
}

func ListUserStates() []string {
	states := []string{}

	database.DB.Model(&models.User{}).Distinct().Pluck("state", &states)

	return states
}

func ListUserCities(state string) []string {
	cities := []string{}

	database.DB.Model(&models.User{}).Where("state = ?", state).Distinct().Pluck("city", &cities)

	return cities
}

func SearchProfiles(page int, pageSize int, searchParams dto.SearchParams) ([]models.User, int) {
	var profiles []models.User
	var total int64

	query := database.DB

	if searchParams.City != "" {
		query = query.Where("city = ?", searchParams.City)
	}

	if searchParams.State != "" {
		query = query.Where("state = ?", searchParams.State)
	}

	if searchParams.Query != "" {
		query = query.Where(
			"name LIKE ?", "%"+searchParams.Query+"%").Or(
			"profile.goal LIKE ?", "%"+searchParams.Query+"%")
	}

	if searchParams.Neighborhood != "" {
		query = query.Where("neighborhood = ?", searchParams.Neighborhood)
	}

	query.Find(&profiles)

	query.Model(&models.User{}).Count(&total)
	query.Scopes(Paginate(page, pageSize)).Find(&profiles)
	return profiles, int(total)
}
