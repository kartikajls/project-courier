package dto

import "time"

// CourierResponse digunakan untuk menampilkan detail Courier.
type CourierResponse struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Level     string    `json:"level"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CourierListResponse digunakan untuk menampilkan daftar Courier.
type CourierListResponse struct {
	Data       []CourierResponse  `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

// PaginationResponse berisi informasi pagination.
type PaginationResponse struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
}
