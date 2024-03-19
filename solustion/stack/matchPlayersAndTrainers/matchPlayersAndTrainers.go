package matchPlayersAndTrainers

type MyStack struct {
	list   []int
	length int
}

func Constructor() MyStack {
	s := new(MyStack)
	s.list = make([]int, 0)
	s.length = 0
	return *s
}

func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)
	this.length++
}

func (this *MyStack) Pop() int {
	if this.length == 0 {
		return -1
	} else {
		result := this.list[this.length-1]
		this.list = this.list[:this.length-1]
		this.length--
		return result
	}
}

func (this *MyStack) Top() int {
	if this.length == 0 {
		return -1
	} else {
		return this.list[this.length-1]
	}
}

func (this *MyStack) Empty() bool {
	if this.length == 0 {
		return true
	} else {
		return false
	}
}

func MatchPlayersAndTrainers(players []int, trainers []int) int {
	pStack, tStack := Constructor(), Constructor()
	players = mergeSort(players)
	trainers = mergeSort(trainers)
	for _, player := range players {
		pStack.Push(player)
	}
	for _, value := range trainers {
		tStack.Push(value)
	}
	result := 0
	for !pStack.Empty() {
		if pStack.Top() <= tStack.Top() {
			pStack.Pop()
			tStack.Pop()
			result++
		} else {
			pStack.Pop()
		}
	}
	return result
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
