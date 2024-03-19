package threeConsecutiveOdds

func ThreeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	left, right, count := 0, 0, 0
	for right = 0; right < 3; right++ {
		if arr[right]%2 != 0 {
			count++
		}
	}
	for right != len(arr)+1 {
		if count == 3 {
			return true
		}
		if arr[left]%2 != 0 {
			count--
		}
		left++
		if arr[right]%2 != 0 {
			count++
		}
		right++
	}
	return false
}
