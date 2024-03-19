package partition

type ListNode struct {
	Val  int
	Next *ListNode
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

func listToNode(list []int) *ListNode {
	result := new(ListNode)
	node := result
	for key, value := range list {
		node.Val = value
		if key == len(list)-1 {
			break
		}
		newOne := new(ListNode)
		node.Next = newOne
		node = newOne
	}
	return result
}

func partition(head *ListNode, x int) *ListNode {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	list = mergeSort(list)
	return listToNode(list)
}
