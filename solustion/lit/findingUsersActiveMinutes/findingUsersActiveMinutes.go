package findingUsersActiveMinutes

type User struct {
	count int
	mMap  map[int]bool
}

func FindingUsersActiveMinutes(logs [][]int, k int) []int {
	uMap, ulist := make(map[int]*User), make([]int, 0)
	for _, log := range logs {

		if uMap[log[0]] == nil {
			ulist = append(ulist, log[0])
			newUser := new(User)
			newUser.count = 1
			newUser.mMap = make(map[int]bool)
			newUser.mMap[log[1]] = true
			uMap[log[0]] = newUser
		} else {
			if !uMap[log[0]].mMap[log[1]] {
				uMap[log[0]].mMap[log[1]] = true
				uMap[log[0]].count++
			}
		}
	}
	result := make([]int, k)
	for i := 0; i < len(ulist); i++ {
		if uMap[ulist[i]] == nil {
			continue
		} else {
			result[uMap[ulist[i]].count-1]++
		}
	}
	return result
}
