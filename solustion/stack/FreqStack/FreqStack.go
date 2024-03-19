package FreqStack

type TimeList struct {
	// 记录该频率的数字
	//Time map[int]bool
	// 记录该数字添加的时间
	TMap map[int][]uint
}

type FreqStack struct {
	// 频率对应map和list
	Im map[int]TimeList
	// 数字的频率
	FTime map[int]int
	PTime uint
	MaxF  uint
}

func Constructor() FreqStack {
	return FreqStack{map[int]TimeList{}, map[int]int{}, 0, 0}
}

func (this *FreqStack) Push(val int) {
	this.PTime++
	if this.FTime[val] == 0 {
		if len(this.Im) == 0 {
			t := TimeList{map[int][]uint{}}
			this.Im[1] = t
		}
		//this.Im[1].Time[val] = 1
		this.FTime[val] = 1

		this.Im[1].TMap[val] = []uint{this.PTime}
	} else {

		this.FTime[val]++

		F := this.FTime[val]
		list := this.Im[F-1].TMap[val]
		list = append(list, this.PTime)
		delete(this.Im[F-1].TMap, val)
		if len(this.Im[F].TMap) == 0 {
			this.Im[F] = TimeList{map[int][]uint{}}
		}
		this.Im[F].TMap[val] = list

	}
	if this.FTime[val] > int(this.MaxF) {
		this.MaxF = uint(this.FTime[val])
	}
}

func (this *FreqStack) Pop() int {
	maxPt, maxValue := uint(0), 0
	for i, uints := range this.Im[int(this.MaxF)].TMap {
		if uints[len(uints)-1] > maxPt {
			maxPt = uints[len(uints)-1]
			maxValue = i
		}
	}
	this.FTime[maxValue]--
	if this.MaxF > 1 {

		list := this.Im[int(this.MaxF)].TMap[maxValue][:this.MaxF-1]
		this.Im[int(this.MaxF-1)].TMap[maxValue] = list
	}

	delete(this.Im[int(this.MaxF)].TMap, maxValue)
	if len(this.Im[int(this.MaxF)].TMap) == 0 {
		delete(this.Im, int(this.MaxF))
		this.MaxF--
	}
	return maxValue
}
