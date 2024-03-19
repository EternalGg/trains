package firstBadVersion

import "fmt"

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	openPoint := n / 2
	if isBadVersion(openPoint) {
		openPoint = 0
	}
	if openPoint == 0 {
		fmt.Println(openPoint)
		for i := 0; i < n; i++ {
			if isBadVersion(i) == true {
				return i
			}
		}
	} else {
		fmt.Println(openPoint)
		for i := n; i >= openPoint; i-- {
			if isBadVersion(i) == false {
				return i - 1
			}
		}
	}
	return 0
}
