package convertToBase7

import "fmt"

func convertToBase7(num int) string {
	blist := []byte{}

	for num/7 >= 1 {
		blist = append(blist, byte(num%7)-49)
		num = num / 7
	}
	fmt.Println(blist)
	return string(blist)
}
