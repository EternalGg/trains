package designerPdfViewer

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	wMap := make(map[int]int32)
	for i := 0; i < len(h); i++ {
		wMap[i] = h[i]
	}
	max := int32(0)
	for i := 0; i < len(word); i++ {
		if wMap[int(word[i])-97] > max {
			max = wMap[int(word[i])-97]
		}
	}
	return int32(len(word)) * max
}
