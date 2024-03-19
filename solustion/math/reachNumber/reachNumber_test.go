package reachNumber

import (
	"fmt"
	"testing"
)

func ReachNumber(target int) int {
	if target == 0 {
		return 0
	}
	acount, flag, sub := 0, 0, 0
	if target <= 0 {
		for true {
			flag += 1
			acount -= flag
			if acount <= target {
				sub = target - acount
				break
			}
		}
	} else {
		for true {
			flag += 1
			acount += flag
			if acount >= target {

				sub = acount - target
				break
			}
		}
	}
	if sub == 0 {
		return flag
	}
	if sub == flag {
		return flag
	}
	if sub <= flag {
		return flag - 1
	} else {
		for i := flag; i >= 0; i++ {
			sub -= flag
			if sub <= flag-1 {
				return flag - 1
			}
		}
	}
	return flag
}

func TestReachNumber(t *testing.T) {
	a := ReachNumber(3)
	fmt.Println(a)
}
