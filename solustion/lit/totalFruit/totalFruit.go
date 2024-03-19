package totalFruit

func TotalFruit(fruits []int) int {
	fList, cList, last, count := make([]int, 0), make([]int, 0), fruits[0], 0
	for i := 0; i < len(fruits); i++ {
		if fruits[i] != last {
			fList = append(fList, last)
			cList = append(cList, count)
			last = fruits[i]
			count = 1
		} else {
			count++
		}
	}
	fList = append(fList, last)
	cList = append(cList, count)

	result, temp := 0, 0
	for i := 0; i < len(fList)-1; i++ {
		temp = 0
		first, second := fList[i], fList[i+1]
		temp = cList[i] + cList[i+1]
		for j := i + 2; j < len(fList); j++ {
			if fList[j] == first || fList[j] == second {
				temp += cList[j]
			} else {
				break
			}
		}
		if temp > result {
			result = temp
		}
	}
	if temp > result {
		result = temp
	}
	return result
}
