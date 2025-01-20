package storage

import (
	"authentication/models"
	"log"

	"gorm.io/gorm"
)


func MigrateDB(db *gorm.DB){
	err :=models.MigrateBook(DB)
	if err != nil {
		log.Fatal("Unable to migrate database")
	}
	err =models.MigrateUser(DB)
	if err != nil {
		log.Fatal("Unable to migrate database")
	}
	err =models.MigrateCredential(DB)
	if err != nil {
		log.Fatal("Unable to migrate database")
	}
}