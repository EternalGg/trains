package isValidSudoku

import "fmt"

func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			bMap := make(map[byte]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[i*3+k][j*3+l] != '.' {
						if bMap[board[i*3+k][j*3+l]] {
							return false
						}
						bMap[board[i*3+k][j*3+l]] = true
					}
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		raw := make(map[byte]bool)
		law := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if raw[board[i][j]] {
					return false
				}
				raw[board[i][j]] = true
			}

			if board[j][i] != '.' {
				if law[board[j][i]] {
					return false
				}
				law[board[j][i]] = true
			}

		}
		fmt.Println("zzzz")
	}
	return true
}
