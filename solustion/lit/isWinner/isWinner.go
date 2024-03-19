package isWinner

import "fmt"

func isWinner(player1 []int, player2 []int) int {
	s1, s2 := player1[0], player2[0]
	if player1[0] == 10 {
		s1 += player1[1] * 2
	}
	for i := 2; i < len(player1); i++ {
		if player1[i-1] == 10 || player1[i-2] == 10 {
			s1 += player1[i] * 2
		}
	}

	if player2[0] == 10 {
		s2 += player2[1] * 2
	}
	for i := 2; i < len(player2); i++ {
		if player2[i-1] == 10 || player2[i-2] == 10 {
			s2 += player2[i] * 2
		}
	}
	fmt.Println(s1, s2)
	if s1 > s2 {
		return 1
	} else if s1 < s2 {
		return 2
	} else {
		return 0
	}
}
