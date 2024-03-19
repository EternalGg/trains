package subsets

import (
	"fmt"
	"testing"
)

//func subsets(nums []int) [][]int {
//	qa := [][]int{}
//	qa = append(qa, []int{})
//	qa = append(qa, []int{})
//
//	for key, value := range nums {
//		qa[1] = append(qa[1], value)
//		//i等于一个数组有多少个数和当前所加 j
//		for i := 0; i < key; i++ {
//			//
//			now := make([]int, i)
//			for j := 0; j < key-1; j++ {
//
//			}
//		}
//	}
//}

func TestSubsets(t *testing.T) {
	list := []int{9, 0, 3, 5, 7}
	fmt.Println(Subsets(list))
}
