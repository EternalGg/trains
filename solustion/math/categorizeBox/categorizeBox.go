package categorizeBox

import "math"

func CategorizeBox(length int, width int, height int, mass int) string {
	volume := length * width * height
	tenPowFour, tenPowNine, isBulky := math.Pow(10, 4), math.Pow(10, 9), false
	if float64(length) >= tenPowFour || float64(width) >= tenPowFour || float64(height) >= tenPowFour || float64(volume) >= tenPowNine {
		isBulky = !isBulky
	}
	if isBulky && mass < 100 {
		return "Bulky"
	}
	if !isBulky && mass >= 100 {
		return "Heavy"
	}
	if isBulky && mass >= 100 {
		return "Both"
	}

	return "Neither"

}
