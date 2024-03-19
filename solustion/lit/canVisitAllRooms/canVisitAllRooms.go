package canVisitAllRooms

func CanVisitAllRooms(rooms [][]int) bool {
	rMap := make(map[int]bool)
	GetAllKey(0, rMap, rooms)
	for i := 1; i < len(rooms); i++ {
		if !rMap[i] {
			continue
		}
		for j := 1; j < len(rooms[i]); j++ {
			if !rMap[rooms[i][j]] {
				GetAllKey(rooms[i][j], rMap, rooms)
			}
		}
	}
	for i := 1; i < len(rooms); i++ {
		if !rMap[i] {
			return false
		}
	}
	return true
}

func GetAllKey(room int, m map[int]bool, rooms [][]int) {
	for i := 0; i < len(rooms[room]); i++ {
		if !m[rooms[room][i]] {
			m[rooms[room][i]] = true
			GetAllKey(rooms[room][i], m, rooms)
		}
	}
}
