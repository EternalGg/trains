package StockSpanner

import "testing"

func TestStockSpanner_Next(t *testing.T) {
	s := Constructor()
	s.Next(100)
	s.Next(60)
	s.Next(80)
	s.Next(70)
	s.Next(22)
	s.Next(445)
	s.Next(98)
	s.Next(3222)
}
