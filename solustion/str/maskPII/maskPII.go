package maskPII

import "fmt"

func MaskPII(s string) string {
	answer := ""
	if (s[0] >= 48 && s[0] <= 57) || s[0] == '(' || s[0] == ')' || s[0] == '-' || s[0] == '+' {
		// phone
		num := make([]byte, 0)
		for i := 0; i < len(s); i++ {
			if s[i] >= 48 && s[i] <= 57 {
				num = append(num, s[i])
			}
		}
		for i := len(num) - 1; i >= len(num)-4; i-- {
			answer = string(num[i]) + answer
		}
		answer = "***-***-" + answer

		switch len(num) - 10 {
		case 1:
			answer = "+*-" + answer
		case 2:
			answer = "+**-" + answer
		case 3:
			answer = "+***-" + answer
		}
	} else {
		// mail
		at, sToB := 0, []byte(s)
		for i := 0; i < len(s); i++ {
			if s[i] >= 65 && s[i] <= 90 {
				sToB[i] = s[i] + 32
			}
		}
		first, last := sToB[0], byte(0)
		for i := 0; i < len(s); i++ {
			if s[i+1] == '@' {
				last = sToB[i]
				at = i + 1
				break
			}
		}

		fmt.Println(string(first), string(last))
		answer = string(first) + "*****" + string(last) + string(sToB[at:])
	}
	return answer
}
