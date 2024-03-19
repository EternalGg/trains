package StockSpanner

type stock struct {
	Key   int
	Value int
}

type StockSpanner struct {
	list []stock
}

func Constructor() StockSpanner {
	s := new(StockSpanner)
	s.list = make([]stock, 0)
	return *s
}

func (this *StockSpanner) Next(price int) int {
	if len(this.list) == 0 {
		s := new(stock)
		s.Key = price
		s.Value = 1
		this.list = append(this.list, *s)
		return 1
	}

	length := len(this.list) - 1
	for length != -1 && this.list[length].Key <= price {
		length--
	}
	result := 1
	for i := len(this.list) - 1; i > length; i-- {
		result += this.list[i].Value
	}
	s := new(stock)
	s.Key = price
	s.Value = result
	if length == -1 {
		this.list = make([]stock, 0)
	} else {
		this.list = this.list[:length+1]
	}
	this.list = append(this.list, *s)
	return result
}
