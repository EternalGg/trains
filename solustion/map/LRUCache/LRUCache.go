package LRUCache

type LRUCache struct {
	List []int
	eMap map[int]bool
	vMap map[int]int
	Max  int
}

func Constructor(capacity int) LRUCache {
	l := new(LRUCache)
	l.Max = capacity
	l.List = make([]int, 0)
	l.eMap = make(map[int]bool)
	l.vMap = make(map[int]int)
	return *l
}

func (this *LRUCache) Get(key int) int {
	if this.eMap[key] {
		point, pointer := make([]int, 1), 0
		for i := 0; i < len(this.List); i++ {
			if this.List[i] == key {
				point[0] = this.List[i]
				pointer = i
			}
		}
		head := this.List[:pointer]
		last := this.List[pointer+1:]
		point = append(point, head...)
		point = append(point, last...)
		this.List = point
		return this.vMap[key]
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.eMap[key] {
		this.vMap[key] = value
	} else {
		if len(this.List) < this.Max {
			this.List = append(this.List, key)
		} else {
			//满了
			this.eMap[this.List[len(this.List)-1]] = false
			this.List[len(this.List)-1] = key
		}
		this.eMap[key] = true
		this.vMap[key] = value
	}
	point, pointer := make([]int, 1), 0
	for i := 0; i < len(this.List); i++ {
		if this.List[i] == key {
			point[0] = this.List[i]
			pointer = i
		}
	}
	head := this.List[:pointer]
	last := this.List[pointer+1:]
	point = append(point, head...)
	point = append(point, last...)
	this.List = point
}
