package repository

import (
	"context"
	"fmt"
	"strings"

	"project-courier/internal/entity"

	"gorm.io/gorm"
)

type courierRepository struct {
	db *gorm.DB
}

func NewCourierRepository(db *gorm.DB) CourierRepository {
	return &courierRepository{
		db: db,
	}
}

// Create: menyimpan data Courier baru.
func (r *courierRepository) Create(ctx context.Context, courier *entity.Courier) error {
	return r.db.WithContext(ctx).Create(courier).Error
}

// GetByID mengambil data Courier berdasarkan user_id.
func (r *courierRepository) GetByID(ctx context.Context, userID string) (*entity.Courier, error) {
	var courier entity.Courier

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&courier).Error

	if err != nil {
		return nil, err
	}

	return &courier, nil
}

// GetAll: mengambil daftar Courier dengan filter dan pagination.
func (r *courierRepository) GetAll(ctx context.Context, filter CourierFilter) ([]entity.Courier, int64, error) {

	var couriers []entity.Courier
	var totalItems int64

	query := r.db.WithContext(ctx).Model(&entity.Courier{})

	// Search berdasarkan setiap kata.
	if filter.Search != "" {
		keywords := strings.Fields(filter.Search)

		for _, keyword := range keywords {
			query = query.Where(
				"name LIKE ?",
				"%"+keyword+"%",
			)
		}
	}

	// Filter berdasarkan Level.
	if len(filter.Level) > 0 {
		query = query.Where("level IN ?", filter.Level)
	}

	// Filter berdasarkan Status.
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	// Menghitung total data sebelum pagination.
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	// Default pagination.
	page := filter.Page
	if page < 1 {
		page = 1
	}

	perPage := filter.PerPage
	if perPage < 1 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	// Default sorting.
	sortBy := filter.SortBy
	if sortBy == "" {
		sortBy = "name"
	}

	order := strings.ToUpper(filter.Order)
	if order != "DESC" {
		order = "ASC"
	}

	// Whitelist kolom sorting.
	allowedSort := map[string]string{
		"name":       "name",
		"created_at": "created_at",
	}

	column, exists := allowedSort[sortBy]
	if !exists {
		return nil, 0, fmt.Errorf("invalid sort column")
	}

	// Mengambil data sesuai filter, sorting, dan pagination.
	err := query.
		Order(column + " " + order).
		Limit(perPage).
		Offset(offset).
		Find(&couriers).Error

	if err != nil {
		return nil, 0, err
	}

	return couriers, totalItems, nil
}

// Update: memperbarui data Courier.
func (r *courierRepository) Update(ctx context.Context, courier *entity.Courier) error {

	result := r.db.WithContext(ctx).
		Model(&entity.Courier{}).
		Where("user_id = ?", courier.UserID).
		Select(
			"name",
			"email",
			"phone",
			"address",
			"level",
			"status",
		).
		Updates(courier)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Delete: menghapus data Courier berdasarkan user_id.
func (r *courierRepository) Delete(ctx context.Context, userID string) error {

	result := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.Courier{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
