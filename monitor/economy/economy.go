package economy

import (
	"sort"
	"train/monitor/conn/room/notice"
	"train/monitor/economy/bank"
	"train/monitor/economy/shop"
	"train/monitor/hero"
	"train/monitor/monitorfile"
)

type (
	Economy struct {
		Money    int                 // 剩余金钱
		Bank     *bank.Bank          // 银行系统
		BaseShop *shop.BaseShop      // 基地商店
		Shop     *shop.Shop          // 商人
		CardsPkg map[int]*shop.Cards // 卡牌
	}
	SaleData struct {
		HeroName  string //英雄名称
		SaleMoney int    //销售价格
		ErrorCode int    //错误代码
		// errorcode 1:手牌中不存在该卡牌
	}
	BuyData struct {
		HeroName  string // 英雄名称
		BuyMoney  int    //购买价格
		ErrorCode int    //错误代码
		// errorcode 1:金钱不足
		// errorcode 2:购买cd
		// errorcode 3:商店不存在该hid
	}
)

func EconomyInit() *Economy {
	Shop, baseShop := shop.ShopInit()
	return &Economy{
		Money:    monitorfile.EconomyMapToint("初始金钱"),
		Bank:     bank.BankInit(),
		Shop:     Shop,
		BaseShop: baseShop,
		CardsPkg: map[int]*shop.Cards{},
	}
}

// shops
func (e *Economy) DeleteCardFromPKG(p int) {
	e.CardsPkg[p] = nil
}

func (e *Economy) AddCardFromPKG(card *shop.Cards) {
	for i := 0; i <= len(e.CardsPkg); i++ {
		if e.CardsPkg[(i)] == nil {
			e.CardsPkg[(i)] = card
			return
		}
	}
}

func (e *Economy) ChoseBefore(hero map[int]*hero.Hero) {
	for _, h := range hero {
		c := shop.MakeCard(*h)
		e.AddCardFromPKG(c)
	}
	e.BaseShop.HerosToShop(hero)
}

// shops
// 商品的购买
// errorcode 1:金钱不足
// errorcode 2:购买cd
// errorcode 3:不存在该hid
func (e *Economy) BuyHero(hid int) *notice.ActionData {
	buy := BuyData{}

	// 3
	if e.BaseShop.Cards[hid] == nil {
		buy.ErrorCode = 3
		return notice.BuyResultMade(false, "购买卡牌不存在该cardid", hid, e.Money, 0, 3)
	}
	// 1.查看剩余金钱是否大于售价
	if e.BaseShop.Cards[hid].Hero.Price > e.Money {
		buy.ErrorCode = 1
		return notice.BuyResultMade(false, "购买卡牌金钱不足", hid, e.Money, 0, 1)

	}
	// 2.查看购买cd
	if e.BaseShop.Cards[hid].BuyCd > 0 {
		buy.ErrorCode = 2
		return notice.BuyResultMade(false, "购买卡牌cd不足", hid, e.Money, 0, 2)
	}
	// 3.购买
	// 扣除钱
	e.Money -= e.BaseShop.Cards[hid].Hero.Price
	// 购买cd加
	e.BaseShop.Cards[hid].BuyCd += e.BaseShop.Cards[hid].CD
	// 卡牌进组
	e.AddCardFromPKG(e.BaseShop.Cards[hid])
	buy.BuyMoney = e.BaseShop.Cards[hid].Hero.Price
	buy.HeroName = e.BaseShop.Cards[hid].Hero.Name
	return notice.BuyResultMade(true, "卡牌购买成功！", hid, e.Money, buy.BuyMoney, 0)

}

// 卖
func (e *Economy) SaleHero(point int) *notice.ActionData {
	// error code 1:没有此卡牌
	sale := SaleData{}

	if e.CardsPkg[point] == nil {
		sale.ErrorCode = 1
		return notice.SaleResultMade(false, "手牌中没有该卡牌", 1, point, 0, e.Money)
	}
	// 钱到账
	saleMoney := (e.CardsPkg[(point)].Hero.Price * 6) / 10
	e.Money += saleMoney
	sale.SaleMoney = saleMoney
	sale.HeroName = e.CardsPkg[point].Hero.Name
	// 删除卡牌从卡组
	e.DeleteCardFromPKG((point))
	return notice.SaleResultMade(true, "出售成功！", 0, point, sale.SaleMoney, e.Money)
}

// banks
// Mortgage to bank 输入卡牌id
// error code
// 1:卡牌重样
// 2:卡组中没有该卡
// 3:贷款卡组没有超过额度1000
// 4:贷款额度不足
func (e *Economy) MakeMortgage(cards []int) int {
	// 1.判断卡牌是否重样
	sort.Slice(cards, func(i, j int) bool {
		return false
	})
	for i := 1; i < len(cards); i++ {
		if cards[i] == cards[i-1] {
			return 1
		}
	}
	// 2.判断卡牌是否存在
	for _, card := range cards {
		if e.CardsPkg[card] == nil {
			return 2
		}
	}
	// 3.判断额度
	count := 0
	for _, card := range cards {
		count += e.CardsPkg[card].Hero.Price
	}
	if count < 1000 {
		return 3
	}
	// 3.判断完成后，卡组卡牌删除
	list := make([]*shop.Cards, 0)
	for _, card := range cards {
		list = append(list, e.CardsPkg[card])
		e.DeleteCardFromPKG(card)
	}
	// 4.借贷失败
	if !e.Bank.MakeLoan(count, list) {
		return 4
	}
	// 5.返回成功code
	return 0
}

// 还贷
// 返回是否成功，成功就把抵押 单位 返回
func (e *Economy) PayMortgage(loanid int) bool {
	loan := e.Bank.Loans[loanid]
	cards := loan.MortgageList
	if e.Bank.Pay(loanid) {
		for _, card := range cards {
			e.AddCardFromPKG(card)
		}
		return true
	} else {
		return false
	}
}

// save money
// 1.判断存款
func (e *Economy) SaveToBank(money int) bool {
	if e.Money >= money {
		e.Bank.Save(money)
		e.Money -= money
		return true
	}
	return false
}

// 取钱
// 1.判断成功，成功money+=钱
func (e *Economy) TakeToBank(money int) bool {
	if e.Bank.Take(money) {
		e.Money += money
		return true
	}
	return false
}

func (e *Economy) DayPast() {
	e.Bank.DayPast()
}

func (e *Economy) ShowMoney() (int, int) {
	return e.Money, e.Bank.Saving
}

func (e *Economy) ShowHero() []shop.Cards {
	result := make([]shop.Cards, 0)
	for _, i2 := range e.CardsPkg {
		result = append(result, *i2)
	}
	return result
}

// BANKING 银行系统
// CREDIT 信用评级机制

// SHOP 商店系统 买以及卖
// REWARD 奖励系统
// UNIT COST 单位花费
// SKILL COST 技能花费
// BANKER/ECONOMY SKILL 银行家/经济技能
