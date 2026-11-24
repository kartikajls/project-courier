package repository

import (
	"context"

	"project-courier/internal/entity"
)

type CourierFilter struct {
	Search  string
	Level   []string
	Status  string
	Page    int
	PerPage int
	SortBy  string
	Order   string
}

type CourierRepository interface {
	Create(ctx context.Context, courier *entity.Courier) error
	GetByID(ctx context.Context, userID string) (*entity.Courier, error)
	GetAll(ctx context.Context, filter CourierFilter) ([]entity.Courier, int64, error)
	Update(ctx context.Context, courier *entity.Courier) error
	Delete(ctx context.Context, userID string) error
}
