package ProductOfNumbers

type ProductOfNumbers struct {
	list   []int
	length int
}

func Constructor() ProductOfNumbers {
	obj := new(ProductOfNumbers)
	obj.list = make([]int, 0)
	obj.length = 0
	return *obj
}

func (this *ProductOfNumbers) Add(num int) {
	this.list = append(this.list, num)
	this.length++
	return
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	now := 1
	for i := this.length; i > this.length-k; i-- {
		now *= this.list[i]
	}
	return now
}
