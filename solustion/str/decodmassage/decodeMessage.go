package decodmassage

func DecodeMessage(key string, message string) string {
	index, indexMap := uint8(97), make(map[uint8]string)

	for i := 0; i < len(key); i++ {
		if key[i] != 32 {
			if indexMap[key[i]] == "" {
				indexMap[key[i]] = string(index)
				index++
			}
		}
	}
	result := ""
	for i := 0; i < len(message); i++ {
		if message[i] != 32 {
			result += indexMap[message[i]]
		} else {
			result += " "
		}
	}
	return result
}
