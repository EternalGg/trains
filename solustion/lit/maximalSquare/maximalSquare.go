package maximalSquare

func MaximalSquare(matrix [][]byte) int {
	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			maxlength := 1
		here:
			if maxlength == min(len(matrix), len(matrix[0])) {
				return maxlength * maxlength
			}
			if i+maxlength > len(matrix) || j+maxlength > len(matrix[0]) {
				goto end
			}
			for k := i; k < i+maxlength; k++ {
				for l := j; l < j+maxlength; l++ {
					if matrix[k][l] != '1' {
						goto end
					}
				}
			}
			if maxlength*maxlength > result {
				result = maxlength * maxlength
			}
			maxlength++
			goto here
		end:
		}
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
