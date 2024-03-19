package judgeCircle

func JudgeCircle(moves string) bool {
	x, y := 0, 0

	for _, letters := range moves {
		switch letters {
		case 68:
			y--
		case 76:
			x--
		case 82:
			x++
		case 85:
			y++
		}
	}
	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}
}
