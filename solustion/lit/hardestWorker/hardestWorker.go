package hardestWorker

func HardestWorker(n int, logs [][]int) int {
	max, reslut := 0, 0
	for i := len(logs) - 1; i > 0; i-- {
		now := logs[i][1] - logs[i-1][1]
		if now == max && logs[i][0] < logs[i-1][0] {
			max = now
			reslut = logs[i][0]
			continue
		}
		if now > max {
			max = now
			reslut = logs[i][0]
		}
	}
	if logs[0][1] == max && logs[0][0] < reslut {
		return logs[0][0]
	}
	if logs[0][1] > max {
		return logs[0][0]
	}
	return reslut
}
