package tool

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func main() {

	hash, _ := HashPassword("test")
	fmt.Println(hash)

	fmt.Println(CheckPasswordHash("test", hash))
}
