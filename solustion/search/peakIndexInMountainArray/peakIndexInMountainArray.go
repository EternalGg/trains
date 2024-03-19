package peakIndexInMountainArray

func PeakIndexInMountainArray(arr []int) int {
	head, last, now := 0, len(arr)-1, (len(arr)-1)/2
	for !(arr[now] > arr[now-1] && arr[now] > arr[now+1]) {
		if arr[now-1] > arr[now] {
			last = now
		} else {
			head = now
		}
		now = (last-head)/2 + head
	}
	return now
}
