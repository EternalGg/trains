package twoOutOfThree

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	n1Map, n2Map, n3Map, cMap, result := make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make([]int, 0)
	l2, l3 := make([]int, 0), make([]int, 0)
	n1Map, _ = listToMap(n1Map, nums1)
	n2Map, l2 = listToMap(n2Map, nums2)
	n3Map, l3 = listToMap(n3Map, nums3)

	// l1 vs l2
	cMap, result = compareListInMap(n1Map, cMap, l2, result)
	// l1 vs l3
	cMap, result = compareListInMap(n1Map, cMap, l3, result)
	// l2 vs l3
	cMap, result = compareListInMap(n2Map, cMap, l3, result)
	return result
}

func listToMap(Map map[int]bool, list []int) (map[int]bool, []int) {
	listrusult := make([]int, 0)
	for _, value := range list {
		if Map[value] == false {
			Map[value] = true
			listrusult = append(listrusult, value)
		}
	}
	return Map, listrusult
}

func compareListInMap(c2, check map[int]bool, l1, result []int) (map[int]bool, []int) {
	for _, num := range l1 {
		if c2[num] == true && !check[num] {
			check[num] = true
			result = append(result, num)
		}
	}
	return check, result
}
