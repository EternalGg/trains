package FrequencyTracker

type FrequencyTracker struct {
	nMap map[int]int
	fMap map[int]int
}

func Constructor() FrequencyTracker {
	f := new(FrequencyTracker)
	f.nMap = map[int]int{}
	f.fMap = map[int]int{}
	return *f
}

func (this *FrequencyTracker) Add(number int) {
	if this.nMap[number] != 0 {
		this.fMap[this.nMap[number]]--
	}
	this.nMap[number]++
	this.fMap[this.nMap[number]]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.nMap[number] <= 0 {
		return
	}
	this.fMap[this.nMap[number]]--
	this.nMap[number]--
	this.fMap[this.nMap[number]]++
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	if frequency == 0 {
		return true
	}
	if this.fMap[frequency] > 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
