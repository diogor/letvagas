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

	new_user := models.User{
		Name:     user.Name,
		Password: password,
		Email:    user.Email,
		Role:     user.Role,
	}

	result := database.DB.Create(&new_user)

	return result.Error
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
