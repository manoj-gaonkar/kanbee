package models

type Project struct {
	Name        string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Tasks       []Task `gorm:"foreignKey:ProjectID"`
	ID          uint   `gorm:"primaryKey"`
}
