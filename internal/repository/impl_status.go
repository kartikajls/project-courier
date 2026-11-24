package repository

import (
	"context"

	"project-courier/internal/entity"

	"gorm.io/gorm"
)

type statusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) StatusRepository {
	return &statusRepository{
		db: db,
	}
}

// GetAll mengambil seluruh data Status.
func (r *statusRepository) GetAll(ctx context.Context) ([]entity.Status, error) {

	var statuses []entity.Status

	err := r.db.WithContext(ctx).
		Order("type ASC").
		Find(&statuses).Error

	if err != nil {
		return nil, err
	}

	return statuses, nil
}

// GetByType mengambil Status berdasarkan type.
func (r *statusRepository) GetByType(ctx context.Context, statusType string) (*entity.Status, error) {

	var status entity.Status

	err := r.db.WithContext(ctx).
		Where("type = ?", statusType).
		First(&status).Error

	if err != nil {
		return nil, err
	}

	return &status, nil
}
