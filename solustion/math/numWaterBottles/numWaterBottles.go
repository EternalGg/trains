package numWaterBottles

func numWaterBottles(numBottles int, numExchange int) int {
	remain, dived, result := 0, 0, numBottles

	remain = numBottles % numExchange
	dived = numBottles / numExchange
	result += dived
	remain += dived
	for remain >= numExchange {
		// 瓶子兑换饮料
		dived = remain / numExchange
		numBottles = dived
		// 计算新的空瓶子
		remain -= numExchange * dived
		remain += numBottles
		result += numBottles
	}
	return result
}
