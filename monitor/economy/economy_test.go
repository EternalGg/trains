package economy

import (
	"fmt"
	"testing"
	"train/monitor/economy/shop"
	crazy2 "train/monitor/hero/crazy"
	ele2 "train/monitor/hero/ele"
)

// 测试完成
func TestDayPast(t *testing.T) {
	e := EconomyInit()
	e.SaveToBank(3000)
	e.DayPast()
	e.DayPast()
	e.DayPast()
	fmt.Println(e.ShowMoney())
	fmt.Println(e.TakeToBank(2000))
	fmt.Println(e.ShowMoney())
	e.DayPast()
	e.DayPast()
	e.DayPast()
	fmt.Println(e.ShowMoney())
}

func TestLoan(t *testing.T) {
	e := EconomyInit()
	fmt.Println(e.Bank.Saving)
	ele := shop.MakeCard(*ele2.EleHeroInit())
	crazy := shop.MakeCard(*crazy2.CrazyHeroInit())
	e.AddCardFromPKG(ele)
	e.AddCardFromPKG(crazy)
	for _, cards := range e.CardsPkg {
		fmt.Println(cards)
	}
	fmt.Println("wtf", e.Bank.Saving)
	e.MakeMortgage([]uint{0, 1})
	fmt.Println("wtf", e.Bank.Saving)
	e.DayPast()
	fmt.Println("wtf", e.Bank.Saving)
	fmt.Println(e.Bank.Loans[0])
	fmt.Println(e.Bank.Saving)
	e.DayPast()
	fmt.Println(e.Bank.Loans[0])
	e.DayPast()
	fmt.Println(e.Bank.Loans[0])
	fmt.Println(e.Bank.Trust, e.Bank.Saving)
	e.Bank.Saving += 2000
	fmt.Println(e.Bank.Saving)
	fmt.Println(e.PayMortgage(0))
	for _, cards := range e.CardsPkg {
		fmt.Println(cards)
	}
	fmt.Println(e.Bank.Trust, e.Bank.Saving)
	//fmt.Println(e.CardsPkg)
	fmt.Println((uint(0) * 30) / 100)
}
