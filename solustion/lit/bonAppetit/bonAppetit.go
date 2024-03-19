package bonAppetit

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	count := int32(0)

	for i := 0; i < len(bill); i++ {
		if int32(i) != k {
			count += bill[i]
		}
	}
	if (count / 2) == b {
		fmt.Println("Bon Appetit")
	}
	if (count / 2) < b {
		fmt.Println((count / 2) - b)
	}
}
