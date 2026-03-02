package handlers

import "github.com/matthewhartstonge/argon2"

var argon = argon2.DefaultConfig()

func HashPassword(password string) (string, error) {
	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(encodedHash, password string) bool {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(encodedHash))
	if err != nil {
		return false
	}
	return ok
}
