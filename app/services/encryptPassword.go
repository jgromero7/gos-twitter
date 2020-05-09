package services

import "golang.org/x/crypto/bcrypt"

// EncryptPassword generates an encrypted string
func EncryptPassword(password string) (string, error) {
	rounds := 10

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
