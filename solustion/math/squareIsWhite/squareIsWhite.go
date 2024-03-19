package squareIsWhite

func squareIsWhite(coordinates string) bool {
	xOdd, yOdd, x, y := false, false, coordinates[1], coordinates[0]
	if x%2 == 0 {
		xOdd = true
	}
	if y%2 == 0 {
		yOdd = true
	}
	if yOdd == xOdd {
		return false
	} else {
		return true
	}

}
