package licenseKeyFormatting

type License struct {
	Now string
	Max int
	Lab []string
}

func InitLicense(max int) *License {
	l := new(License)
	l.Max = max
	l.Lab = make([]string, 0)
	l.Now = ""
	return l
}

func (l *License) push(b byte) {
	switch {
	case b == 45:
		return
	case b >= 97 && b <= 122:
		l.Now += string(b - 32)
	default:
		l.Now += string(b)
	}
	if len(l.Now) == l.Max {
		l.Lab = append(l.Lab, converseString(l.Now))
		l.Now = ""
	}
}

func converseString(str string) string {
	result := ""
	for _, letter := range str {
		result = string(letter) + result
	}
	return result
}

func (l *License) format() string {

	if len(l.Now) != 0 {
		l.Lab = append(l.Lab, converseString(l.Now))
	}
	if len(l.Lab) == 0 {
		return ""
	}
	result := ""
	for i := len(l.Lab) - 1; i > 0; i-- {
		result += l.Lab[i]
		result += "-"
	}
	result += l.Lab[0]
	return result
}

func LicenseKeyFormatting(s string, k int) string {
	l := InitLicense(k)
	for i := len(s) - 1; i >= 0; i-- {
		l.push(s[i])
	}
	return l.format()
}
