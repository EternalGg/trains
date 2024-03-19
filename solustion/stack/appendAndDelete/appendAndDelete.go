package appendAndDelete

func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	lastDiff := 0
	for i := 0; i < len(t); i++ {
		if t[i] != s[i] {
			lastDiff = i
		}
	}
	if lastDiff == 0 && s[0] == t[0] {
		lastDiff = len(t) - 1
	}
	cut := 0
	cut += len(s) - lastDiff - 1
	if len(t) > len(s) {
		cut += len(t) - lastDiff
	}
	if int32(cut) <= k {
		return "Yes"
	} else {
		return "No"
	}
}
