package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Id       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId   *uuid.UUID `gorm:"type:uuid;column:user_id" json:"userId"`
	User     User       `gorm:"foreignKey:UserId"`
	Name     string     `gorm:"column:name;" json:"name"`
	Location string     `gorm:"column:location;" json:"location"`
}

func (Folder) TableName() string {
	return "folders"
}
