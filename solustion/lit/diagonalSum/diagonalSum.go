package diagonalSum

import "fmt"

func DiagonalSum(mat [][]int) int {
	length, heigth, count := len(mat[0]), len(mat), 0
	if heigth == 1 && length == 1 {
		return mat[0][0]
	}

	for i := 0; i < length; i++ {
		if length%2 != 0 && length/2 == i {
			count += mat[i][i]
			continue
		}
		count += mat[i][i]
		count += mat[i][length-i-1]
		fmt.Println(count)
	}

	return count
}
