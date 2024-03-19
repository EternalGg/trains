package isValidSudoku

import (
	"fmt"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	m := [][]byte{}
	m = append(m, []byte{'.', '.', '.', '.', '5', '.', '.', '1', '.'})
	m = append(m, []byte{'.', '4', '.', '3', '.', '.', '.', '.', '.'})
	m = append(m, []byte{'.', '.', '.', '.', '.', '3', '.', '.', '1'})
	m = append(m, []byte{'8', '.', '.', '.', '.', '.', '.', '2', '.'})
	m = append(m, []byte{'.', '.', '2', '.', '7', '.', '.', '.', '.'})
	m = append(m, []byte{'.', '1', '5', '.', '.', '.', '.', '.', '.'})
	m = append(m, []byte{'.', '.', '.', '.', '.', '2', '.', '.', '.'})
	m = append(m, []byte{'.', '2', '.', '9', '.', '.', '.', '.', '.'})
	m = append(m, []byte{'.', '.', '4', '.', '.', '.', '.', '.', '.'})
	fmt.Println(IsValidSudoku(m))
}
