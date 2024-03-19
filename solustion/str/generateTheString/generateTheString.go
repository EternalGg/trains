package generateTheString

func generateTheString(n int) string {
	qa := ""
	if n%2 == 0 {
		for i := 0; i < n-2; i++ {
			qa += "a"
		}
		qa += "b"
	} else {
		for i := 0; i < n-1; i++ {
			qa += "a"
		}
	}
	return qa
}
