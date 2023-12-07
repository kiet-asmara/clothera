package cli

import (
	"pair-project/pkg/validator"
	"strconv"
	"strings"
)

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
	v.Check(!strings.ContainsAny(password, " \n\t\r"), "password", "must not contain white space")
}

func ValidateUsername(v *validator.Validator, username string) {
	v.Check(username != "", "username", "must be provided")
	v.Check(len(username) >= 2, "username", "must be at least 2 bytes long")
	v.Check(!strings.ContainsAny(username, " \n\t\r!@#$%^&*()_+-=?><';:{}[]|"), "username", "must not contain special character")
}

func ValidateCountry(v *validator.Validator, country string) {
	v.Check(country != "", "country", "must be provided")
	v.Check(len(country) >= 4, "country", "must be at least 4 bytes long")
	v.Check(!strings.ContainsAny(country, " \n\t\r!@#$%^&*()_+-=?><';:{}[]|123456789"), "country", "must not contain special character or number")
}

func ValidateCity(v *validator.Validator, city string) {
	v.Check(city != "", "city", "must be provided")
	v.Check(!strings.ContainsAny(city, " \n\t\r!@#$%^&*()_+-=?><';:{}[]|123456789"), "city", "must not contain special character or number")
}

func ValidateStreet(v *validator.Validator, street string) {
	v.Check(street != "", "street", "must be provided")
	v.Check(!strings.ContainsAny(street, "\n\t\r!@#$%^&*()_+-=?><';:{}[]|"), "street", "must not contain special character")
}

func ValidateStringNoSpecialChar(v *validator.Validator, input string) {
	v.Check(input != "", "value", "must be provided")
	v.Check(!strings.ContainsAny(input, "\n\t\r!@#$%^&*()_+-=?><';:{}[]|123456789"), "vlaue", "must not contain special character or number")
}

func ValidateNumber(v *validator.Validator, input string) {
	v.Check(input != "", "value", "must be provided")

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		v.AddError("value", "must be valid number")
	}

	v.Check(num > 0, "value", "must be positive")
}
