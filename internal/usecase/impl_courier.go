package usecase

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/mail"
	"strings"

	"project-courier/internal/dto"
	"project-courier/internal/entity"
	"project-courier/internal/repository"

	"gorm.io/gorm"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("data not found")
)

type courierUsecase struct {
	courierRepo repository.CourierRepository
	levelRepo   repository.LevelRepository
	statusRepo  repository.StatusRepository
}

func NewCourierUsecase(
	courierRepo repository.CourierRepository,
	levelRepo repository.LevelRepository,
	statusRepo repository.StatusRepository,
) CourierUsecase {
	return &courierUsecase{
		courierRepo: courierRepo,
		levelRepo:   levelRepo,
		statusRepo:  statusRepo,
	}
}

// CREATE COURIER
func (u *courierUsecase) Create(ctx context.Context, request dto.CreateCourierRequest) (*dto.CourierResponse, error) {

	request.UserID = strings.TrimSpace(request.UserID)
	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.TrimSpace(request.Email)
	request.Phone = strings.TrimSpace(request.Phone)
	request.Address = strings.TrimSpace(request.Address)
	request.Level = strings.TrimSpace(request.Level)
	request.Status = strings.TrimSpace(request.Status)

	if request.UserID == "" ||
		request.Name == "" ||
		request.Email == "" ||
		request.Phone == "" ||
		request.Address == "" ||
		request.Level == "" {
		return nil, fmt.Errorf("%w: required fields must be filled", ErrInvalidInput)
	}

	if _, err := mail.ParseAddress(request.Email); err != nil {
		return nil, fmt.Errorf("%w: invalid email format", ErrInvalidInput)
	}

	if len(request.Phone) > 20 {
		return nil, fmt.Errorf("%w: phone maximum is 20 characters", ErrInvalidInput)
	}

	// Check duplicate Courier ID
	_, err := u.courierRepo.GetByID(ctx, request.UserID)
	if err == nil {
		return nil, fmt.Errorf("%w: courier ID already exists", ErrInvalidInput)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Validate Level
	_, err = u.levelRepo.GetByID(ctx, request.Level)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: level does not exist", ErrInvalidInput)
	}
	if err != nil {
		return nil, err
	}

	// Default Status
	if request.Status == "" {
		request.Status = "active"
	}

	// Validate Status
	_, err = u.statusRepo.GetByType(ctx, request.Status)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: status does not exist", ErrInvalidInput)
	}
	if err != nil {
		return nil, err
	}

	courier := &entity.Courier{
		UserID:  request.UserID,
		Name:    request.Name,
		Email:   request.Email,
		Phone:   request.Phone,
		Address: request.Address,
		Level:   request.Level,
		Status:  request.Status,
	}

	if err := u.courierRepo.Create(ctx, courier); err != nil {
		return nil, err
	}

	return toCourierResponse(courier), nil
}

// GET COURIER BY ID
func (u *courierUsecase) GetByID(ctx context.Context, userID string) (*dto.CourierResponse, error) {

	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil, fmt.Errorf("%w: courier ID is required", ErrInvalidInput)
	}

	courier, err := u.courierRepo.GetByID(ctx, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return toCourierResponse(courier), nil
}

// GET ALL COURIERS
func (u *courierUsecase) GetAll(ctx context.Context, request dto.CourierQueryRequest) (*dto.CourierListResponse, error) {

	page := request.Page
	if page < 1 {
		page = 1
	}

	perPage := request.PerPage
	if perPage < 1 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100
	}

	sortBy := strings.TrimSpace(request.Sort)
	if sortBy == "" {
		sortBy = "name"
	}

	if sortBy != "name" && sortBy != "created_at" {
		return nil, fmt.Errorf("%w: invalid sort field", ErrInvalidInput)
	}

	order := strings.ToUpper(strings.TrimSpace(request.Order))
	if order == "" {
		order = "ASC"
	}
	if order != "ASC" && order != "DESC" {
		return nil, fmt.Errorf("%w: invalid sort order", ErrInvalidInput)
	}

	// Convert level query: "2,3" -> []string{"2", "3"}
	var levels []string

	if strings.TrimSpace(request.Level) != "" {
		levelMap := make(map[string]bool)

		for _, levelID := range strings.Split(request.Level, ",") {
			levelID = strings.TrimSpace(levelID)

			if levelID == "" {
				return nil, fmt.Errorf("%w: invalid level filter", ErrInvalidInput)
			}

			if !levelMap[levelID] {
				levels = append(levels, levelID)
				levelMap[levelID] = true
			}
		}

		for _, levelID := range levels {
			_, err := u.levelRepo.GetByID(ctx, levelID)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("%w: level %s does not exist", ErrInvalidInput, levelID)
			}
			if err != nil {
				return nil, err
			}
		}
	}

	status := strings.TrimSpace(request.Status)
	if status != "" {
		_, err := u.statusRepo.GetByType(ctx, status)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: status does not exist", ErrInvalidInput)
		}
		if err != nil {
			return nil, err
		}
	}

	filter := repository.CourierFilter{
		Search:  strings.TrimSpace(request.Search),
		Level:   levels,
		Status:  status,
		Page:    page,
		PerPage: perPage,
		SortBy:  sortBy,
		Order:   order,
	}

	couriers, totalItems, err := u.courierRepo.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	data := make([]dto.CourierResponse, 0, len(couriers))

	for _, courier := range couriers {
		data = append(data, *toCourierResponse(&courier))
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))

	return &dto.CourierListResponse{
		Data: data,
		Pagination: dto.PaginationResponse{
			CurrentPage: page,
			PerPage:     perPage,
			TotalItems:  totalItems,
			TotalPages:  totalPages,
		},
	}, nil
}

// UPDATE COURIER
func (u *courierUsecase) Update(ctx context.Context, userID string, request dto.UpdateCourierRequest) (*dto.CourierResponse, error) {

	userID = strings.TrimSpace(userID)
	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.TrimSpace(request.Email)
	request.Phone = strings.TrimSpace(request.Phone)
	request.Address = strings.TrimSpace(request.Address)
	request.Level = strings.TrimSpace(request.Level)
	request.Status = strings.TrimSpace(request.Status)

	if userID == "" ||
		request.Name == "" ||
		request.Email == "" ||
		request.Phone == "" ||
		request.Address == "" ||
		request.Level == "" ||
		request.Status == "" {
		return nil, fmt.Errorf("%w: required fields must be filled", ErrInvalidInput)
	}

	if _, err := mail.ParseAddress(request.Email); err != nil {
		return nil, fmt.Errorf("%w: invalid email format", ErrInvalidInput)
	}

	if len(request.Phone) > 20 {
		return nil, fmt.Errorf("%w: phone maximum is 20 characters", ErrInvalidInput)
	}

	// Check existing Courier
	courier, err := u.courierRepo.GetByID(ctx, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	// Validate Level
	_, err = u.levelRepo.GetByID(ctx, request.Level)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: level does not exist", ErrInvalidInput)
	}
	if err != nil {
		return nil, err
	}

	// Validate Status
	_, err = u.statusRepo.GetByType(ctx, request.Status)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: status does not exist", ErrInvalidInput)
	}
	if err != nil {
		return nil, err
	}

	courier.Name = request.Name
	courier.Email = request.Email
	courier.Phone = request.Phone
	courier.Address = request.Address
	courier.Level = request.Level
	courier.Status = request.Status

	if err := u.courierRepo.Update(ctx, courier); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return toCourierResponse(courier), nil
}

// DELETE COURIER
func (u *courierUsecase) Delete(ctx context.Context, userID string) error {

	userID = strings.TrimSpace(userID)
	if userID == "" {
		return fmt.Errorf("%w: courier ID is required", ErrInvalidInput)
	}

	err := u.courierRepo.Delete(ctx, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}

	return err
}

// ENTITY TO RESPONSE DTO
func toCourierResponse(courier *entity.Courier) *dto.CourierResponse {
	return &dto.CourierResponse{
		UserID:    courier.UserID,
		Name:      courier.Name,
		Email:     courier.Email,
		Phone:     courier.Phone,
		Address:   courier.Address,
		Level:     courier.Level,
		Status:    courier.Status,
		CreatedAt: courier.CreatedAt,
		UpdatedAt: courier.UpdatedAt,
	}
}
