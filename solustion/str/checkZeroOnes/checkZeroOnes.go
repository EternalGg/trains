package checkZeroOnes

import "fmt"

func checkZeroOnes(s string) bool {
	oneMax, zeroMax, lastIszero, freq := 0, 0, false, 0
	for _, value := range s {
		if value == 48 {
			if lastIszero {
				freq++
			} else {
				if oneMax < freq {
					oneMax = freq
				}
				lastIszero = true
				freq = 1
			}
		} else {
			if !lastIszero {
				freq++
			} else {
				if zeroMax < freq {
					zeroMax = freq
				}
				lastIszero = false
				freq = 1
			}
		}
	}
	if lastIszero {
		if freq > zeroMax {
			zeroMax = freq
		}
	} else {
		if freq > oneMax {
			oneMax = freq
		}
	}
	fmt.Println(oneMax, zeroMax)
	if oneMax > zeroMax {
		return true
	} else {
		return false
	}
}
