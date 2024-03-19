package minDeletionSize

import "fmt"

func MinDeletionSize(strs []string) int {
	rows := make([][]int32, len(strs[0]))
	for _, str := range strs {
		for i2, i32 := range str {
			rows[i2] = append(rows[i2], i32)
		}
	}
	result := 0
	for _, row := range rows {
		last := int32(0)
		for _, list := range row {
			if list >= last {
				last = list
			} else {
				fmt.Println(row)
				result++
				break
			}
		}
	}
	return result
}
