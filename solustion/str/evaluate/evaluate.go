package evaluate

import "fmt"

func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string)
	for i := 0; i < len(knowledge); i++ {
		m[knowledge[i][0]] = m[knowledge[i][1]]
	}

	pair := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == ')' {
			pair = append(pair, i)
		}
	}

	for i := 0; i < len(pair); i += 2 {
		fmt.Println(s)
		fmt.Println(s[:pair[i]] + ("?") + s[pair[i+1]+1:])
		if m[s[pair[i]+1:pair[i+1]]] == "" {
			s = s[:pair[i]] + ("?") + s[pair[i+1]+1:]
		} else {
			s = s[:pair[i]] + m[s[pair[i]+1:pair[i+1]]] + s[pair[i+1]+1:]
		}
	}
	return s
}
