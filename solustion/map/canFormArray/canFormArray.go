package canFormArray

func CanFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = i + 1
	}
	for i := 0; i < len(pieces); i++ {
		for j := 0; j < len(pieces[i])-1; j++ {
			if m[pieces[i][j]]+1 != m[pieces[i][j+1]] {
				return false
			}
		}
		for j := 0; j < len(pieces[i]); j++ {
			if m[pieces[i][j]] == 0 {
				return false
			}
		}
	}
	return true
}
