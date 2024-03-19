package RandomizedSet

import "math/rand"

type RandomizedSet struct {
	eMap map[int]bool
	kMap map[int]int
	list []int
}

func Constructor() RandomizedSet {
	r := new(RandomizedSet)
	r.eMap = make(map[int]bool)
	r.kMap = make(map[int]int)
	r.list = make([]int, 0)
	return *r
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.eMap[val] {
		return false
	} else {
		this.eMap[val] = true
		this.list = append(this.list, val)
		this.kMap[val] = len(this.list) - 1
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if !this.eMap[val] {
		return false
	} else {
		delete(this.eMap, val)
		this.list[this.kMap[val]] = this.list[len(this.list)-1]
		this.kMap[this.list[this.kMap[val]]] = this.kMap[val]
		this.list = this.list[:len(this.list)-1]
		delete(this.kMap, val)
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	rand := this.list[rand.Intn(len(this.list))]
	return rand
}
