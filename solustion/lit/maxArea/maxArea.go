package maxArea

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			now := FindMin(height[i], height[j]) * (j - i)
			if now > max {
				max = now
			}
		}
	}
	return max
}

func MaxArea1(height []int) int {
	left, right, max := 0, len(height)-1, 0
	for right > left {
		if (FindMin(height[left], height[right]) * (right - left)) > max {
			max = FindMin(height[left], height[right]) * (right - left)
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func FindMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
