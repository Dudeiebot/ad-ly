package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id              string `gorm:"primaryKey"`
	Name            string
	Password        string
	Email           string
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Links           []Link `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Link struct {
	Code      string `gorm:"primaryKey"`
	UserId    string
	Url       string
	ExpireAt  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Empty() bool {
	return u.Id == ""
}

func (u *User) EmailVerified() bool {
	if u.EmailVerifiedAt != nil {
		if u.EmailVerifiedAt.IsZero() {
			return false
		}
		return true
	}
	return false
}

func (l *Link) Expired() bool {
	if l.ExpireAt == nil || l.ExpireAt.IsZero() {
		return false
	}
	return time.Now().After(*l.ExpireAt)
}
