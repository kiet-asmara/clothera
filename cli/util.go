package cli

import (
	"bufio"
	"fmt"
	"os"
	"pair-project/pkg/validator"
	"strings"
)

// input untuk satu kata
func promptword(prompt string) (string, error) {
	fmt.Printf("%-10s: ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}

// input untuk satu lines
func promptline(prompt string) (string, error) {
	fmt.Printf("%-10s: ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}


/* ---------------------------------------------------------------- */
/*                            input auth                            */
/* ---------------------------------------------------------------- */


func inputUsername(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if ValidateUsername(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputEmail(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if ValidateEmail(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputPassword(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if ValidatePasswordPlaintext(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputCountry(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if ValidateCountry(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputCity(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if ValidateCity(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputStreet(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptline(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		input = strings.TrimSpace(input)

		if ValidateStreet(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}


/* ---------------------------------------------------------------- */
/*                           input update                           */
/* ---------------------------------------------------------------- */

func inputUpdateUsername(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if input == "-" {
			return input
		}

		if ValidateUsername(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputUpdateEmail(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if input == "-" {
			return input
		}

		if ValidateEmail(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputUpdatePassword(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if input == "-" {
			return input
		}

		if ValidatePasswordPlaintext(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputUpdateCountry(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if input == "-" {
			return input
		}

		if ValidateCountry(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputUpdateCity(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptword(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if input == "-" {
			return input
		}

		if ValidateCity(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

func inputUpdateStreet(v *validator.Validator, prompt string) string {
	for {
		v.Clear()
		input, err := promptline(prompt)
		if err != nil {
			fmt.Println("err:", err)
		}

		if input == "-" {
			return input
		}

		input = strings.TrimSpace(input)

		if ValidateStreet(v, input); !v.Valid() {
			fmt.Println(v.ShowError())
		} else {
			return input
		}
	}
}

