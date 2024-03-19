package maxDivScore

func maxDivScore(nums []int, divisors []int) int {
	uMap, max, key := make(map[int]bool), 0, divisors[0]
	for i := 0; i < len(divisors); i++ {
		if !uMap[divisors[i]] {
			uMap[divisors[i]] = true
			now := 0
			for j := 0; j < len(nums); j++ {
				if nums[j]%divisors[i] == 0 {
					now++
				}
			}
			if now > max {
				max = now
				key = divisors[i]
			}
			if now == max && divisors[i] < key {
				key = divisors[i]
			}
		}
	}
	return key
}
