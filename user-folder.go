package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFolder struct {
	gorm.Model
	Id       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId   *uuid.UUID `gorm:"type:uuid;column:user_id" json:"userId"`
	User     User       `gorm:"foreignKey:UserId"`
	FolderId *uuid.UUID `gorm:"type:uuid;column:folder_id" json:"folderId"`
	Folder   Folder     `gorm:"foreignKey:FolderId"`
}

func (UserFolder) TableName() string {
	return "user_folders"
}
