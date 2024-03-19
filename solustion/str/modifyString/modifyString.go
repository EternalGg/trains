package modifyString

func modifyString(s string) string {
	str := s
	for key, value := range s {
		if value == '?' {
			switch key {
			case 0:
				if s[1] == 122 {
					str += string(s[1] - 1)
				} else {
					str += string(s[1] + 1)
				}
			case len(s) - 1:
			default:

			}
		} else {

		}
	}
	return s
}
