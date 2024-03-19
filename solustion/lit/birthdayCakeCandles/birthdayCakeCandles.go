package birthdayCakeCandles

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	max, count := int32(0), int32(0)

	for _, candle := range candles {
		if candle > max {
			max = candle
			count = 0
			continue
		}
		if candle == max {
			count++
			continue
		}

	}

	return count
}
