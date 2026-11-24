package usecase

import (
	"context"

	"project-courier/internal/dto"
)

type CourierUsecase interface {
	Create(ctx context.Context, request dto.CreateCourierRequest) (*dto.CourierResponse, error)
	GetByID(ctx context.Context, userID string) (*dto.CourierResponse, error)
	GetAll(ctx context.Context, request dto.CourierQueryRequest) (*dto.CourierListResponse, error)
	Update(ctx context.Context, userID string, request dto.UpdateCourierRequest) (*dto.CourierResponse, error)
	Delete(ctx context.Context, userID string) error
}
