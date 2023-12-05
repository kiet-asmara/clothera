package handler

import "golang.org/x/crypto/bcrypt"

func hashPassword(plainPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
}

func comparePassword(src []byte, target string) error {
	err := bcrypt.CompareHashAndPassword(src, []byte(target))
	if err != nil {
		return err
	}
	return nil
}
