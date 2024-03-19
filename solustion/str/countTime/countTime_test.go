package countTime

import (
	"fmt"
	"testing"
)

func countTime(time string) int {
	num, isQ, qa, aCount := make([]int32, 5), make([]bool, 5), 0, 0

	for key, ward := range time {
		if key == 2 {
			continue
		}
		if ward != 63 {
			num[key] = ward - int32(48)
		} else {
			isQ[key] = true
			aCount++
		}
	}
	if aCount == 0 {
		return 0
	}

	// 小时判断
	if !isQ[0] && !isQ[1] {
		qa = 1
		goto minute
	}
	if isQ[0] && isQ[1] {
		qa = 24
		goto minute
	}
	// 第一位空 第二位存在 判断是不是《=4 如果《=4 qa=3
	if isQ[0] && !isQ[1] {
		if num[1] <= 4 {
			qa = 3
		} else {
			qa = 2
		}
		goto minute
	}
	// 第二位空 第一位存在 判断是不是 》=4 如果《=4 qa=3
	if !isQ[0] && isQ[1] {
		if num[0] == 2 {
			qa = 4
		} else {
			qa = 10
		}
		goto minute
	}
minute:
	if !isQ[3] && !isQ[4] {
		return qa
	}
	if isQ[3] && isQ[4] {
		return qa * 60
	}
	// 第一位空 第二位存在 判断是不是《=4 如果《=4 qa=3
	if isQ[3] && !isQ[4] {
		return qa * 6
	}

	if !isQ[3] && isQ[4] {
		return qa * 10
	}

	return qa
}

func TestCountTime(t *testing.T) {
	fmt.Println(countTime("0?:0?"))
}
