package largestValues

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

func largestValues(root *TreeNode) []int {
	list := levelOrder(root)
	result := make([]int, len(list))
	for i, ints := range list {
		now := mergeSort(ints)
		result[i] = now[len(now)-1]
	}
	return result
}

type LevelTree struct {
	level int
	max   int
	lMap  map[int][]int
}

func LevelTreeInit() LevelTree {
	l := new(LevelTree)
	l.level = 0
	l.max = 0
	l.lMap = make(map[int][]int)
	return *l
}

func levelOrder(root *TreeNode) [][]int {
	l := LevelTreeInit()

	result := make([][]int, 0)
	l.lorder(root)
	for i := 0; i < l.max; i++ {
		result = append(result, l.lMap[i])
	}

	return result
}

func (l *LevelTree) lorder(root *TreeNode) {
	if root != nil {
		l.lMap[l.level] = append(l.lMap[l.level], root.Val)
		l.level++
		if l.level > l.max {
			l.max = l.level
		}
		l.lorder(root.Left)
		l.lorder(root.Right)
		l.level--
	}
	return
}
