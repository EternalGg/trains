package getImportance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func GetImportance(employees []*Employee, id int) int {
	m := make(map[int]Employee)

	for i := 0; i < len(employees); i++ {
		m[employees[i].Id] = *employees[i]
	}
	list := m[id].Subordinates
	result := m[id].Importance
	for len(list) != 0 {
		result += m[list[0]].Importance
		if len(m[list[0]].Subordinates) != 0 {
			list = append(list, m[list[0]].Subordinates...)
		}
		list = list[1:]
	}
	return result
}
