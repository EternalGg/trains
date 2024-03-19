package totalMoney

import "fmt"

func TotalMoney(n int) int {
	result := 0
	week := n / 7
	cut := n % 7
	for i := 0; i < week; i++ {
		for j := 0; j < 7; j++ {
			fmt.Println(i + 1 + j)
			result += i + 1 + j
		}
	}
	for j := 0; j < cut; j++ {
		fmt.Println(week + 1 + j)
		result += week + 1 + j
	}
	return result
}
