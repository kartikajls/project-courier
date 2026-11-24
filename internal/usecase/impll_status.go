package usecase

import (
	"context"
	"errors"
	"strings"

	"project-courier/internal/dto"
	"project-courier/internal/entity"
	"project-courier/internal/repository"

	"gorm.io/gorm"
)

type statusUsecase struct {
	statusRepo repository.StatusRepository
}

func NewStatusUsecase(statusRepo repository.StatusRepository) StatusUsecase {
	return &statusUsecase{
		statusRepo: statusRepo,
	}
}

// GET ALL STATUSES
func (u *statusUsecase) GetAll(ctx context.Context) ([]dto.StatusResponse, error) {

	statuses, err := u.statusRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.StatusResponse, 0, len(statuses))

	for _, status := range statuses {
		responses = append(responses, *toStatusResponse(&status))
	}

	return responses, nil
}

// GET STATUS BY TYPE
func (u *statusUsecase) GetByType(ctx context.Context, statusType string) (*dto.StatusResponse, error) {

	statusType = strings.TrimSpace(statusType)
	if statusType == "" {
		return nil, ErrInvalidInput
	}

	status, err := u.statusRepo.GetByType(ctx, statusType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return toStatusResponse(status), nil
}

// ENTITY TO RESPONSE DTO
func toStatusResponse(status *entity.Status) *dto.StatusResponse {
	return &dto.StatusResponse{
		Type: status.Type,
	}
}
