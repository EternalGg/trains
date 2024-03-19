package decodeMessage

func DecodeMessage(key string, message string) string {
	newIndex, nMap, qa := make([]int32, 0), make(map[int32]bool), ""

	for _, word := range key {
		if word >= 97 && word <= 122 {
			if !nMap[word] {
				nMap[word] = true
				newIndex = append(newIndex, word)
			}
		}
	}
	//length := int32(len(newIndex))
	for _, word := range message {

		if word >= 97 && word <= 122 {

			qa += string(newIndex[int(word)-97])

		} else {
			qa += " "
		}
	}
	return qa
}
