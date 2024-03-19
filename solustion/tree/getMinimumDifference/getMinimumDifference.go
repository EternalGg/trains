package getMinimumDifference

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func getMinimumDifference(root *TreeNode) int {
	list := preorderTraversal(root)
	list = mergeSort(list)
	min := list[1] - list[0]
	for i := 1; i < len(list); i++ {
		if list[i]-list[i-1] < min {
			min = list[i] - list[i-1]
		}
	}
	return min
}

type PreTree struct {
	list []int
}

func PreTreeInit() PreTree {
	p := new(PreTree)
	p.list = make([]int, 0)
	return *p
}

func preorderTraversal(root *TreeNode) []int {
	p := PreTreeInit()
	p.preorder(root)
	return p.list
}

func (t *PreTree) preorder(root *TreeNode) {
	if root != nil {
		t.list = append(t.list, root.Val)
		t.preorder(root.Left)
		t.preorder(root.Right)
	}
	return
}
