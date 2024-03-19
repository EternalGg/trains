package distributeCandies

func DistributeCandies(candies int, num_people int) []int {
	result, round := make([]int, num_people), 0
	for candies != 0 {
		for i := 0; i < num_people; i++ {
			now := (round * num_people) + i + 1
			if candies >= now {
				candies -= now
				result[i] += now
			} else {
				result[i] += candies
				return result
			}
		}
		round++
	}
	return result
}
