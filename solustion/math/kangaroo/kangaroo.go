package kangaroo

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	for x1 < x2 {
		x1 += v1
		x2 += v2
	}
	if x1 != x2 {
		return "NO"
	} else {
		return "YES"
	}
}
