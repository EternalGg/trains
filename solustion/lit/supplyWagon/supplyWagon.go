package supplyWagon

func SupplyWagon(supplies []int) []int {
	target := len(supplies)
	for len(supplies) != target/2 {
		min, key := supplies[1]+supplies[0], 0
		for i := 1; i < len(supplies); i++ {
			if supplies[i]+supplies[i-1] < min {
				min = supplies[i] + supplies[i-1]
				key = i - 1
			}
		}
		supplies[key] = min
		head := supplies[:key+1]
		last := supplies[key+2:]
		supplies = append(head, last...)
	}
	return supplies
}
