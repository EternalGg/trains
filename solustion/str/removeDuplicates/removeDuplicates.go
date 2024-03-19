package removeDuplicates

type Stack struct {
	nList []int32
}

func (s *Stack) push(num int32) {
	if len(s.nList) == 0 {
		s.nList = append(s.nList, num)
		return
	}
	if s.nList[len(s.nList)-1] == num {
		s.nList = s.nList[:len(s.nList)-1]
		return
	}
	s.nList = append(s.nList, num)

}

func stackInit() *Stack {
	s := new(Stack)
	s.nList = make([]int32, 0)
	return s
}

func removeDuplicates(s string) string {
	stack := stackInit()
	for _, value := range s {
		stack.push(value)
	}

	return string(stack.nList)
}

func removeDuplicatesTwo(nums []int) int {
	//realPoint, last, t := 0, -1, 0
	//
	//for i := 0; i < len(nums); i++ {
	//	if last == nums[i] {
	//
	//	} else {
	//
	//	}
	//}
	//nums = nums[:realPoint]
	//return len(nums)
	return 0
}
