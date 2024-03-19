package consecutiveNumbersSum

import "fmt"

func ConsecutiveNumbersSum(n int) int {
	result := 1

	half := n / 2
	// single 1 2 3 4 5 (15) 15/2==2*7+1
	// double 1 2 3 4 (10/2)=5 -1
	for i := 1; i < n; i++ {

		if half-i <= n && (half-i) > 0 && (half-i)%2 == 0 {
			fmt.Println(half - i)
			result++
		} else if i*(i+1)/2 == n {
			fmt.Println(i)
			result++
		}

	}
	return result
}
