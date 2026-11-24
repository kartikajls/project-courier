package usecase

import (
	"context"

	"project-courier/internal/dto"
)

type StatusUsecase interface {
	GetAll(ctx context.Context) ([]dto.StatusResponse, error)
	GetByType(ctx context.Context, statusType string) (*dto.StatusResponse, error)
}
