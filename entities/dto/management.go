package dto

import (
	"letvagas/entities/models"

	"github.com/google/uuid"
)

type CreatePositionRequest struct {
	Company     string              `json:"company" binding:"required"`
	Title       string              `json:"title" binding:"required"`
	Level       models.Level        `json:"level" binding:"required"`
	Type        models.PositionType `json:"type" binding:"required"`
	Allocation  models.Allocation   `json:"allocation" binding:"required"`
	Wage        string              `json:"wage"`
	Contract    models.ContractType `json:"contract" binding:"required"`
	Location    string              `json:"location"`
	Description string              `json:"description"`
	PCD         bool                `json:"pcd"`
}

type PositionResponse struct {
	CreatePositionRequest
	ID uuid.UUID `json:"id"`
}

type ListPositionResponse struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Company    string    `json:"company"`
	Location   string    `json:"location"`
	Level      string    `json:"level"`
	Type       string    `json:"type"`
	Allocation string    `json:"allocation"`
	Contract   string    `json:"contract"`
	Wage       string    `json:"wage"`
	PCD        bool      `json:"pcd"`
}
