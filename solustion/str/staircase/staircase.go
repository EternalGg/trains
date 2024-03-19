package staircase

import "fmt"

func Staircase(n int32) {
	// Write your code here
	sharp := ""

	for i := n - 1; i >= 0; i-- {
		now := ""

		for j := 0; j < int(i); j++ {
			now += " "
		}
		sharp = fmt.Sprint(sharp, "#")
		fmt.Println(now + sharp)
	}
}
