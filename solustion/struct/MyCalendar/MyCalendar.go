package MyCalendar

type Node struct {
	Start int
	End   int
	Next  *Node
	Front *Node
}

type MyCalendar struct {
	N *Node
}

func Constructor() MyCalendar {
	return MyCalendar{N: &Node{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
