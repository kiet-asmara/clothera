package testing

import (
	"pair-project/handler"
	"testing"
)

func TestDaysBetween(t *testing.T) {
	var tests = []struct {
		name       string
		inputStart string
		inputEnd   string
		want       int
	}{
		{"Days should be 10", "2023-11-25", "2023-12-05", 10},
		{"Days should be 10", "2023-11-20", "2023-11-30", 10},
		{"Days should be 1", "2023-11-25", "2023-11-25", 1}, // result = 1 because 0 days is rounded up to 1
	}

	for _, tt := range tests {
		res := handler.DaysBetween(tt.inputStart, tt.inputEnd)
		if res != tt.want {
			t.Errorf("got %d, want %d", res, tt.want)
		}

	}
}
