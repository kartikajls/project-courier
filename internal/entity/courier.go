package entity

import "time"

type Courier struct {
	UserID    string    `json:"user_id" gorm:"column:user_id;primaryKey"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"column:email"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	Address   string    `json:"address" gorm:"column:address"`
	Level     string    `json:"level" gorm:"column:level"`
	Status    string    `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (Courier) TableName() string {
	return "couriers"
}
