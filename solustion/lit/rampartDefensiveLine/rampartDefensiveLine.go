package rampartDefensiveLine

func RampartDefensiveLine(rampart [][]int) int {
	min := Min(rampart[0][1]-rampart[0][0], rampart[len(rampart)-1][1]-rampart[len(rampart)-1][0])
	for i := 1; i < len(rampart)-1; i++ {
		if (rampart[i][0]-rampart[i-1][1])+(rampart[i+1][0]-rampart[i][1]) < min {
			min = (rampart[i][0] - rampart[i-1][1]) + (rampart[i+1][0] - rampart[i][1])
		}
	}
	return min
}

func Min(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
