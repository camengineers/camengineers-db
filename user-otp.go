package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserOtp struct {
	gorm.Model
	Id           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId       *uuid.UUID `gorm:"type:uuid;column:user_id" json:"userId"`
	User         User       `gorm:"foreignKey:UserId"`
	Code         string     `gorm:"column:code;" json:"code"`
	IsUsed       bool       `gorm:"column:is_used;default:false" json:"isUsed"`
	CheckedCount int16      `gorm:"column:checked_count;default:0" json:"checkedCount"`
	OtpType      int16      `gorm:"column:otp_type;default:1" json:"otpType"`
}

func (UserOtp) TableName() string {
	return "user_otps"
}
