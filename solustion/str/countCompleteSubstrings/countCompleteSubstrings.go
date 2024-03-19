package countCompleteSubstrings

import (
	"math"
)

func CountCompleteSubstrings(word string, k int) int {

	result := 0

	for multi := 1; multi <= len(word)/k; multi++ {
		l := k * multi
		if l == 1 {
			result += len(word)
			continue
		}
		left, right := 0, l
		m := map[int]int{}
		for i := 0; i < l; i++ {
			m[int(word[i])]++
		}
		if len(m) == multi {
			for _, i2 := range m {
				if i2 != k {
					goto here
				}
			}
			for i := left; i < right; i++ {
				if i == left {
					if math.Abs(float64(int(word[i])-int(word[i+1]))) <= 2 {
						continue
					} else {
						goto here
					}
				}
				if i == right-1 {
					if math.Abs(float64(int(word[i])-int(word[i-1]))) <= 2 {
						continue
					} else {
						goto here
					}
				}
				if math.Abs(float64(int(word[i])-int(word[i-1]))) <= 2 && math.Abs(float64(int(word[i])-int(word[i+1]))) <= 2 {
					continue
				} else {
					goto here
				}
			}
			result++
		here:
		}
		for right != len(word) {

			m[int(word[left])]--
			if m[int(word[left])] == 0 {
				delete(m, int(word[left]))
			}
			left++
			m[int(word[right])]++
			right++
			if len(m) == multi {
				for _, i2 := range m {
					if i2 != k {
						goto here1
					}
				}
				for i := left; i < right; i++ {
					if i == left {
						if math.Abs(float64(int(word[i])-int(word[i+1]))) <= 2 {
							continue
						} else {
							goto here1
						}
					}
					if i == right-1 {
						if math.Abs(float64(int(word[i])-int(word[i-1]))) <= 2 {
							continue
						} else {
							goto here1
						}
					}
					if math.Abs(float64(int(word[i])-int(word[i-1]))) <= 2 && math.Abs(float64(int(word[i])-int(word[i+1]))) <= 2 {
						continue
					} else {
						goto here1
					}
				}
				result++
			here1:
			}
		}
	}

	return result
}
