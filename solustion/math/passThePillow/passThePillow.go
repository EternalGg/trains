package passThePillow

import "fmt"

func PassThePillow(n int, time int) int {
	half := n - 1
	if time/(half*2) != 0 {
		time = time % (half * 2)
	}
	fmt.Println(time)
	if time >= half {
		time -= half
		for i := 0; i < time; i++ {
			half -= 1
		}
		return half
	} else {
		return half - time
	}
}
