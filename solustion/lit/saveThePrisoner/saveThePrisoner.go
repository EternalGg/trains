package saveThePrisoner

func SaveThePrisoner(n int32, m int32, s int32) int32 {
	// Write your code here
	result := (m + s - 1) % n
	// if result is last seat, it will produce zero due to modulus, so return the last seat, aka N
	if result == 0 {
		result = n
	}
	return result

}
