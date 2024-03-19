package permute

//             1
// 3         2       4
// 2 4    34    2 3
//4   2  4  3  3      2
//2 3 1
//3 2 1

func Permute(nums []int) [][]int {
	matrix := make([][]int, 0)
	length := len(nums)
	count := 1
	for i := length; i > 0; i-- {
		count *= i
	}
	for i := 0; i < count; i++ {
		matrix = append(matrix, make([]int, len(nums)))
	}

	for i := 0; i < len(nums); i++ {
		j, k := 0, 0
		for i := 0; i < k; i++ {
			j++
		}
	}
	return matrix
}
