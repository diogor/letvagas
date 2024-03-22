package dto

import "letvagas/entities/models"

type CreateUserRequest struct {
	Name     string           `json:"name" binding:"required"`
	Password string           `json:"password" binding:"required"`
	Email    string           `json:"email" binding:"required"`
	Role     *models.UserRole `json:"role" binding:"required"`
}
