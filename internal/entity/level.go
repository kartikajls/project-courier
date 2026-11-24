package entity

type Level struct {
	LevelID     string `json:"level_id" gorm:"column:level_id;primaryKey"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (Level) TableName() string {
	return "levels"
}
