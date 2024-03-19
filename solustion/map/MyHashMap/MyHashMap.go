package MyHashMap

type MyHashMap struct {
	mMap map[int]int
	eMap map[int]bool
}

func Constructor() MyHashMap {
	m := new(MyHashMap)
	m.mMap = make(map[int]int)
	m.eMap = make(map[int]bool)
	return *m
}

func (this *MyHashMap) Put(key int, value int) {
	this.mMap[key] = value
	this.eMap[key] = true
}

func (this *MyHashMap) Get(key int) int {
	if this.eMap[key] {
		return this.mMap[key]
	} else {
		return -1
	}
}

func (this *MyHashMap) Remove(key int) {
	this.eMap[key] = false
	this.mMap[key] = 0
}
