package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) ([]byte, error) {
	str := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(str, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashed, nil
}

func CompareHashAndPassword(hashed []byte, password string) error {
	inputPassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashed, inputPassword)
	if err != nil {
		return err
	}
	return nil
}
