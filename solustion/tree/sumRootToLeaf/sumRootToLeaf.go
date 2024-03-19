package sumRootToLeaf

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeList struct {
	BList []byte
	Count int
}

func TreeListInit() *TreeList {
	t := new(TreeList)
	t.BList = make([]byte, 0)
	t.Count = 0
	return t
}

func sumRootToLeaf(root *TreeNode) int {
	t := TreeListInit()
	t.CountAllLeaf(root)
	return t.Count
}

func (t *TreeList) Calculate() {
	bToInt, _ := strconv.ParseInt(string(t.BList), 2, 64)
	t.BList = t.BList[:len(t.BList)-1]
	t.Count += int(bToInt)
}

func (t *TreeList) CountAllLeaf(root *TreeNode) {
	if root != nil {
		t.BList = append(t.BList, uint8(root.Val+48))
		if root.Left == nil && root.Right == nil {
			t.Calculate()
			return
		}
		t.CountAllLeaf(root.Left)
		t.CountAllLeaf(root.Right)
		t.BList = t.BList[:len(t.BList)-1]
	}
	return
}

func IntegerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
