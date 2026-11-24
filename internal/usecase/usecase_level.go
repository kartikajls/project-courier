package usecase

import (
	"context"

	"project-courier/internal/dto"
)

type LevelUsecase interface {
	GetAll(ctx context.Context) ([]dto.LevelResponse, error)
	GetByID(ctx context.Context, levelID string) (*dto.LevelResponse, error)
}
