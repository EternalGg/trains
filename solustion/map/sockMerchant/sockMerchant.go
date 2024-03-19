package sockMerchant

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	smap := make(map[int32]int)
	for i := 0; i < int(n); i++ {
		smap[ar[i]]++
	}
	result := 0
	for _, i2 := range smap {
		result += i2 / 2
	}
	return int32(result)
}
