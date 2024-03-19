package vowelStrings

import "testing"

func TestVowelStrings(t *testing.T) {
	str := []string{"aba", "bcb", "ece", "aa", "e"}
	num := [][]int{}
	l1, l2, l3 := []int{0, 2}, []int{1, 4}, []int{1, 1}
	num = append(num, l1)
	num = append(num, l2)
	num = append(num, l3)
	VowelStrings(str, num)
	VowelStrings2([]string{"are", "amy", "u"}, 0, 2)
}
