package reverseEvenLengthGroups

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	list, key, slist := make([][]*ListNode, 0), 1, make([]*ListNode, 0)
	for head != nil {
		slist = append(slist, head)
		if key == len(slist) {
			if key%2 == 0 && key != 1 {
				newlist := make([]*ListNode, key)
				for i := key - 1; i >= 0; i-- {
					newlist[i] = slist[key-i-1]
				}
				list = append(list, newlist)
			} else {
				list = append(list, slist)
			}
			slist = make([]*ListNode, 0)
			key++
		}
	}

	if len(slist) != 0 {
		if key%2 == 0 && key != 1 {
			newlist := make([]*ListNode, key)
			for i := key - 1; i >= 0; i-- {
				newlist[i] = slist[key-i-1]
			}
			list = append(list, newlist)
		} else {
			list = append(list, slist)
		}
		key++
	}

	if len(list) >= 2 {
		list[0][0].Next = list[1][0]
	}
	fmt.Println(list)
	for i := 1; i < len(list)-1; i++ {
		for j := 0; j < len(list[i]); j++ {
			if !(j == len(list[i])-1) {
				list[i][j].Next = list[i][j+1]
			} else {
				list[i][j].Next = list[i+1][0]
			}
		}
	}

	for i := 0; i < len(list[len(list)-1])-1; i++ {
		list[len(list)-1][i].Next = list[len(list)-1][i+1]
	}
	list[len(list)-1][len(list[len(list)-1])-1].Next = nil
	return list[0][0]
}
