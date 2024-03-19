package bestHand

const (
	one   = "Flush"
	two   = "Three of a Kind"
	three = "Pair"
	four  = "High Card"
)

func BestHand(ranks []int, suits []byte) string {
	for _, suit := range suits {
		if suit != suits[0] {
			goto here
		}
	}
	return one
here:
	uMap, frequency := make(map[int]int), 0
	for _, rank := range ranks {
		uMap[rank]++
		if uMap[rank] > frequency {
			frequency = uMap[rank]
		}
	}
	switch frequency {
	case 2:
		return three
	case 3:
		return two
	case 1:
		return four
	default:
		return two
	}
}
