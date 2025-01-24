package models

import (
	"time"

	"gorm.io/gorm"
)

type Credential struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	Username     string    `json:"username"`
	Password     string    `json:"password,omitempty"`
	UserID       uint      `json:"-"`
	RefreshToken string    `json:"refreshToken,omitempty"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

func MigrateCredential(db *gorm.DB) error {
	err := db.AutoMigrate(&Credential{})
	return err
}
