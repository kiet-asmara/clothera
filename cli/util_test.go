package cli

import (
	"pair-project/pkg/validator"
	"testing"
)

func TestInputUsername(t *testing.T) {
	testcase := []struct {
		prompt    string
		validator *validator.Validator
	}{}

	v := validator.New()
	result := inputUsername(v, ">")
}
