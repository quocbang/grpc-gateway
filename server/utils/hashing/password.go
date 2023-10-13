package hashing

import (
	"github.com/quocbang/grpc-gateway/server/errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func IsMatchedPassword(requestPassword string, storedPassword []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(storedPassword, []byte(requestPassword)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, errors.Error{
				Code:    errors.Code_WRONG_PASSWORD,
				Details: "wrong password",
			}
		}
		return false, err
	}
	return true, nil
}
