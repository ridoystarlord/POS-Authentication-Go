package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID 				uint 		`gorm:"primaryKey;autoIncrement" json:"-"`
	Name 			string 		`json:"name"`
	Credential 		Credential	`gorm:"foreignKey:UserID"`
	CreatedAt    	time.Time    `json:"-"`  
  	UpdatedAt    	time.Time  	 `json:"-"`
}

func MigrateUser(db *gorm.DB) error {
	err:=db.AutoMigrate(&User{})
	return err
}