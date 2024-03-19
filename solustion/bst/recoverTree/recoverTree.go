package recoverTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PreTree struct {
	list []*TreeNode
}

func PreTreeInit() PreTree {
	p := new(PreTree)
	p.list = make([]*TreeNode, 0)
	return *p
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	p := PreTreeInit()
	p.preorder(root)
	return p.list
}

func (t *PreTree) preorder(root *TreeNode) {
	if root != nil {
		t.preorder(root.Left)
		t.list = append(t.list, root)
		t.preorder(root.Right)
	}
	return
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		// 分治，两两拆分，一直拆到基础元素才向上递归。
		return nums
	}
	i := len(nums) / 2
	left := mergeSort(nums[0:i])
	// 左侧数据递归拆分
	right := mergeSort(nums[i:])
	// 右侧数据递归拆分
	result := merge(left, right)
	// 排序 & 合并
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

func recoverTree(root *TreeNode) {
	list, c1, c2 := preorderTraversal(root), -1, -1
	mergeList := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		mergeList[i] = list[i].Val
	}
	mergeList = mergeSort(mergeList)
	for i := 0; i < len(list); i++ {
		if list[i].Val != mergeList[i] {
			if c1 == -1 {
				c1 = i
			} else {
				c2 = i
			}
		}
	}

	list[c1].Val = mergeList[c2]
	list[c2].Val = mergeList[c1]
}
