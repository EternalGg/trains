package shop

import (
	"fmt"
	"testing"
	"train/monitor/hero"
	crazy2 "train/monitor/hero/crazy"
	ele2 "train/monitor/hero/ele"
	fly2 "train/monitor/hero/fly"
	"train/monitor/hero/soldiers"
)

func TestDevi(t *testing.T) {
	fmt.Println(100 / 100)
}

func TestBaseShop_ChoseHero(t *testing.T) {
	ele := ele2.EleHeroInit()
	cra := crazy2.CrazyHeroInit()
	fly := fly2.FlyHeroInit()
	so := soldiers.SoldiersHeroInit()
	shop, baseshop := ShopInit()
	cards := []hero.Hero{*ele, *cra, *fly, *so}
	baseshop.ChoseHero(cards)
	DayPast(shop, baseshop)
	for _, c := range baseshop.Cards {
		fmt.Println(c.CD, c.BuyCd, c.Hero.Price, c)
	}
}

func TestBuyFromShop(t *testing.T) {

}
