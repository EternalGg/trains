package maximumWhiteTiles

import "sort"

func MaximumWhiteTiles(tiles [][]int, carpetLen int) int {
	m := make(map[int]int)
	list := []int{}
	for i := 0; i < len(tiles); i++ {
		m[tiles[i][0]] = tiles[i][1] - tiles[i][0] + 1
		list = append(list, tiles[i][0])
	}
	sort.Ints(list)
	max := 0
	for i := 0; i < len(list); i++ {
		count := m[list[i]]
		if count == carpetLen {
			return count
		}
		tl := carpetLen - count
		ti := i

		for tl > 0 {
			//j
			if ti+1 == len(list) {
				goto here
			}
			split := list[ti+1] - (m[list[ti]] + list[ti])
			if split >= tl {
				break
			} else {
				tl -= split

				if tl < m[list[ti+1]] {
					count += tl
					break
				} else {
					count += m[list[ti+1]]
					tl -= m[list[ti+1]]
					ti++
				}
			}
		}
	here:
		if count > max {
			max = count
		}

	}
	return max
}
