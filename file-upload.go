package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileUpload struct {
	gorm.Model
	Id       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	FilePath string     `gorm:"column:file_path" json:"filePath"`
	FileType int64      `gorm:"column:file_type" json:"fileType"`
	UserId   *uuid.UUID `gorm:"type:uuid;column:user_id" json:"userId"`
	User     User       `gorm:"foreignKey:UserId"`
}

func (FileUpload) TableName() string {
	return "file_uploads"
}

func (o FileUpload) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        o.Id,
		"filePath":  o.FilePath,
		"fileType":  o.FileType,
		"createdAt": o.CreatedAt.Format(time.RFC3339),
	}
	return payload
}
