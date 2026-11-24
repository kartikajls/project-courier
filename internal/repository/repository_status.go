package repository

import (
	"context"

	"project-courier/internal/entity"
)

type StatusRepository interface {
	GetAll(ctx context.Context) ([]entity.Status, error)
	GetByType(ctx context.Context, statusType string) (*entity.Status, error)
}
