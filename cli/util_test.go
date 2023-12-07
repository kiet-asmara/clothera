package cli

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"pair-project/pkg/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateRandomString(length int) string {
	var container = "abcdefghijklmnopqrstuvwxyz123456789"

	var result string
	for i := 0; i < length; i++ {
		idx := rand.Intn(len(container))
		result += string(container[idx])
	}
	return result
}

func getSTDOUT() (*os.File, func(), error) {
	tmpfile, err := os.CreateTemp("", "testfile*")
	if err != nil {
		return nil, nil, err
	}

	oldStdout := os.Stdout
	os.Stdout = tmpfile

	return os.Stdin, func() {
		os.Stdout = oldStdout
		os.Remove(tmpfile.Name())
	}, nil
}

func getSTDIN() (*os.File, func(), error) {
	tmpfile, err := os.CreateTemp("", "testfile*")
	if err != nil {
		return nil, nil, err
	}

	oldStdin := os.Stdin
	os.Stdin = tmpfile

	return os.Stdin, func() {
		os.Stdin = oldStdin
		os.Remove(tmpfile.Name())
	}, nil
}

func truncateSTDIN(stdin *os.File) error {
	if err := stdin.Truncate(0); err != nil {
		return err
	}

	if _, err := stdin.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	return nil
}

func TestInputUsername(t *testing.T) {
	testcase := []struct {
		testname    string
		input       string
		expectError bool
		errorstring string
		expected    string
	}{
		{testname: "UsernameOK",
			input:       "examplename",
			expectError: false,
			errorstring: "",
			expected:    "examplename",
		},
		{testname: "FailUsernameEmpty",
			input:       "",
			expectError: true,
			errorstring: "must be provided",
			expected:    "",
		},
		{testname: "FailUsernameNotLengthEnough",
			input:       "a",
			expectError: true,
			errorstring: "must be at least 2 bytes long",
			expected:    "",
		},
		{testname: "FailUsernameContainSpecialCharacter",
			input:       "@#!@#",
			expectError: true,
			errorstring: "must not contain special character",
			expected:    "",
		},
	}

	// initialization
	v := validator.New()

	// mock stdin
	stdin, cleanup, err := getSTDIN()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// mock stdout to not cluttering output
	_, cleanfn, err := getSTDOUT()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanfn()

	// start test
	for _, tc := range testcase {
		if err := truncateSTDIN(stdin); err != nil {
			log.Fatal(err)
		}

		t.Run(tc.testname, func(t *testing.T) {
			if _, err := stdin.WriteString(tc.input); err != nil {
				log.Fatal(err)
			}

			if _, err := stdin.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			result := inputUsername(v, ">", true)
			fmt.Println()

			if tc.expectError {
				assert.Len(t, v.Errors, 1, "expected validator to have error")

				for _, err := range v.Errors {
					assert.Equal(t, tc.errorstring, err, "expected error to be: %v\nGot %v instead.", tc.errorstring, err)
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Expcted result: %v. Got %v instead\n", result, tc.expected)
			}

		})
	}
}

func TestInputEmail(t *testing.T) {
	testcase := []struct {
		testname    string
		input       string
		expectError bool
		errorstring string
		expected    string
	}{
		{testname: "EmailOK",
			input:       "example@example.com",
			expectError: false,
			errorstring: "",
			expected:    "example@example.com",
		},
		{testname: "FailEmailEmpty",
			input:       "",
			expectError: true,
			errorstring: "must be provided",
			expected:    "",
		},
		{testname: "FailEmailNotValidForm",
			input:       "a.com",
			expectError: true,
			errorstring: "must be a valid email address",
			expected:    "",
		},
	}

	// initialization
	v := validator.New()

	// mock stdin
	stdin, cleanup, err := getSTDIN()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// mock stdout to not cluttering output
	_, cleanfn, err := getSTDOUT()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanfn()

	// start test
	for _, tc := range testcase {
		if err := truncateSTDIN(stdin); err != nil {
			log.Fatal(err)
		}

		t.Run(tc.testname, func(t *testing.T) {
			if _, err := stdin.WriteString(tc.input); err != nil {
				log.Fatal(err)
			}

			if _, err := stdin.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			result := inputEmail(v, ">", true)
			fmt.Println()

			if tc.expectError {
				assert.Len(t, v.Errors, 1, "expected validator to have error")

				for _, err := range v.Errors {
					assert.Equal(t, tc.errorstring, err, "expected error to be: %v\nGot %v instead.", tc.errorstring, err)
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Expcted result: %v. Got %v instead\n", result, tc.expected)
			}

		})
	}
}

func TestInputPassword(t *testing.T) {
	testcase := []struct {
		testname    string
		input       string
		expectError bool
		errorstring string
		expected    string
	}{
		{testname: "PasswordOK",
			input:       "12345678",
			expectError: false,
			errorstring: "",
			expected:    "12345678",
		},
		{testname: "FailPasswordEmpty",
			input:       "",
			expectError: true,
			errorstring: "must be provided",
			expected:    "",
		},
		{testname: "FailPasswordNotLongEnough",
			input:       "12345",
			expectError: true,
			errorstring: "must be at least 8 bytes long",
			expected:    "",
		},
		{testname: "FailPasswordTooLong",
			input:       generateRandomString(80),
			expectError: true,
			errorstring: "must not be more than 72 bytes long",
			expected:    "",
		},
	}

	// initialization
	v := validator.New()

	// mock stdin
	stdin, cleanup, err := getSTDIN()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// mock stdout to not cluttering output
	_, cleanfn, err := getSTDOUT()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanfn()

	// start test
	for _, tc := range testcase {
		if err := truncateSTDIN(stdin); err != nil {
			log.Fatal(err)
		}

		t.Run(tc.testname, func(t *testing.T) {
			if _, err := stdin.WriteString(tc.input); err != nil {
				log.Fatal(err)
			}

			if _, err := stdin.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			result := inputPassword(v, ">", true)
			fmt.Println()

			if tc.expectError {
				assert.Len(t, v.Errors, 1, "expected validator to have error")

				for _, err := range v.Errors {
					assert.Equal(t, tc.errorstring, err, "expected error to be: %v\nGot %v instead.", tc.errorstring, err)
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Expcted result: %v. Got %v instead\n", result, tc.expected)
			}

		})
	}
}

func TestInputCountry(t *testing.T) {
	testcase := []struct {
		testname    string
		input       string
		expectError bool
		errorstring string
		expected    string
	}{
		{testname: "CountryOK",
			input:       "Nigeria",
			expectError: false,
			errorstring: "",
			expected:    "Nigeria",
		},
		{testname: "FailCountryEmpty",
			input:       "",
			expectError: true,
			errorstring: "must be provided",
			expected:    "",
		},
		{testname: "FailCountryTooShort",
			input:       generateRandomString(2),
			expectError: true,
			errorstring: "must be at least 4 bytes long",
			expected:    "",
		},
		{testname: "FailCountryHaveSpecialCharter",
			input:       "!@#!@#",
			expectError: true,
			errorstring: "must not contain special character or number",
			expected:    "",
		},
	}

	// initialization
	v := validator.New()

	// mock stdin
	stdin, cleanup, err := getSTDIN()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// mock stdout to not cluttering output
	_, cleanfn, err := getSTDOUT()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanfn()

	// start test
	for _, tc := range testcase {
		if err := truncateSTDIN(stdin); err != nil {
			log.Fatal(err)
		}

		t.Run(tc.testname, func(t *testing.T) {
			if _, err := stdin.WriteString(tc.input); err != nil {
				log.Fatal(err)
			}

			if _, err := stdin.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			result := inputCountry(v, ">", true)
			fmt.Println()

			if tc.expectError {
				assert.Len(t, v.Errors, 1, "expected validator to have error")

				for _, err := range v.Errors {
					assert.Equal(t, tc.errorstring, err, "expected error to be: %v\nGot %v instead.", tc.errorstring, err)
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Expcted result: %v. Got %v instead\n", result, tc.expected)
			}

		})
	}
}

func TestInputCity(t *testing.T) {
	testcase := []struct {
		testname    string
		input       string
		expectError bool
		errorstring string
		expected    string
	}{
		{testname: "CityOK",
			input:       "Semarang",
			expectError: false,
			errorstring: "",
			expected:    "Semarang",
		},
		{testname: "FailCityEmpty",
			input:       "",
			expectError: true,
			errorstring: "must be provided",
			expected:    "",
		},
		{testname: "FailCityyHaveSpecialCharter",
			input:       "!@#!@#",
			expectError: true,
			errorstring: "must not contain special character or number",
			expected:    "",
		},
	}

	// initialization
	v := validator.New()

	// mock stdin
	stdin, cleanup, err := getSTDIN()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// mock stdout to not cluttering output
	_, cleanfn, err := getSTDOUT()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanfn()

	// start test
	for _, tc := range testcase {
		if err := truncateSTDIN(stdin); err != nil {
			log.Fatal(err)
		}

		t.Run(tc.testname, func(t *testing.T) {
			if _, err := stdin.WriteString(tc.input); err != nil {
				log.Fatal(err)
			}

			if _, err := stdin.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			result := inputCity(v, ">", true)
			fmt.Println()

			if tc.expectError {
				assert.Len(t, v.Errors, 1, "expected validator to have error")

				for _, err := range v.Errors {
					assert.Equal(t, tc.errorstring, err, "expected error to be: %v\nGot %v instead.", tc.errorstring, err)
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Expcted result: %v. Got %v instead\n", result, tc.expected)
			}

		})
	}
}

func TestInputStreet(t *testing.T) {
	testcase := []struct {
		testname    string
		input       string
		expectError bool
		errorstring string
		expected    string
	}{
		{testname: "StreetOK",
			input:       "Semarang street",
			expectError: false,
			errorstring: "",
			expected:    "Semarang street",
		},
		{testname: "FailStreetEmpty",
			input:       "",
			expectError: true,
			errorstring: "must be provided",
			expected:    "",
		},
		{testname: "FailStreetHaveSpecialCharter",
			input:       "!@#!@#",
			expectError: true,
			errorstring: "must not contain special character",
			expected:    "",
		},
	}

	// initialization
	v := validator.New()

	// mock stdin
	stdin, cleanup, err := getSTDIN()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// mock stdout to not cluttering output
	_, cleanfn, err := getSTDOUT()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanfn()

	// start test
	for _, tc := range testcase {
		if err := truncateSTDIN(stdin); err != nil {
			log.Fatal(err)
		}

		t.Run(tc.testname, func(t *testing.T) {
			if _, err := stdin.WriteString(tc.input); err != nil {
				log.Fatal(err)
			}

			if _, err := stdin.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			result := inputStreet(v, ">", true)
			fmt.Println()

			if tc.expectError {
				assert.Len(t, v.Errors, 1, "expected validator to have error")

				for _, err := range v.Errors {
					assert.Equal(t, tc.errorstring, err, "expected error to be: %v\nGot %v instead.", tc.errorstring, err)
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Expcted result: %v. Got %v instead\n", result, tc.expected)
			}

		})
	}
}
