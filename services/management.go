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

func CreatePosition(position dto.CreatePositionRequest, user_id uuid.UUID) error {

	profile := GetProfile(user_id)

	new_position := models.Position{
		Company:     position.Company,
		Title:       position.Title,
		Level:       position.Level,
		Type:        position.Type,
		Allocation:  position.Allocation,
		Wage:        &position.Wage,
		Contract:    position.Contract,
		Location:    position.Location,
		Description: position.Description,
		PCD:         position.PCD,
		CreatedByID: profile.ID,
	}

	result := database.DB.Create(&new_position)

	return result.Error
}

func GetPositionByID(position_id uuid.UUID) *models.Position {
	position := models.Position{ID: position_id}
	database.DB.First(&position)
	return &position
}

func ListPositions(page, pageSize int) ([]dto.ListPositionResponse, int) {
	var result []dto.ListPositionResponse
	positions := []models.Position{}
	var total int64

	database.DB.Model(&models.Position{}).Count(&total)
	database.DB.Scopes(Paginate(page, pageSize)).Find(&positions)

	for _, pos := range positions {
		result = append(result, dto.ListPositionResponse{
			ID:         pos.ID,
			Title:      pos.Title,
			Company:    pos.Company,
			Level:      pos.GetLevel(),
			Type:       pos.GetType(),
			Allocation: pos.GetAllocation(),
			Wage:       *pos.Wage,
			Contract:   pos.GetContract(),
			Location:   pos.Location,
		})
	}

	return result, int(total)
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

	if searchParams.Query != "" {
		query = query.Joins(
			"left join profiles p on users.profile_id = p.id").Joins(
			"left join experiences e on p.id = e.profile_id").Joins(
			"left join educations ed on e.profile_id = ed.profile_id").Joins(
			"left join courses c on ed.profile_id = c.profile_id").Where(
			"users.name LIKE ?", "%"+searchParams.Query+"%").Or(
			"p.goal LIKE ?", "%"+searchParams.Query+"%").Or(
			"users.social_name LIKE ?", "%"+searchParams.Query+"%").Or(
			"e.activities LIKE ?", "%"+searchParams.Query+"%").Or(
			"ed.institution LIKE ?", "%"+searchParams.Query+"%").Or(
			"ed.graduation LIKE ?", "%"+searchParams.Query+"%").Or(
			"ed.description LIKE ?", "%"+searchParams.Query+"%").Or(
			"c.name LIKE ?", "%"+searchParams.Query+"%").Or(
			"c.description LIKE ?", "%"+searchParams.Query+"%")
	}

	if searchParams.City != "" {
		query = query.Where("city = ?", searchParams.City)
	}

	if searchParams.State != "" {
		query = query.Where("state = ?", searchParams.State)
	}

	if searchParams.EducationType != "" {
		query = query.Where("education_type = ?", searchParams.EducationType)
	}

	if searchParams.Neighborhood != "" {
		query = query.Where("neighborhood = ?", searchParams.Neighborhood)
	}

	query.Find(&profiles)

	query.Model(&models.User{}).Count(&total)
	query.Scopes(Paginate(page, pageSize)).Find(&profiles)
	return profiles, int(total)
}
