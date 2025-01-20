package models

import (
	"time"

	"gorm.io/gorm"
)



type Book struct {
	ID 				uint 		`gorm:"primaryKey;autoIncrement" json:"-"`
	Title 			string 		`json:"title"`
	Author 			*string 	`json:"author"`
	CreatedAt    	time.Time    `json:"-"`  
  	UpdatedAt    	time.Time  	 `json:"-"`
}

func MigrateBook(db *gorm.DB) error {
	err:=db.AutoMigrate(&Book{})
	return err
}