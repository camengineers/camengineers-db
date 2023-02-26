package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	Id        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId    *uuid.UUID `gorm:"type:uuid;column:user_id" json:"userId"`
	User      User       `gorm:"foreignKey:UserId"`
	LoggedIn  time.Time  `gorm:"column:logged_in;" json:"loggedIn"`
	UserAgent string     `gorm:"column:user_agent;" json:"userAgent"`
	IpAddress string     `gorm:"column:ip_address;" json:"ipAddress"`
	Platform  int16      `gorm:"column:platform;default:1" json:"platform"`
	Browser   string     `gorm:"column:browser;" json:"browser"`
	Version   string     `gorm:"column:version;" json:"version"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}

func (u UserSession) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":       u.Id,
		"loggedIn": u.LoggedIn,
	}
	return payload
}
