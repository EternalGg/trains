package KthLargest

import "sort"

type KthLargest struct {
	kList []int
	k     int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	kl := KthLargest{k: k}
	if len(nums) >= k {
		kl.kList = nums[:k]
	} else {
		kl.kList = nums
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if len(this.kList) < this.k {
		return 0
	}
	return 0
}
