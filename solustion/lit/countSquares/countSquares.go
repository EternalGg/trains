package countSquares

func CountSquares(matrix [][]int) int {
	result := 0
	for i := 1; i <= len(matrix); i++ {
		//max i == 3 ->  1 2 3 -> 1x1 2x2 3x3
		for j := 0; j < i; j++ {
			//left, right := j, j+i
			// j -> 1x1 2x2 3x3
		}
	}
	return result
}
