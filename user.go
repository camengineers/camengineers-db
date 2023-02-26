package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name       string         `gorm:"column:name" json:"name"`
	Phone      string         `gorm:"column:phone;unique" json:"phone"`
	Email      string         `gorm:"column:email;unique" json:"email"`
	Role       int16          `gorm:"column:role;default:3" json:"role"`
	Active     bool           `gorm:"column:active;default:true" json:"active"`
	ProfilePic sql.NullString `gorm:"column:profile_pic" json:"profilePic"`
	Gender     sql.NullInt16  `gorm:"column:gender" json:"gender"`
	BirthDate  sql.NullTime   `gorm:"column:birth_date" json:"birthDate"`
}

func (User) TableName() string {
	return "users"
}

func (u User) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":     u.Id,
		"name":   u.Name,
		"phone":  u.Phone,
		"email":  u.Email,
		"role":   u.Role,
		"active": u.Active,
	}
	return payload
}

func (u User) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"name":       u.Name,
		"email":      u.Email,
		"phone":      u.Phone,
		"role":       u.Role,
		"active":     u.Active,
		"profilePic": nil,
		"gender":     nil,
		"birthDate":  nil,
		"createdAt":  u.CreatedAt.Format(time.RFC3339),
	}
	if u.ProfilePic.Valid {
		payload["profilePic"] = u.ProfilePic.String
	}
	if u.Gender.Valid {
		payload["gender"] = u.Gender.Int16
	}
	if u.BirthDate.Valid {
		payload["birthDate"] = u.BirthDate.Time.Format(time.RFC3339)
	}
	return payload
}
