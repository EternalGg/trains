package combinationSum

import (
	"fmt"
	"sort"
)

func CombinationSum(candidates []int, target int) [][]int {
	resultMap := make(map[int][][]int, 0)
	cl, result := make([][]int, 0), make([][]int, 0)
	sort.Ints(candidates)
	// 查找比数字大的 切割数组
	for i := 0; i < len(candidates); i++ {
		if candidates[i] >= target {
			if candidates[i] == target {
				result = append(result, []int{target})
			}
			candidates = candidates[:i]
			break
		}
	}

	for _, candidate := range candidates {
		cl = append(cl, []int{candidate})
	}
	for len(cl) != 0 {
		l := make([]int, len(cl[0]))
		copy(l, cl[0])
		ttt := []int{3, 3, 3, 3, 3}
		if len(l) == 5 {
			if listequal(cl[0], ttt) {
				fmt.Println("zz")
			}
		}
		for i := 0; i < len(candidates); i++ {
			tl := make([]int, len(l))
			copy(tl, l)
			tl = append(tl, candidates[i])
			sort.Ints(tl)
			c := ListCount(tl)
			if c > target {
				break
			} else if c == target {
				//fmt.Println(c)
				if !ListCompare(resultMap[len(tl)], tl) {
					if c == target {
						resultMap[len(tl)] = append(resultMap[len(tl)], tl)
					}
				}
			} else {
				cl = append(cl, tl)
			}
		}

		cl = cl[1:]
		fmt.Println(cl)
	}

	for _, i2 := range resultMap {
		result = append(result, i2...)
	}
	return result
}

func ListCompare(matrix [][]int, list []int) bool {
	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(list); i++ {
			if listequal(list, matrix[j]) {
				goto here
			}
		}
		return true
	here:
	}
	return false
}

func ListCount(list []int) (count int) {
	for _, v := range list {
		count += v
	}
	return
}
func listequal(l1, l2 []int) bool {
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
