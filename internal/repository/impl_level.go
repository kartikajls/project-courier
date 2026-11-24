package repository

import (
	"context"

	"project-courier/internal/entity"

	"gorm.io/gorm"
)

type levelRepository struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) LevelRepository {
	return &levelRepository{
		db: db,
	}
}

// GetAll mengambil seluruh data Level.
func (r *levelRepository) GetAll(ctx context.Context) ([]entity.Level, error) {

	var levels []entity.Level

	err := r.db.WithContext(ctx).
		Order("level_id ASC").
		Find(&levels).Error

	if err != nil {
		return nil, err
	}

	return levels, nil
}

// GetByID mengambil Level berdasarkan level_id.
func (r *levelRepository) GetByID(ctx context.Context, levelID string) (*entity.Level, error) {

	var level entity.Level

	err := r.db.WithContext(ctx).
		Where("level_id = ?", levelID).
		First(&level).Error

	if err != nil {
		return nil, err
	}

	return &level, nil
}
