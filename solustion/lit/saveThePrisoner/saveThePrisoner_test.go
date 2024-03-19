package saveThePrisoner

import (
	"fmt"
	"testing"
)

func TestSaveThePrisoner(t *testing.T) {
	fmt.Println(SaveThePrisoner(352926151, 380324688, 94730870))
	fmt.Println(SaveThePrisoner(499999999, 999999998, 2))
	fmt.Println(SaveThePrisoner(999999999, 999999999, 1))
}
