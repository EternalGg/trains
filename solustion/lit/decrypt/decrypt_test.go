package decrypt

import (
	"fmt"
	"testing"
)

func TestDecrypt(t *testing.T) {
	fmt.Println(Decrypt([]int{2, 4, 9, 3}, -2))
}
