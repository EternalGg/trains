package room

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	mc "train/monitor"
	"train/monitor/battlefiled"
	"train/monitor/economy/shop"
	"train/monitor/hero"
	"train/monitor/hero/heros"
	"train/monitor/monitorfile"
	skills2 "train/monitor/skills"
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
		GameState     *State //游戏状态 0结束  1选择卡牌 2早上 3早上routine 4中午 5下午routine 6晚上 7夜晚routine
		T             int    // 时间
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
		Time      int       //时间
		GameState string    //游戏时间状态
		DataState int       //数据类型 1 卡牌选择 2 英雄轮次
		HT        HeroTurn  //hero turn data
		CC        CardChose //chose card
	}
	GameSession struct {
		GameDatatype int //1 选择卡牌
		GameData     []byte
	}
	// s->c
	CardChose struct {
		RemainMoney int                //剩余金钱
		CardPool    map[int]*hero.Hero //可选择卡牌池
		ChoseCards  map[int]*hero.Hero //已选择卡牌池
	}
	CardChoseResult struct {
		SessionState string // 1err已经被选择 2success选择卡牌成功
	}
	// c->s
	Landing struct {
		CardId   int //卡牌数据
		Position int //部署位置
	}
	// c->s
	Shop struct {
		IsSale bool //是否售卖
		CardId int  //卡牌id
	}
	// s->c
	PKG struct {
		CardsPkg map[int]*shop.Cards // 卡牌
		Money    int                 // 金钱
	}
	Skill struct {
		Skill      *skills2.Skill // skill
		Selectable Selectable     // 技能可以选中的目标
	}
	Selectable struct {
		Unit []*hero.Hero            // 可选择的单位
		Pos  []*battlefiled.Position // 可选择的地图
	}
	// s->c
	HeroTurn struct {
		Skills      []*skills2.Skill
		RemainMoney int
		Hero        *hero.Hero
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
		GameState:     &State{},
	}
	r.Ch = make(chan *ClientSession, 10)
	r.ChOut = make(chan *ServerSession, 10)
	r.Gamer.ID = 1
	r.GameState.GameState = "选择卡牌"
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
	r.GameState.GameState = "选择卡牌"
	go func() {
		for {
			time.Sleep(1000000000)
			r.T++
			r.GameState.Time = r.T
			s := ServerSession{}
			s.Type = 4
			jss, _ := json.Marshal(r.GameState)
			s.Data = jss
			r.ChOut <- &s
		}
	}()
	cc := CardChose{
		RemainMoney: r.MonitorCenter.Economy[r.Gamer.ID].Money,
		CardPool:    heros.SelectAllHeroByMap(),
		ChoseCards:  map[int]*hero.Hero{},
	}

	//card chose turn
	r.GameState.DataState = 1
	r.GameState.CC = cc
	for len(cc.ChoseCards) != 3 {
		select {
		case CardChose := <-r.Ch:
			// 如果为single chose
			gs := GameSession{}
			json.Unmarshal(CardChose.Data, &gs)
			ParseInt, _ := strconv.ParseInt(string(gs.GameData), 10, 32)
			id := int(ParseInt)
			// 如果为选择卡牌
			result := CardChoseResult{}
			// 如果卡牌选择为空（未选择）
			if cc.CardPool[id] != nil {
				cc.RemainMoney -= cc.CardPool[id].Price
				cc.ChoseCards[id] = cc.CardPool[id]
				delete(cc.CardPool, id)
				result.SessionState = "卡牌选择成功!"
				r.GameState.CC = cc
			} else {
				result.SessionState = "Error已经被选择!"
			}
			s, gs := ServerSession{}, GameSession{}
			s.Type = 3
			gs.GameDatatype = 1
			resultJson, _ := json.Marshal(result)
			gs.GameData = resultJson
			jgs, _ := json.Marshal(gs)
			s.Data = jgs
			r.ChOut <- &s
		}

	}
	r.MonitorCenter.Economy[r.Gamer.ID].ChoseBefore(cc.ChoseCards)
	r.MonitorCenter.Economy[r.Gamer.ID].Money = cc.RemainMoney
	//fmt.Println(r.MonitorCenter.Economy[r.Gamer.ID].ShowHero())
	// card chose 选择结束 进入routine
	r.GameState.GameState = "早上"
	r.GameState.DataState = 0
	//fmt.Println("game state" + strconv.Itoa(r.GameState))
	// 自由模式 游戏结束根据 游戏外quit
	for r.GameState.GameState != "结束" {
		switch r.GameState.GameState {
		case "早上":
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
							sale := r.MonitorCenter.Economy[r.Gamer.ID].SaleHero(shop.CardId)
							s, gs := ServerSession{}, GameSession{}
							s.Type = 3
							gs.GameDatatype = 6
							jsale, _ := json.Marshal(sale)
							gs.GameData = jsale
							jgs, _ := json.Marshal(gs)
							s.Data = jgs
							r.ChOut <- &s
						} else {
							buy := r.MonitorCenter.Economy[r.Gamer.ID].BuyHero(shop.CardId)
							s, gs := ServerSession{}, GameSession{}
							s.Type = 3
							gs.GameDatatype = 4
							jbuy, _ := json.Marshal(buy)
							gs.GameData = jbuy
							jgs, _ := json.Marshal(gs)
							s.Data = jgs
							r.ChOut <- &s
						}
					}
				case 3:
					// 地图以及登陆
					land := Landing{}
					json.Unmarshal(gs.GameData, &land)
					//fmt.Println(land.Position, land.CardId)
					if string(gs.GameData) != ("map") {
						//for i, i2 := range r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg {
						//	fmt.Println(i, i2)
						//}
						//fmt.Println(r.MonitorCenter.BattleFiled.Positions[land.Position] != nil, r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId], r.MonitorCenter.BattleFiled.Positions[land.Position].Hero == nil, r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId], r.MonitorCenter.BattleFiled.Positions[land.Position].Hero == nil)
						if r.MonitorCenter.BattleFiled.Positions[land.Position] != nil && r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId] != nil {
							// 如果没有这个bf
							if r.MonitorCenter.BattleFiled.Positions[land.Position].Hero == nil {
								//fmt.Println(r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero)
								//fmt.Println(r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId])
								// position
								r.MonitorCenter.BattleFiled.Positions[land.Position].Hero = &r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero
								// 卡牌中的卡牌out
								// time map
								r.MonitorCenter.PutHeroInHeroMap(&r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero)
								// monitor init
								heros.LandingById(r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero.Id, r.MonitorCenter)
								delete(r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg, land.CardId)
							}
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
					r.GameState.GameState = NextGameState(r.GameState.GameState)
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
		case "早上routine":
			// 12次round past

			for i := 0; i < 12; i++ {
				for _, action := range r.MonitorCenter.Time.Actions {
					fmt.Print(action)
				}
				for len(r.MonitorCenter.Time.Actions) != 0 {
					// 速度排序
					// 判断事件是否结束 v1 object is dead? v2 target is dead?
					// 时序问题 monitor first ---- monitor speed?
					// hero secend ---- 充钱-speed random
					if !(r.MonitorCenter.Time.Actions[0].HeroAction) {
						// 定时被动
					} else {
						// 当前英雄结束
						end := false
						nowHero := r.MonitorCenter.HeroMap[r.MonitorCenter.Time.Actions[0].HID]

						ht := HeroTurn{}
						// skill id select
						skillsIdList := r.MonitorCenter.HeroMap[r.MonitorCenter.Time.Actions[0].HID].PositiveSkills
						skills := []*skills2.Skill{}
						for _, i2 := range skillsIdList {
							skill := skills2.StrToSkills(monitorfile.SkillsIntToStrMap(i2))
							// skill selected unit/position calculate

							skills = append(skills, skill)
						}
						// hero turn
						ht.Skills = skills
						ht.RemainMoney = r.MonitorCenter.Economy[r.Gamer.ID].Money
						ht.Hero = r.MonitorCenter.HeroMap[r.MonitorCenter.Time.Actions[0].HID]
						r.GameState.HT = ht
						r.GameState.GameState = nowHero.Name + "行动回合！"
						r.GameState.DataState = 2
						// hero turn send to client
						for !end {
							select {
							case heroTurn := <-r.Ch:

								gs := GameSession{}
								json.Unmarshal(heroTurn.Data, &gs)
								switch gs.GameDatatype {
								case 0:
									end = true
								default:

								}
							}
						}
					}
					r.MonitorCenter.Time.Actions = r.MonitorCenter.Time.Actions[1:]
				}
				r.MonitorCenter.RoundPast()
			}

		case "中午":
			// 中午
		case "下午routine":
			// 下午routine
		case "傍晚":
			// 晚上
		case "夜晚routine":
			// 晚上routine
		}
	}

	return
	//r.MonitorCenter.
	// 2.游戏循环到十回合

}

func NextGameState(str string) string {
	switch str {
	case "早上":
		return "早上routine"
	case "早上routine":
		return "中午"
	case "中午":
		return "下午routine"
	case "下午routine":
		return "傍晚"
	case "傍晚":
		return "夜晚routine"
	case "夜晚routine":
		return "早上"
	default:
		return ""
	}
}

// skill found selected

func (t *TestingGame) FindSelected(skill *skills2.Skill) *Skill {
	selected := Selectable{
		Unit: []*hero.Hero{},
		Pos:  []*battlefiled.Position{},
	}

	for _, target := range skill.Targets {
		switch target {
		case 0:
		case 1:
		case 2:
		case 3:
			// 自己
			selected.Unit = append(selected.Unit, skill.Owner)
		case 4:
			// 地形
			//selected.Pos
		}
	}
	result := Skill{
		Skill:      skill,
		Selectable: selected,
	}
	return &result
}
