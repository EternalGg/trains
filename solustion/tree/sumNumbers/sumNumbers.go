package sumNumbers

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Count struct {
	list   []int
	result [][]int
}

func sumNumbers(root *TreeNode) int {
	c := new(Count)
	c.list = make([]int, 0)
	c.result = make([][]int, 0)
	c.cTree(root)

	result := 0
	for _, ints := range c.result {
		str := ""
		for _, num := range ints {
			str += strconv.Itoa(num)
		}
		now, _ := strconv.ParseInt(str, 10, 64)
		result += int(now)
	}
	return result
}

func (t *Count) cTree(root *TreeNode) {
	if root != nil {
		t.list = append(t.list, root.Val)
		if root.Left == nil && root.Right == nil {
			newList := make([]int, len(t.list))
			for i, num := range t.list {
				newList[i] = num
			}
			t.result = append(t.result, newList)
		}
		t.cTree(root.Left)
		t.cTree(root.Right)
		t.list = t.list[:len(t.list)-1]
	}
	return
}
