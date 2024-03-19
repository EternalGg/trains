package createTargetArray

func CreateTargetArray(nums []int, index []int) []int {
	target := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		key, value := index[i], nums[i]
		switch key {
		case 0:
			this := []int{value}
			target = append(this, target...)
		default:
			head, last := target[:key], target[key:]
			this := []int{value}
			last = append(this, last...)
			head = append(head, last...)
			target = head
		}
	}
	return target
}
