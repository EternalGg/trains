package checkIfPangram

func CheckIfPangram(sentence string) bool {
	count, nMap := 26, make(map[int32]bool)

	for _, i2 := range sentence {

		if nMap[i2] == false {
			count--
			nMap[i2] = true
		}
		if count == 0 {
			return true
		}
	}
	return false
}
