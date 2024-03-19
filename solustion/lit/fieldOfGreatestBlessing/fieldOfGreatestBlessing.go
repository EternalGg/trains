package fieldOfGreatestBlessing

import (
	"fmt"
)

func FieldOfGreatestBlessing(forceField [][]int) int {
	m, max := make(map[string]int), 0
	for i := 0; i < len(forceField); i++ {
		x, y, force := float32(forceField[i][0]), float32(forceField[i][1]), forceField[i][2]
		list := make([]string, 0)
		for j := 0; j <= (force/2)+1; j++ {
			neg := x - float32(j)/2

			odd := x + float32(j)/2
			for k := 0; k <= (force/2)+1; k++ {

				up := y + float32(k)/2
				dow := y - float32(k)/2
				nu := fmt.Sprintf("%.1f%.1f", neg, up)
				nd := fmt.Sprintf("%.1f%.1f", neg, dow)
				ou := fmt.Sprintf("%.1f%.1f", odd, up)
				od := fmt.Sprintf("%.1f%.1f", odd, dow)
				list = append(list, ou, nu, nd, od)

			}
		}
		list = diffPrint(list)
		for j := 0; j < len(list); j++ {
			m[list[j]]++
			if m[list[j]] > max {
				max = m[list[j]]

			}
		}
	}
	return max
}

func diffPrint(list []string) []string {
	uMap := make(map[string]bool)
	result := make([]string, 0)
	for i := 0; i < len(list); i++ {
		if !uMap[list[i]] {
			result = append(result, list[i])
			uMap[list[i]] = true
		}
	}
	return result
}

//func FieldOfGreatestBlessing2(forceField [][]int) int {
//	for i := 0; i < len(forceField); i++ {
//		forceField[i][0] += forceField[i][2]
//		forceField[i][0] += forceField[i][2]
//	}
//}
