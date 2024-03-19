package findDigits

import "testing"

func FuzzFindDigits(f *testing.F) {
	FindDigits(12)
	FindDigits(1012)
}
