package testing

import (
	"os"
	"pair-project/handler"
	"testing"
)

func TestGetSelectedCategoryFromUserValidCategory(t *testing.T) {
	input := "Kemeja\n"

	simulateUserInput(input, func() {
		categories := []string{"T-Shirt", "Hoodie", "Jaket"}
		selectedCategory := handler.GetSelectedCategoryFromUser(categories)

		expectedCategory := "Kemeja"
		if selectedCategory != expectedCategory {
			t.Errorf("Expected selected category to be %s, got %s", expectedCategory, selectedCategory)
		}
	})
}

func simulateUserInput(input string, f func()) {
	originalStdin := os.Stdin

	r, w, _ := os.Pipe()
	w.Write([]byte(input))
	w.Close()
	os.Stdin = r

	f()

	os.Stdin = originalStdin
}
