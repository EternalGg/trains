package splitNum

import "strconv"

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

func SplitNum(num int) int {
	bList := []byte(strconv.Itoa(num))
	nlist := make([]int, 0)
	for i := 0; i < len(bList); i++ {
		if bList[i] != '0' {
			nlist = append(nlist, int(bList[i]-48))
		}
	}
	nlist = mergeSort(nlist)
	n1, n2 := "", ""
	for i := 0; i < len(nlist)/2; i++ {
		n1 += strconv.Itoa(nlist[i*2])
		n2 += strconv.Itoa(nlist[(i*2)+1])
	}
	if len(nlist)%2 != 0 {
		n1 += strconv.Itoa(nlist[len(nlist)-1])
	}
	h, l := string(n1), string(n2)
	hn, _ := strconv.ParseInt(h, 10, 64)
	ln, _ := strconv.ParseInt(l, 10, 64)

	return int(hn) + int(ln)
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
