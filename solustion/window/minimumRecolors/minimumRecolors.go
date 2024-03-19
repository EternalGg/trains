package minimumRecolors

func MinimumRecolors(blocks string, k int) int {
	left, right, now, result := 0, 0, 0, 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			now++
		}
	}
	result = now
	right = k - 1
	for right != len(blocks)-1 {
		head := blocks[left]
		if head == 'W' {
			now--
		}
		left++
		right++
		last := blocks[right]
		if last == 'W' {
			now++
		}
		if result > now {
			result = now
		}
	}
	return result
}
