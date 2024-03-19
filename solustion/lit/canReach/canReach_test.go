package canReach

import "fmt"

type graph struct {
	val      int
	upNode   *graph
	downNode *graph
}

func canReach(arr []int, start int) bool {
	isExit, length, g := make(map[int]bool), len(arr), new(graph)
	fmt.Println(isExit, length)
	for true {
		g.val = arr[start]
		upValue := start + arr[start]
		downValue := start - arr[start]
		if g.val == 0 || arr[upValue] == 0 || arr[downValue] == 0 {
			return true
		}

	}
	return false
}

func Find() {

}
