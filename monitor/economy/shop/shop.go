package shop

import (
	"train/monitor/hero"
)

type Cards struct {
	Hero  hero.Hero // 单位
	CD    int       // cd
	BuyCd int       // 购买cd 价格-（150 or 100）/100
}

// 购买cd
// 购买价格
type Shop struct {
	Cards map[int]*Cards
}

type BaseShop struct {
	Cards map[int]*Cards
}

func ShopInit() (*Shop, *BaseShop) {
	shop := Shop{Cards: map[int]*Cards{}}
	baseShop := BaseShop{Cards: map[int]*Cards{}}
	return &shop, &baseShop
}

// 替换卡牌or增加卡牌到shop中
func (b *BaseShop) AddToShop(cards []Cards) {
	for i := 0; i < len(cards); i++ {
		b.Cards[cards[i].Hero.Id] = &cards[i]
	}
}
func (b *BaseShop) HerosToShop(heros map[int]*hero.Hero) {
	for _, h := range heros {

		c := MakeCard(*h)
		c.BuyCd = (c.Hero.Price - 100) / 100
		c.CD = c.BuyCd
		b.Cards[h.Id] = c
	}
}
func (b *BaseShop) HerosToShopList(heros []*hero.Hero) {
	for _, h := range heros {

		c := MakeCard(*h)
		c.BuyCd = (c.Hero.Price - 100) / 100
		c.CD = c.BuyCd
		b.Cards[h.Id] = c
	}
}

func (b *BaseShop) ReMoveToShop(cards Cards) {
	b.Cards[cards.Hero.Id] = nil
}

// 选择卡牌 beta
func (b *BaseShop) ChoseHero(hero []hero.Hero) {
	if len(hero) == 0 {
		return
	}
	var cardList []Cards
	for i := 0; i < len(hero); i++ {
		card := MakeCard(hero[i])
		card.BuyCd = (card.Hero.Price - 100) / 100
		card.CD = card.BuyCd
		cardList = append(cardList, *card)
	}
	b.AddToShop(cardList)
}

func MakeCard(hero hero.Hero) *Cards {
	return &Cards{
		Hero:  hero,
		CD:    0,
		BuyCd: 0,
	}
}

func DayPast(shop *Shop, baseShop *BaseShop) {
	for u, _ := range shop.Cards {
		if shop.Cards[u].CD != 0 {
			shop.Cards[u].CD--
		}
	}
	for u, _ := range baseShop.Cards {
		if baseShop.Cards[u].CD != 0 {
			baseShop.Cards[u].CD--
		}
	}
}
