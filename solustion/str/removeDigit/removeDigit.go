package removeDigit

func RemoveDigit(number string, digit byte) string {
	b := []byte(number)
	for i := 0; i <= len(number)-2; i++ {
		if b[i] == digit {
			if b[i] < b[i+1] {
				head := string(b[:i])
				last := string(b[i+1:])
				return head + last
			}
		}
	}
	for i := len(number); i >= 0; i-- {
		if b[i] == digit {
			head := string(b[:i])
			last := string(b[i+1:])
			return head + last
		}
	}
	return number
	//max, maxKey := int64(0), 0
	//for i := 0; i < len(sList); i++ {
	//	now, _ := strconv.ParseInt(sList[i], 10, 64)
	//	if now > max {
	//		max = now
	//		maxKey = i
	//	}
	//}

}
