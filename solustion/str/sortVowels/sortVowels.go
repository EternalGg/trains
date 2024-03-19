package sortVowels

import "sort"

func sortVowels(s string) string {
	V, NV := []byte{}, []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			V = append(V, s[i])
		} else {
			NV = append(NV, s[i])
		}
	}
	sort.Slice(V, func(i, j int) bool {
		return V[i] < V[j]
	})
	result := []byte{}

	NP, VP := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			result = append(result, V[VP])
			VP++
		} else {
			result = append(result, NV[NP])
			NP++
		}
	}
	return string(result)
}
