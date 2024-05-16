package services

import (
	"letvagas/database"
	"letvagas/entities/dto"
	"letvagas/entities/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(user *dto.CreateUserRequest) error {
	password, err := hashPassword(user.Password)

	if err != nil {
		return err
	}

	new_user := &models.User{
		Name:         user.Name,
		SocialName:   &user.SocialName,
		Password:     password,
		Email:        user.Email,
		Role:         models.APPLICANT,
		Cpf:          user.Cpf,
		BirthDate:    user.BirthDate,
		City:         &user.City,
		State:        &user.State,
		Cep:          &user.Cep,
		Street:       &user.Street,
		Neighborhood: &user.Neighborhood,
		Number:       &user.Number,
		Complement:   &user.Complement,
		Phone1:       user.Phone1,
		AreaCode1:    user.AreaCode1,
		Phone2:       &user.Phone2,
		AreaCode2:    &user.AreaCode2,
		Linkedin:     &user.Linkedin,
		PCD:          user.PCD,
		Profile:      models.Profile{ID: uuid.New()},
	}

	result := database.DB.Create(new_user)

	return result.Error
}

func GetProfile(user_id uuid.UUID) *models.Profile {
	user := models.User{ID: user_id}
	database.DB.First(&user).Association("Profile")
	profile := models.Profile{ID: user.ProfileID}
	database.DB.First(&profile)

	return &profile
}

func GetUser(id string) *models.User {
	user := models.User{ID: uuid.MustParse(id)}
	database.DB.First(&user)
	return &user
}

func GetUserByEmail(email string) *models.User {
	user := models.User{Email: email}
	database.DB.First(&user, "email = ?", email)

	if user.ID == uuid.Nil {
		return nil
	}

	return &user
}
