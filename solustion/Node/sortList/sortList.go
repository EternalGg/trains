package sortList

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

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	list, valMap := make([]int, 0), make(map[int][]*ListNode)
	for head != nil {
		if valMap[head.Val] != nil {
			valMap[head.Val][len(valMap[head.Val])-1].Next = head
			valMap[head.Val] = append(valMap[head.Val], head)
			head = head.Next
			continue
		}
		list = append(list, head.Val)
		if len(valMap[head.Val]) == 0 {
			valMap[head.Val] = make([]*ListNode, 0)
			valMap[head.Val] = append(valMap[head.Val], head)
		}
		head = head.Next
	}
	list = mergeSort(list)
	for i := 0; i < len(list)-1; i++ {
		valMap[list[i]][len(valMap[list[i]])-1].Next = valMap[list[i+1]][0]
	}
	valMap[list[len(list)-1]][len(valMap[list[len(list)-1]])-1].Next = nil
	return valMap[list[0]][0]
}
