package commons

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	newHash := md5.Sum([]byte(password))
	hashedPassword := hex.EncodeToString(newHash[:])
	if hashedPassword == hash {
		return true
	}
	return false
}
