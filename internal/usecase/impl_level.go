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

type levelUsecase struct {
	levelRepo repository.LevelRepository
}

func NewLevelUsecase(levelRepo repository.LevelRepository) LevelUsecase {
	return &levelUsecase{
		levelRepo: levelRepo,
	}
}

// GET ALL LEVELS
func (u *levelUsecase) GetAll(ctx context.Context) ([]dto.LevelResponse, error) {

	levels, err := u.levelRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.LevelResponse, 0, len(levels))

	for _, level := range levels {
		responses = append(responses, *toLevelResponse(&level))
	}

	return responses, nil
}

// GET LEVEL BY ID
func (u *levelUsecase) GetByID(ctx context.Context, levelID string) (*dto.LevelResponse, error) {

	levelID = strings.TrimSpace(levelID)
	if levelID == "" {
		return nil, ErrInvalidInput
	}

	level, err := u.levelRepo.GetByID(ctx, levelID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return toLevelResponse(level), nil
}

// ENTITY TO RESPONSE DTO
func toLevelResponse(level *entity.Level) *dto.LevelResponse {
	return &dto.LevelResponse{
		LevelID:     level.LevelID,
		Name:        level.Name,
		Description: level.Description,
	}
}
