package getTriggerTime

import "fmt"

func GetTriggerTime(increase [][]int, requirements [][]int) []int {

	iLength := len(increase)
	if iLength == 1 && requirements[0][0] == 0 && requirements[0][1] == 0 && requirements[0][2] == 0 {
		return []int{0}
	}
	cList, rList, hList := make([]int, iLength), make([]int, iLength), make([]int, iLength)
	cList[0], rList[0], hList[0] = increase[0][0], increase[0][1], increase[0][2]
	for i := 1; i < iLength; i++ {
		cList[i] += increase[i][0] + cList[i-1]
		rList[i] += increase[i][1] + rList[i-1]
		hList[i] += increase[i][2] + hList[i-1]
	}

	max := []int{cList[iLength-1], rList[iLength-1], hList[iLength-1]}
	result := make([]int, len(requirements))
	for i := 0; i < len(requirements); i++ {
		now := requirements[i]
		if now[0] > max[0] || now[1] > max[1] || now[2] > max[2] {
			result[i] = -1
		} else {
			begain := thirdMin(BinaryFindSmall(cList, now[0]), BinaryFindSmall(rList, now[1]), BinaryFindSmall(hList, now[2]))
			fmt.Println(begain)
			for j := begain; j < iLength; j++ {
				if now[0] <= cList[j] && now[1] <= rList[j] && now[2] <= hList[j] {
					result[i] = j
					break
				}
			}
		}
	}
	return result
}

func BinaryFindSmall(nums []int, target int) int {
	result := BinaryFind(nums, 0, len(nums)-1, target)
	return result
}

func BinaryFind(nums []int, head, last, target int) int {
	half := ((last - head) / 2) + head
	if head > last {
		return half - 1
	}
	switch {
	case nums[head] == target:
		return head
	case nums[last] == target:
		return last
	case nums[half] == target:
		return half
		// 当前 > 目标
	case nums[half] > target:
		return BinaryFind(nums, head, half-1, target)
		// 当前 < 目标
	case nums[half] < target:
		return BinaryFind(nums, half+1, last, target)
	}
	return half
}

func thirdMin(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
