package dto

// CreateCourierRequest digunakan untuk menambahkan Courier.
type CreateCourierRequest struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Level   string `json:"level"`
	Status  string `json:"status"`
}

// UpdateCourierRequest digunakan untuk memperbarui data Courier.
type UpdateCourierRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Level   string `json:"level"`
	Status  string `json:"status"`
}

// CourierQueryRequest digunakan untuk filter daftar Courier.
type CourierQueryRequest struct {
	Search  string `query:"search"`
	Level   string `query:"level"`
	Status  string `query:"status"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
	Sort    string `query:"sort"`
	Order   string `query:"order"`
}
