package breakingRecords

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	good, bad, gb, bb := scores[0], scores[0], 0, 0
	for i := 0; i < len(scores); i++ {
		if scores[i] > good {
			good = scores[i]
			gb++
		}
		if scores[i] < bad {
			bad = scores[i]
			bb++
		}
	}
	return []int32{int32(gb), int32(bb)}
}
