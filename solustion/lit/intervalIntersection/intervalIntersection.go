package intervalIntersection

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	M := make(map[int]bool)
	for _, nums := range firstList {
		head, cut := nums[0], nums[len(nums)-1]
		for i := head; i <= cut; i++ {
			M[i] = true
		}
	}
	result := make([][]int, 0)
	for _, nums := range secondList {
		fmt.Println(nums)
		head := -1
		header, cut := nums[0], nums[len(nums)-1]
		for i := header; i <= cut; i++ {
			if M[i] {
				if head == -1 {
					head = i
					if i == cut {
						now := make([]int, 2)
						now[0] = head
						now[1] = head
						result = append(result, now)
					}
				}

			} else {
				if head != -1 {
					now := make([]int, 2)
					now[0] = head
					now[1] = i - 1
					result = append(result, now)
					head = -1
				}
			}
		}

	}
	return result
}
