package entity

type Status struct {
	Type string `json:"type" gorm:"column:type;primaryKey"`
}

func (Status) TableName() string {
	return "status"
}
