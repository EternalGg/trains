package extraLongFactorials

import (
	"fmt"
)

func ExtraLongFactorials(n int32) {
	// Write your code here
	fmt.Println(factorial(int(n)))
}

func factorial(n int) uint64 {
	facVal := uint64(1)
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			facVal *= uint64(i)
		}
	}
	return facVal

}
