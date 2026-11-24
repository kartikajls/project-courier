package repository

import (
	"context"

	"project-courier/internal/entity"
)

type LevelRepository interface {
	GetAll(ctx context.Context) ([]entity.Level, error)
	GetByID(ctx context.Context, levelID string) (*entity.Level, error)
}
