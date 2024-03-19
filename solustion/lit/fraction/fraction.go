package fraction

func Fraction(cont []int) []int {

	result := []int{1, cont[len(cont)-1]}
	if len(cont) == 1 {
		result[0] = cont[0]
		result[1] = 1
		return result
	}
	for i := len(cont) - 2; i > 0; i-- {
		result[0] += cont[i] * result[1]
		middle := result[0]
		result[0] = result[1]
		result[1] = middle
	}

	result[0] += cont[0] * result[1]
	if result[0]%result[1] == 0 {
		result[0] = result[0] / result[1]
		result[1] = 1
	}
	return result
}
