package checkIfExist

func checkIfExist(arr []int) bool {

	nMap := make(map[float64]bool)
	for _, num := range arr {
		if num == 0 {
			if nMap[float64(num)] {
				return true
			}
			nMap[float64(num)] = true
			continue
		}
		nMap[float64(num)] = true
		half, double := float64(num)/2, float64(num)*2
		if nMap[half] || nMap[double] {
			return true
		}
	}
	return false
}
