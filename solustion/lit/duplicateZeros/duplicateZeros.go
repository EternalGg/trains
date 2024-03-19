package duplicateZeros

func DuplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			befor := arr[:i]
			befor = append(befor, 0)
			last := arr[i : len(arr)-1]
			arr = append(befor, last...)
			i++
		}
	}
}
