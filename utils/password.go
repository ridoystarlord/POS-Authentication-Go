package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
 hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
 return string(hash)
}

func ComparePassword(hashedPassword, password string) bool {
 err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
 return err == nil
}