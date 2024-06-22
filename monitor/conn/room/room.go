package room

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	mc "train/monitor"
	"train/monitor/economy/shop"
	"train/monitor/hero"
	"train/monitor/hero/heros"
)

type (
	NormalGame struct {
		MonitorCenter *mc.MonitorCenter
		Gamer1        *Gamer
		Gamer2        *Gamer
	}
	TestingGame struct {
		MonitorCenter *mc.MonitorCenter
		Gamer         *Gamer
		Ch            chan *ClientSession
		ChOut         chan *ServerSession
		GameState     int //游戏状态 0结束  1选择卡牌 2早上 3早上routine 4中午 5下午routine 6晚上 7夜晚routine
		T             int // 时间

	}
	Gamer struct {
		ID          int
		WaitingTime int //等待时间
	}
	ServerSession struct {
		Token []byte // token验证
		Type  int    // 0未登陆 1登陆player信息 2大厅等待信息 3卡牌选择信息 4游戏内时间与State
		// 5 购买商店信息 6 出售物品信息
		Data []byte // when Type=1
	}
	ClientSession struct {
		Token []byte // token验证
		Type  int    // 0未登陆 1登陆player信息 2大厅等待信息 3卡牌选择信息 4游戏内时间与State
		// 5 end 6 input channel
		Data []byte // when Type=1
	}

	State struct {
		Time      int //时间
		GameState int //游戏时间状态
	}
	GameSession struct {
		GameDatatype int //1 选择卡牌
		GameData     []byte
	}
	CardChose struct {
		SessionState int                //1卡牌阅览 2err已经被选择 3success选择卡牌成功
		RemainMoney  int                //剩余金钱
		CardPool     map[int]*hero.Hero //可选择卡牌池
		ChoseCards   map[int]*hero.Hero //已选择卡牌池
	}
	Landing struct {
		CardId   int //卡牌数据
		Position int //部署位置
	}
	Shop struct {
		IsSale bool //是否售卖
		CardId int  //卡牌id
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
	PKG struct {
		CardsPkg map[int]*shop.Cards // 卡牌
		Money    int                 // 金钱
	}
)

func GamerInit() *Gamer {
	return &Gamer{
		WaitingTime: 300,
	}
}

func TestingGameInit(userid int) *TestingGame {
	r := TestingGame{
		MonitorCenter: mc.MonitorCenterInit(userid),
		Gamer:         GamerInit(),
	}
	r.Ch = make(chan *ClientSession, 10)
	r.ChOut = make(chan *ServerSession, 10)
	r.Gamer.ID = 1
	r.GameState = 1
	r.T = 0
	return &r
}
func (r *NormalGame) GameStart() {
	// todo
	// 游戏倒计时，到达正确的时间内
	// 正确的时间内监控正确的事情
	// 当时间超过时 扣除金钱
	// 监视游戏胜利
}
func (r *TestingGame) GameStart() {
	// todo
	// 游戏倒计时，到达正确的时间内
	// 正确的时间内监控正确的事情
	// 当时间超过时 扣除金钱
	// 监视游戏胜利
	// 返回正确的信息
	// 计时以及
	r.GameState = 1
	go func() {
		for {
			r.T++
			time.Sleep(1000000000)
			s, ss := ServerSession{}, State{}
			s.Type = 4
			ss.Time = r.T
			ss.GameState = r.GameState

			jss, _ := json.Marshal(ss)
			s.Data = jss
			r.ChOut <- &s
		}
	}()
	cc := CardChose{
		RemainMoney: r.MonitorCenter.Economy[r.Gamer.ID].Money,
		CardPool:    heros.SelectAllHeroByMap(),
		ChoseCards:  map[int]*hero.Hero{},
	}

	//每十秒发送卡牌信息给客户端
	go func() {
		for r.GameState == 1 {
			time.Sleep(10000000000)
			s, gs := ServerSession{}, GameSession{}
			s.Type = 3
			gs.GameDatatype = 1
			cc.SessionState = 1
			jc, _ := json.Marshal(cc)
			gs.GameData = jc
			jgs, _ := json.Marshal(gs)
			s.Data = jgs
			r.ChOut <- &s
		}
	}()
	//card chose turn
	for len(cc.ChoseCards) != 7 {
		select {
		case CardChose := <-r.Ch:
			// 如果为single chose
			gs := GameSession{}
			json.Unmarshal(CardChose.Data, &gs)
			ParseInt, _ := strconv.ParseInt(string(gs.GameData), 10, 32)
			id := int(ParseInt)
			// 如果为选择卡牌

			// 如果卡牌选择为空（未选择）
			if cc.CardPool[id] != nil {
				cc.RemainMoney -= cc.CardPool[id].Price
				cc.ChoseCards[id] = cc.CardPool[id]
				delete(cc.CardPool, id)
				cc.SessionState = 3
			} else {
				cc.SessionState = 2
			}
			s, gs := ServerSession{}, GameSession{}
			s.Type = 3
			gs.GameDatatype = 1
			jc, _ := json.Marshal(cc)
			gs.GameData = jc
			jgs, _ := json.Marshal(gs)
			s.Data = jgs
			r.ChOut <- &s
		}

	}
	r.MonitorCenter.Economy[r.Gamer.ID].ChoseBefore(cc.ChoseCards)
	r.MonitorCenter.Economy[r.Gamer.ID].Money = cc.RemainMoney
	//fmt.Println(r.MonitorCenter.Economy[r.Gamer.ID].ShowHero())
	// card chose 选择结束 进入routine
	r.GameState = 2
	//fmt.Println("game state" + strconv.Itoa(r.GameState))
	// 自由模式 游戏结束根据 游戏外quit
	for r.GameState != 0 {
		switch r.GameState {
		case 2:
			//早上 同时接收部署和购买 当结束进入routine
			select {
			case LandingOrBuying := <-r.Ch:
				// buy  buy+id
				// shop 查询商店
				// sale sale+id
				// land land+id+位置
				// map 查询地图
				// end 结束回合
				// pkg 查询金钱和卡包
				gs := GameSession{}
				json.Unmarshal(LandingOrBuying.Data, &gs)
				switch gs.GameDatatype {
				case 2:
					if string(gs.GameData) == "shop" {
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 2
						bs := r.MonitorCenter.Economy[r.Gamer.ID].BaseShop
						jbs, _ := json.Marshal(bs)
						gs.GameData = jbs
						jgs, _ := json.Marshal(gs)
						s.Data = jgs
						r.ChOut <- &s
					} else {
						shop := Shop{}
						json.Unmarshal(gs.GameData, &shop)
						if shop.IsSale {
							r.MonitorCenter.Economy[r.Gamer.ID].SaleHero(shop.CardId)
						} else {
							r.MonitorCenter.Economy[r.Gamer.ID].BuyHero(shop.CardId)
						}
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 5
						pkg := PKG{
							r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg,
							r.MonitorCenter.Economy[r.Gamer.ID].Money,
						}
						jpkg, _ := json.Marshal(pkg)
						gs.GameData = jpkg
						jgs, _ := json.Marshal(gs)
						s.Data = jgs
						r.ChOut <- &s
					}
				case 3:
					// 地图以及登陆
					if string(gs.GameData) != ("map") {
						land := Landing{}
						json.Unmarshal(gs.GameData, &land)
						if r.MonitorCenter.BattleFiled.Positions[land.Position] != nil && r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId] != nil {
							// 如果没有这个bf
							if r.MonitorCenter.BattleFiled.Positions[land.Position].Hero == nil {
								r.MonitorCenter.BattleFiled.Positions[land.Position].Hero = &r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero
								r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId] = nil
								heros.LandingById(r.MonitorCenter.BattleFiled.Positions[land.Position].Hero.Id, r.MonitorCenter)
							}
						}
						for _, monitor := range r.MonitorCenter.MonitorMap {
							fmt.Println(monitor)
						}
					}
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 3
					m := r.MonitorCenter.BattleFiled
					jm, _ := json.Marshal(m)
					gs.GameData = jm
					jgs, _ := json.Marshal(gs)
					s.Data = jgs
					r.ChOut <- &s
				case 4:
					//结束该回合
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 5
					r.GameState = NextGameState(r.GameState)
				case 5:
					// 查询卡牌和经济
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 5
					pkg := PKG{
						r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg,
						r.MonitorCenter.Economy[r.Gamer.ID].Money,
					}
					jpkg, _ := json.Marshal(pkg)
					gs.GameData = jpkg
					jgs, _ := json.Marshal(gs)
					s.Data = jgs
					r.ChOut <- &s
				}
			}
		case 3:

		case 4:
		case 5:
		case 6:

		}
	}

	return
	//r.MonitorCenter.
	// 2.游戏循环到十回合

}

func NextGameState(i int) int {
	if i != 7 {
		i++
	} else {
		i = 2
	}
	return i
}
