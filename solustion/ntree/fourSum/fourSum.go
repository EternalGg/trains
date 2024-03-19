package fourSum

type Node struct {
	Val      int
	count    int
	layer    int
	Children []*Node
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	headList := []Node{}
	for i := 0; i < len(nums)-3; i++ {
		n := Node{Val: nums[i], count: nums[i], layer: 1, Children: []*Node{}}
		headList = append(headList, n)
	}
	for i := 0; i < len(headList); i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			twon := Node{Val: nums[i], count: nums[i] + headList[i].Val, Children: []*Node{}}
			headList[i].Children = append(headList[i].Children, &twon)
			for k := i + 2; k < len(nums)-1; k++ {
				threen := Node{Val: nums[i], count: nums[i] + headList[i].Val, Children: []*Node{}}
				headList[i].Children = append(twon.Children, &threen)
				for l := i + 3; l < len(nums); l++ {

				}
			}
		}
	}
	return result
}
