package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppFile struct {
	gorm.Model
	Id       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	FolderId *uuid.UUID `gorm:"type:uuid;column:folder_id" json:"folderId"`
	Folder   Folder     `gorm:"foreignKey:FolderId"`
	Name     string     `gorm:"column:name;" json:"name"`
}

func (AppFile) TableName() string {
	return "app_files"
}
