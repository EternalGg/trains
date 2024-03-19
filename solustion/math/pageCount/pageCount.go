package pageCount

import "math"

func pageCount(n int32, p int32) int32 {
	//// Write your code here
	//if p == 0 || p == 1 {
	//	return 0
	//}
	//if p == n || (n%2 != 0 && p == n-1) {
	//	return 0
	//}
	//hc, min := 0, 0
	//for i := -1; i < int(n); i += 2 {
	//	hc++
	//	if int(p) == i || int(p) == i-1 {
	//		min = hc
	//		hc = 0
	//		break
	//	}
	//}
	//if n%2 == 0 {
	//	n++
	//}
	//for i := n; i > 0; i -= 2 {
	//	hc++
	//	if p == i-1 || p == i {
	//		if min > hc {
	//			hc = min
	//		}
	//		break
	//	}
	//}
	//return int32(min)
	p = p / 2
	n = n / 2

	return int32(math.Min(math.Abs(float64((n - p))), float64(p)))
}
