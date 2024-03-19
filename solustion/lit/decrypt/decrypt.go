package decrypt

import "math"

func Decrypt(code []int, k int) []int {
	result := []int{}
	if k == 0 {
		return make([]int, len(code))
	} else if k > 0 {
		newCode := append(code, code[:k]...)
		for i := 0; i < len(newCode)-k; i++ {
			c := 0
			for j := i + 1; j < k+i+1; j++ {
				c += newCode[j]
			}
			result = append(result, c)
		}
	} else if k < 0 {
		newCode := append(code[len(code)-int(math.Abs(float64(k))):], code...)
		for i := 0; i < len(newCode)-int(math.Abs(float64(k))); i++ {
			c := 0
			for j := i; j < int(math.Abs(float64(k)))+i; j++ {
				c += newCode[j]
			}
			result = append(result, c)
		}
	}
	return result
}
