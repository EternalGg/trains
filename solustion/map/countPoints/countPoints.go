package countPoints

import "fmt"

func countPoints(points [][]int, queries [][]int) []int {
	m := make(map[int][]int)
	for i := 0; i < len(points); i++ {
		if len(m[points[i][0]]) == 0 {
			m[points[i][0]] = make([]int, 0)
		}
		m[points[i][0]] = append(m[points[i][0]], points[i][1])
	}

	result := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		for j := 0; j <= queries[i][2]; j++ {
			up, down, left, right := queries[i][0]-j, queries[i][0]+j, queries[i][1]-j, queries[i][1]+j
			for k := up; k <= down; k++ {
				fmt.Println(k)
				if len(m[k]) == 0 {
					break
				} else {
					for l := 0; l < len(m[k]); l++ {
						//fmt.Println(m[k][l], i, l)
						if m[k][l] >= left || m[k][l] <= right {
							result[i]++
						}
					}
				}
			}
		}
	}
	return result
}

func countPoints1(rings string) int {
	r, g, b := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	for i := 0; i < len(rings)/2; i++ {
		switch rings[i*2] {
		case 'R':
			if !r[int(rings[(i*2)+1]-48)] {
				r[int(rings[(i*2)+1]-48)] = true
			}
		case 'G':
			if !g[int(rings[(i*2)+1]-48)] {
				g[int(rings[(i*2)+1]-48)] = true
			}
		case 'B':
			if !b[int(rings[(i*2)+1]-48)] {
				b[int(rings[(i*2)+1]-48)] = true
			}
		}
	}
	result := 0
	for i := 0; i < 9; i++ {
		if r[i] && g[i] && b[i] {
			result++
		}
	}
	return result
}
