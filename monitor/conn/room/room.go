package room

import (
	"encoding/json"
	"strconv"
	"time"
	mc "train/monitor"
	"train/monitor/battlefiled"
	"train/monitor/conn/room/notice"
	"train/monitor/economy/shop"
	"train/monitor/hero"
	"train/monitor/hero/attribute"
	"train/monitor/hero/heros"
	"train/monitor/hero/soldiers"
	"train/monitor/monitorfile"
	monitors2 "train/monitor/monitors"
	"train/monitor/routines"
	"train/monitor/routines/attack"
	skills2 "train/monitor/skills"
)

// inner game series
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
		GameState     *BaseState //游戏状态 0结束  1选择卡牌 2早上 3早上routine 4中午 5下午routine 6晚上 7夜晚routine
		T             int        // 时间
		CC            *CardChose
		HT            *HeroTurn
		HeroStr       string
		SkillsStr     string
	}
	Gamer struct {
		ID          int
		WaitingTime int //等待时间
	}
	ServerSession struct {
		Token string // token验证
		Type  int    // 0未登陆 1登陆player信息 2大厅等待信息 3卡牌选择信息 4游戏内时间与State
		// 5 购买商店信息 6 出售物品信息
		Data string // when Type=1
	}
	ClientSession struct {
		Token string // token验证
		Type  int    // 0未登陆 1登陆player信息 2
		// 5 end 6 input channel
		Data string // when Type=1
	}
	BaseState struct {
		Version   int    //信息版本
		Time      int    //时间
		GameState string //游戏时间状态
	}
	// 从客户端发来的session
	GameSession struct {
		GameDatatype int //1 选择卡牌
		GameData     string
	}
	CardChose struct {
		RemainMoney int //剩余金钱
		CardPool    []*hero.Hero
		CardPoolStr []string
		ChoseCount  int
	}
	PKG struct {
		CardsPkg map[int]*shop.Cards // 卡牌
		Money    int                 // 金钱
	}
	Skill struct {
		Skill        *skills2.Skill // skill
		Selectable   *Selectable    // 技能可以选中的目标
		Available    bool           // 是否可用
		NotAvailable []string       // 不可用id 1 资源不足 2 cd不足 3 范围内无可用目标
	}
	Selectable struct {
		Unit        []*hero.Hero // 可选择的单位
		Pos         []int        // 可选择的地图id
		SelectPoint int          // 可选择长度
	}
	HeroTurn struct {
		Skills []*Skill
		Hero   *hero.Hero
	}
)

type (
	SkillStr struct {
		Skill        string // skill struct
		Selectable   string // cut SelectableStr list
		Available    string // 是否可用
		NotAvailable string // cut string list t or f
	}
	SelectableStr struct {
		Unit   string // cut SelectableStr id  list
		Pos    string // cut Pos id  list
		Length string
	}
	PKGStr struct {
		Money  string
		CardId []string
		TId    []string
	}
)
type (
	Landing struct {
		CardId   int //卡牌数据
		Position int //部署位置
	}
	Shop struct {
		IsSale bool //是否售卖
		CardId int  //卡牌id
	}
	Move struct {
		SkillId  int //技能序列
		TargetId int //目标序列 TargetId = SelectList[i] + 1
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
		GameState:     &BaseState{},
	}
	r.Ch = make(chan *ClientSession, 10)
	r.ChOut = make(chan *ServerSession, 10)
	r.Gamer.ID = 1
	r.GameState.GameState = "选择卡牌"
	r.T = 0
	return &r
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
			ns := BaseState{}
			ns.Time = r.GameState.Time
			ns.GameState = r.GameState.GameState
			ns.Version = r.GameState.Version
			jss, _ := json.Marshal(ns)
			s.Data = string(jss)
			r.ChOut <- &s
		}
	}()
	cc := CardChose{
		RemainMoney: r.MonitorCenter.Economy[r.Gamer.ID].Money,
		CardPool:    heros.SelectAllHeroByList(),
		CardPoolStr: HeroListToStrList(heros.SelectAllHeroByList()),
	}

	//card chose turn

	r.CC = &cc
	for cc.ChoseCount != 4 {
		select {
		case CardChose := <-r.Ch:
			// 如果为single chose
			gs := GameSession{}
			json.Unmarshal([]byte(CardChose.Data), &gs)
			ParseInt, _ := strconv.ParseInt(string(gs.GameData), 10, 32)
			id := int(ParseInt)
			// 如果卡牌选择为空（未选择）
			if cc.CardPool[id] != nil {
				if !cc.CardPool[id].AreadyChose {
					cc.RemainMoney -= cc.CardPool[id].Price
					cc.CardPool[id].AreadyChose = true
					cc.CardPoolStr = HeroListToStrList(cc.CardPool)
					r.CC = &cc
					r.GameState.Version++
					cc.ChoseCount++

					s := ServerSession{}
					s.Type = 3
					jss, _ := json.Marshal(cc)
					s.Data = string(jss)
					r.ChOut <- &s
				} else {
					//已经选择
				}

			} else {
			}
			//s, gs := ServerSession{}, GameSession{}
			//s.Type = 3
			//gs.GameDatatype = 1
			//jgs, _ := json.Marshal(gs)
			//s.Data = string(jgs)
			//r.ChOut <- &s
		}
	}
	r.MonitorCenter.Economy[r.Gamer.ID].ChoseBefore(cc.CardPool)
	r.MonitorCenter.Economy[r.Gamer.ID].Money = cc.RemainMoney

	// card chose 选择结束 进入routine
	r.GameState.GameState = "早上"
	// robot
	// r.LandingCrazyRobot()
	// resource shop bank map -> public search
	for r.GameState.GameState != "结束" {
		if r.GameState.GameState == "早上" || r.GameState.GameState == "中午" || r.GameState.GameState == "傍晚" {
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
				json.Unmarshal([]byte(LandingOrBuying.Data), &gs)
				switch gs.GameDatatype {
				case 2:
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 2
					//bs := r.MonitorCenter.Economy[r.Gamer.ID].BaseShop
					bsstrs := make([]shop.BaseShopStr, 0)
					for _, v := range r.MonitorCenter.Economy[r.Gamer.ID].BaseShop.Cards {
						bsstr := shop.BaseShopStr{
							CardId: "",
							Price:  "",
							CD:     "",
							BuyCd:  "",
						}
						bsstr.Price = strconv.Itoa(v.Hero.Price)
						bsstr.CardId = strconv.Itoa(v.Hero.Id)
						bsstr.BuyCd = strconv.Itoa(v.BuyCd)
						bsstr.CD = strconv.Itoa(v.CD)
						bsstrs = append(bsstrs, bsstr)
					}
					jbs, _ := json.Marshal(bsstrs)
					gs.GameData = string(jbs)
					jgs, _ := json.Marshal(gs)
					s.Data = string(jgs)
					r.ChOut <- &s

					// 2 shop
				case 3:
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 3
					//m := r.MonitorCenter.BattleFiled
					result := make([]battlefiled.PositionStr, 0)
					for _, v := range r.MonitorCenter.BattleFiled.Positions {
						s := battlefiled.PositionStr{
							Id:      "",
							HeroId:  "",
							Machine: "",
						}
						s.Id = strconv.Itoa(v.Id)
						if v.Hero != nil {
							s.HeroId = strconv.Itoa(v.Hero.Tid)
						}
						if v.Machine != nil {
							s.Machine = strconv.Itoa(v.Machine.Tid)
						}
						result = append(result, s)
					}
					jm, _ := json.Marshal(result)
					gs.GameData = string(jm)
					jgs, _ := json.Marshal(gs)
					s.Data = string(jgs)
					r.ChOut <- &s
					// 4 resource
					// 3 map
				case 4:
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 4
					pkg := PKG{
						r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg,
						r.MonitorCenter.Economy[r.Gamer.ID].Money,
					}

					pkgStr := PKGStr{
						Money:  strconv.Itoa(pkg.Money),
						CardId: make([]string, 0),
						TId:    make([]string, 0),
					}
					for k, v := range pkg.CardsPkg {
						pkgStr.TId = append(pkgStr.TId, strconv.Itoa(k))
						pkgStr.CardId = append(pkgStr.CardId, strconv.Itoa(v.Hero.Id))
					}

					jpkg, _ := json.Marshal(pkgStr)
					gs.GameData = string(jpkg)
					jgs, _ := json.Marshal(gs)
					s.Data = string(jgs)
					r.ChOut <- &s
					// 4 resource
				case 5:
					//结束该回合
					r.GameState.GameState = NextGameState(r.GameState.GameState)
					te := notice.TurnEndResultMade(true, 0)
					r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, te)

					n := notice.ActionToNotice([]notice.ActionData{*te}, "结束回合", 10)
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 5
					jn, _ := json.Marshal(n)
					gs.GameData = string(jn)
					jgs, _ := json.Marshal(gs)
					s.Data = string(jgs)
					r.ChOut <- &s

					// 5 end turn
				case 6:
				// normal banking data
				// 1.存钱
				// 2.取钱
				// 3.贷款
				// 4.还贷
				// 6 bank
				case 7:
				// wild shop  time
				case 8:
					//bluff time

				case 11:
					// 登陆
					land := Landing{}
					json.Unmarshal([]byte(gs.GameData), &land)
					//fmt.Println(land.Position, land.CardId)
					if r.MonitorCenter.BattleFiled.Positions[land.Position] != nil && r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId] != nil {
						// 如果没有这个bf
						if r.MonitorCenter.BattleFiled.Positions[land.Position].Hero == nil {
							// position
							// 放入英雄map以及时间map
							ad := r.CardLanding(land)
							r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, ad)
							r.GameState.Version++
							n := notice.ActionToNotice([]notice.ActionData{*ad}, "登场", 4)
							s, gs := ServerSession{}, GameSession{}
							s.Type = 3
							gs.GameDatatype = 11
							jn, _ := json.Marshal(n)
							gs.GameData = string(jn)
							jgs, _ := json.Marshal(gs)
							s.Data = string(jgs)
							r.ChOut <- &s
						}
					}

				case 12:
					shop := Shop{}
					json.Unmarshal([]byte(gs.GameData), &shop)
					if shop.IsSale {
						sale := r.MonitorCenter.Economy[r.Gamer.ID].SaleHero(shop.CardId)
						r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, sale)
						n := notice.ActionToNotice([]notice.ActionData{*sale}, "出售卡牌", 3)
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 12
						jsale, _ := json.Marshal(n)
						gs.GameData = string(jsale)
						jgs, _ := json.Marshal(gs)
						s.Data = string(jgs)
						r.ChOut <- &s
					} else {
						buy := r.MonitorCenter.Economy[r.Gamer.ID].BuyHero(shop.CardId)
						r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, buy)

						n := notice.ActionToNotice([]notice.ActionData{*buy}, "购买卡牌", 3)
						r.GameState.Version++
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 12
						jbuy, _ := json.Marshal(n)
						gs.GameData = string(jbuy)
						jgs, _ := json.Marshal(gs)
						s.Data = string(jgs)
						r.ChOut <- &s
					}
				case 101:
					// catch ht hero str
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 101
					m := r.HeroStr
					jm, _ := json.Marshal(m)
					gs.GameData = string(jm)
					jgs, _ := json.Marshal(gs)
					s.Data = string(jgs)
					r.ChOut <- &s
				case 102:
					// catch ht hero str
					s, gs := ServerSession{}, GameSession{}
					s.Type = 3
					gs.GameDatatype = 102
					m := r.SkillsStr
					jm, _ := json.Marshal(m)
					gs.GameData = string(jm)
					jgs, _ := json.Marshal(gs)
					s.Data = string(jgs)
					r.ChOut <- &s
				}
			}
		} else {
			for r.GameState.GameState == "早上routine" || r.GameState.GameState == "中午routine" || r.GameState.GameState == "傍晚routine" {
				r.GameRoutine()
			}
			r.GameState.GameState = NextGameState(r.GameState.GameState)
		}
		// next function

	}
	return
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

// skill find selected
func (t *TestingGame) FindSelected(skill *skills2.Skill, remainMoney int, h *hero.Hero) *Skill {
	selected := Selectable{
		Unit:        []*hero.Hero{},
		Pos:         []int{},
		SelectPoint: 0,
	}
	result := Skill{
		Skill:        skill,
		Selectable:   &selected,
		Available:    true,
		NotAvailable: []string{},
	}
	// 返回所有距离内的pos
	pos := []*battlefiled.Position{}
	for i := 0; i < 45; i++ {
		if t.MonitorCenter.BattleFiled.Positions[i] != nil {
			//fmt.Println(skill.Owner.Pos)
			//fmt.Println(t.MonitorCenter.BattleFiled.Positions[skill.Owner.Pos].Distance[i])
			if t.MonitorCenter.BattleFiled.Positions[skill.Owner.Pos].Distance[i] <= skill.Distance {
				pos = append(pos, t.MonitorCenter.BattleFiled.Positions[i])
			}
		}
	}
	// 筛选
	for _, target := range skill.Targets {
		//0 队友 1敌方单位 2中立单位 3自己 4 无单位的地形 5 无物品的地形 6 有单位的地形 7 有物品的地形
		switch target {
		case 0:
			for _, p := range pos {
				if p.Hero != nil {
					if p.Hero.Owner == skill.Owner.Owner && p.Hero != skill.Owner {
						selected.Unit = append(selected.Unit, p.Hero)
					}
				}
				if p.Machine != nil {
					if p.Machine.Owner == skill.Owner.Owner && p.Machine != skill.Owner {
						selected.Unit = append(selected.Unit, p.Machine)
					}
				}
			}
		case 1:
			for _, p := range pos {
				if p.Hero != nil {
					if p.Hero.Owner != skill.Owner.Owner && p.Hero.Owner != 2 {
						selected.Unit = append(selected.Unit, p.Hero)
					}
				}
				if p.Machine != nil {
					if p.Machine.Owner != skill.Owner.Owner && p.Machine.Owner != 2 {
						selected.Unit = append(selected.Unit, p.Machine)
					}
				}
			}
		case 2:
			for _, p := range pos {
				if p.Hero != nil {
					if p.Hero.Owner == 2 {
						selected.Unit = append(selected.Unit, p.Hero)
					}
				}
				if p.Machine != nil {
					if p.Machine.Owner == 2 {
						selected.Unit = append(selected.Unit, p.Machine)
					}
				}
			}
		case 3:
			// 自己
			selected.Unit = append(selected.Unit, skill.Owner)
		case 4:
			for _, p := range pos {
				if p.Hero == nil {
					selected.Pos = append(selected.Pos, p.Id)
				}
			}
		case 5:
			for _, p := range pos {
				if p.Machine == nil {
					selected.Pos = append(selected.Pos, p.Id)
				}
			}
		case 6:
			for _, p := range pos {
				if p.Hero != nil {
					selected.Pos = append(selected.Pos, p.Id)
				}
			}
		case 7:
			for _, p := range pos {
				if p.Machine != nil {
					selected.Pos = append(selected.Pos, p.Id)
				}
			}
		}
	}
	result.Selectable.SelectPoint = len(selected.Unit) + len(selected.Pos)
	// 判断技能目标是否足够，不足而且不能无目标释放 A = false
	if selected.SelectPoint == 0 {
		if !skill.NoTarget {
			result.Available = false
			result.NotAvailable = append(result.NotAvailable, "3")
		}
	}
	// 判断英雄剩余金钱是否足够
	if result.Skill.Money > remainMoney || result.Skill.MovePoint > h.TActionPoint {
		result.Available = false
		result.NotAvailable = append(result.NotAvailable, "1")
	}
	// 判断英雄剩余金钱是否足够
	//fmt.Println(result.Selectable, result.NotAvailable, result.Available)
	return &result
}

// skill
func (r *TestingGame) SkillsDynamicUpdates(h *hero.Hero) *HeroTurn {
	ht := HeroTurn{}
	// skill id select
	skillsIdList := h.PositiveSkills
	//ht.RemainMoney = r.MonitorCenter.Economy[r.Gamer.ID].Money
	ht.Hero = h
	var skills []*Skill
	for _, i2 := range skillsIdList {
		skill := skills2.StrToSkills(monitorfile.SkillsIntToStrMap(i2))
		skill.Owner = ht.Hero
		// skill selected unit/position calculate
		s := r.FindSelected(skill, r.MonitorCenter.Economy[r.Gamer.ID].Money, h)
		//fmt.Print(s.Skill, s.Selectable.SelectPoint)
		//for _, po := range s.Selectable.Pos {
		//	fmt.Println(po)
		//}
		skills = append(skills, s)
	}

	// hero turn
	ht.Skills = skills

	return &ht
}

// success := 0 成功扣减
// errcode := 1 没有该技能！
// errcode := 2 资源不足或者没有目标！
// errcode := 3 没有该目标！
func (r *TestingGame) SkillJudgeAndReduce(h *HeroTurn, targetId, skillId int) (int, *Skill) {
	// 判断技能是否存在
	var skill *Skill

	for _, skills := range h.Skills {
		if skillId == skills.Skill.Id {
			skill = skills
			break
		}
	}
	if skill == nil {
		return 1, nil
	}

	if !skill.Available {
		return 2, nil
	}
	// 判断目标是否存在
	if skill.Selectable.SelectPoint < targetId {
		return 3, nil
	}
	// 判断成功 扣除资源

	r.MonitorCenter.Economy[r.Gamer.ID].Money -= skill.Skill.Money
	switch skill.Skill.Id {
	case 7:
		// 自然疗法
		h.Hero.TActionPoint -= mc.RandomNumber(3)
	default:
		h.Hero.TActionPoint -= skill.Skill.MovePoint
	}

	return 0, skill
}

func (r *TestingGame) LogToNotice(lens int) *notice.Notice {
	if lens == len(r.MonitorCenter.MonitorLogs) {
		return nil
	} else {
		result := notice.Notice{
			NoticeName: r.MonitorCenter.MonitorLogs[lens].ActionName,
			Id:         r.MonitorCenter.MonitorLogs[lens].ID,
			Actions:    []notice.ActionData{},
		}

		for i := lens; i < len(r.MonitorCenter.MonitorLogs); i++ {
			result.Actions = append(result.Actions, *r.MonitorCenter.MonitorLogs[i])
		}
		return &result
	}

}

func (s *Skill) FindTarget(p int) (*hero.Hero, int) {

	if s.Selectable.Unit != nil {

		for i := 0; i < len(s.Selectable.Unit); i++ {
			p--
			if p == 0 {
				return s.Selectable.Unit[i], 0
			}
		}

	}
	if s.Selectable.Pos != nil {
		for i := 0; i < len(s.Selectable.Pos); i++ {
			p--
			if p == 0 {
				return nil, s.Selectable.Pos[i]
			}
		}
	}
	return nil, 0
}

// 卡牌登陆
func (r TestingGame) CardLanding(land Landing) *notice.ActionData {
	r.MonitorCenter.PutHeroInHeroMap(&r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero)
	r.MonitorCenter.PutHeroInTimeMap(&r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero)

	// 放入地区
	r.MonitorCenter.BattleFiled.Positions[land.Position].Hero = &r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero
	// 英雄的pos更改
	r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero.Pos = land.Position
	// 卡牌中的卡牌out
	// time map

	// monitor init
	heros.Landing(&r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg[land.CardId].Hero, r.MonitorCenter)
	// 删除卡牌map中的卡牌
	delete(r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg, land.CardId)
	ad := notice.LandingResultMade(true, "登场成功", r.MonitorCenter.BattleFiled.Positions[land.Position].Hero, 0, land.Position)
	r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, ad)
	return ad
}

// 登场狂爆机器人到3和20
func (r *TestingGame) LandingCrazyRobot() {
	crazy1 := soldiers.SoldiersHeroInit()
	crazy1.Name = "稻草人"
	heros.Landing(crazy1, r.MonitorCenter)
	crazy1.Pos = 20
	r.MonitorCenter.BattleFiled.Positions[20].Hero = crazy1
	r.MonitorCenter.PutHeroInHeroMap(crazy1)
	crazy1.Owner = 1

	crazy2 := soldiers.SoldiersHeroInit()
	crazy2.Name = "稻草人"
	heros.Landing(crazy2, r.MonitorCenter)
	crazy2.Pos = 3
	r.MonitorCenter.BattleFiled.Positions[3].Hero = crazy2
	r.MonitorCenter.PutHeroInHeroMap(crazy2)
	crazy2.Owner = 1
}

// 根据英雄周围距离找到所有英雄
func (r TestingGame) FindHerosByDistance(h *hero.Hero, distance int) []*hero.Hero {
	result := []*hero.Hero{}
	for i := 0; i < 45; i++ {
		if r.MonitorCenter.BattleFiled.Positions[i] != nil {
			//fmt.Println(skill.Owner.Pos)
			//fmt.Println(t.MonitorCenter.BattleFiled.Positions[skill.Owner.Pos].Distance[i])
			if r.MonitorCenter.BattleFiled.Positions[h.Pos].Distance[i] <= distance && r.MonitorCenter.BattleFiled.Positions[i].Hero != nil {
				result = append(result, r.MonitorCenter.BattleFiled.Positions[i].Hero)
			}
		}
	}
	return result
}

func (r TestingGame) FindHerosByDistanceWithOutSelf(h *hero.Hero, distance int) []*hero.Hero {
	result := []*hero.Hero{}
	for i := 0; i < 45; i++ {
		if r.MonitorCenter.BattleFiled.Positions[i] != nil {
			//fmt.Println(skill.Owner.Pos)
			//fmt.Println(t.MonitorCenter.BattleFiled.Positions[skill.Owner.Pos].Distance[i])
			//判断是否是自己
			if r.MonitorCenter.BattleFiled.Positions[h.Pos].Distance[i] <= distance && r.MonitorCenter.BattleFiled.Positions[i].Hero != nil {
				if r.MonitorCenter.BattleFiled.Positions[i].Hero == h {
					continue
				}
				result = append(result, r.MonitorCenter.BattleFiled.Positions[i].Hero)
			}
		}
	}
	return result
}

// 游戏routinue
func (r *TestingGame) GameRoutine() {
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
			if nowHero == nil {
				break
			}
			nowHero.TActionPoint = nowHero.ActionPoint
			HT := r.SkillsDynamicUpdates(nowHero)
			r.HT = HT
			r.GameState.GameState = nowHero.Name + "行动回合！"
			r.HeroStr = HeroStrInit(r.HT)
			r.SkillsStr = HeroSkillStrInit(r.HT)
			// hero turn send to client
			for !end {
				select {
				// 判断skill的Available 如果==true
				// 如果存在 查询目标是否在selected中
				// 如果可以释放 两个判断结束 资源扣除后释放技能
				case heroTurn := <-r.Ch:
					//fmt.Println(heroTurn.Type, heroTurn.Data)
					gs := GameSession{}
					json.Unmarshal([]byte(heroTurn.Data), &gs)
					switch gs.GameDatatype {
					case 2:
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 2
						//bs := r.MonitorCenter.Economy[r.Gamer.ID].BaseShop
						bsstrs := make([]shop.BaseShopStr, 0)
						for _, v := range r.MonitorCenter.Economy[r.Gamer.ID].BaseShop.Cards {
							bsstr := shop.BaseShopStr{
								CardId: "",
								Price:  "",
								CD:     "",
								BuyCd:  "",
							}
							bsstr.Price = strconv.Itoa(v.Hero.Price)
							bsstr.CardId = strconv.Itoa(v.Hero.Id)
							bsstr.BuyCd = strconv.Itoa(v.BuyCd)
							bsstr.CD = strconv.Itoa(v.CD)
							bsstrs = append(bsstrs, bsstr)
						}
						jbs, _ := json.Marshal(bsstrs)
						gs.GameData = string(jbs)
						jgs, _ := json.Marshal(gs)
						s.Data = string(jgs)
						r.ChOut <- &s

						// 2 shop
					case 3:
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 3
						//m := r.MonitorCenter.BattleFiled
						result := make([]battlefiled.PositionStr, 0)
						for _, v := range r.MonitorCenter.BattleFiled.Positions {
							s := battlefiled.PositionStr{
								Id:      "",
								HeroId:  "",
								Machine: "",
							}
							s.Id = strconv.Itoa(v.Id)
							if v.Hero != nil {
								s.HeroId = strconv.Itoa(v.Hero.Tid)
							}
							if v.Machine != nil {
								s.Machine = strconv.Itoa(v.Machine.Tid)
							}
							result = append(result, s)
						}
						jm, _ := json.Marshal(result)
						gs.GameData = string(jm)
						jgs, _ := json.Marshal(gs)
						s.Data = string(jgs)
						r.ChOut <- &s
						// 4 resource
						// 3 map
					case 4:
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 4
						pkg := PKG{
							r.MonitorCenter.Economy[r.Gamer.ID].CardsPkg,
							r.MonitorCenter.Economy[r.Gamer.ID].Money,
						}

						pkgStr := PKGStr{
							Money:  strconv.Itoa(pkg.Money),
							CardId: make([]string, 0),
							TId:    make([]string, 0),
						}
						for _, v := range pkg.CardsPkg {
							pkgStr.TId = append(pkgStr.TId, strconv.Itoa(v.Hero.Tid))
							pkgStr.CardId = append(pkgStr.CardId, strconv.Itoa(v.Hero.Id))
						}

						jpkg, _ := json.Marshal(pkgStr)
						gs.GameData = string(jpkg)
						jgs, _ := json.Marshal(gs)
						s.Data = string(jgs)
						r.ChOut <- &s
					// 4 resource
					// 4 resource

					case 10:

						move := Move{}
						json.Unmarshal([]byte(gs.GameData), &move)
						no := &notice.Notice{}
						var state int
						var skill *Skill

						state, skill = r.SkillJudgeAndReduce(HT, move.TargetId, move.SkillId)
						monitors := []*monitors2.Monitor{}
						for _, i2 := range r.MonitorCenter.MonitorMap {
							monitors = append(monitors, i2)
						}
						if state == 0 {
							var h *hero.Hero
							var pos int
							if !skill.Skill.NoTarget {
								h, pos = skill.FindTarget(move.TargetId)
							}
							switch state {
							case 1:
							case 2:
							case 3:
							default:
								switch skill.Skill.Id {
								case 1:
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, notice.ActionStart())
									l := len(r.MonitorCenter.MonitorLogs)
									attack := attack.Attack{
										Name:     "Attack",
										Attacker: nowHero,
										Targets:  h,
										Damage:   0,
										Mc:       r.MonitorCenter,
									}
									routines.Gates(attack, r.MonitorCenter)
									no = r.LogToNotice(l)
									HT = r.SkillsDynamicUpdates(nowHero)
									r.HT = HT
								case 2:
									// 移动
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, notice.ActionStart())
									move := r.MonitorCenter.BattleFiled.HeroMove(
										r.MonitorCenter.BattleFiled.Positions[nowHero.Pos],
										r.MonitorCenter.BattleFiled.Positions[pos],
										nowHero)
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, move)
									l := len(r.MonitorCenter.MonitorLogs)
									no = r.LogToNotice(l)
									HT = r.SkillsDynamicUpdates(nowHero)
									r.HT = HT
								case 3:
								case 4:
									// 早上回合结束
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, notice.ActionStart())

									n := notice.TurnEndResultMade(true, 0)
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, n)
									r.MonitorCenter.TurnEndListen(nowHero)
									// 检测turnend里需要end的monitor

									l := len(r.MonitorCenter.MonitorLogs)
									no = r.LogToNotice(l)
									end = true
								case 5:
									// 野兽冲击

									// 对周围的所有单位造成1点攻击
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, notice.ActionStart())
									l := len(r.MonitorCenter.MonitorLogs)

									// 对周围的所有单位造成1点攻击
									hs := r.FindHerosByDistanceWithOutSelf(nowHero, 1)
									gh := &monitors2.Monitor{}
									for _, monitor := range monitors {
										if monitor.MID == monitorfile.MonitorIdMap("巨型食草动物") {
											gh = monitor
										}
									}
									for _, h := range hs {

										a := attack.Attack{
											Name:     "移动攻击",
											Attacker: nowHero,
											Targets:  h,
											Damage:   1,
											Mc:       r.MonitorCenter,
										}
										if gh.LifeTime == 1 {
											a.Damage = 2
										}
										routines.Gates(a, r.MonitorCenter)
									}
									// move
									move := r.MonitorCenter.BattleFiled.HeroMove(
										r.MonitorCenter.BattleFiled.Positions[nowHero.Pos],
										r.MonitorCenter.BattleFiled.Positions[pos],
										nowHero)
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, move)
									// 对周围的所有单位造成1点攻击
									hs2 := r.FindHerosByDistanceWithOutSelf(nowHero, 1)
									for _, h := range hs2 {
										a := attack.Attack{
											Name:     "移动攻击",
											Attacker: nowHero,
											Targets:  h,
											Damage:   1,
											Mc:       r.MonitorCenter,
										}
										if gh.LifeTime == 1 {
											a.Damage = 2
										}
										routines.Gates(a, r.MonitorCenter)
									}

									no = r.LogToNotice(l)
									HT = r.SkillsDynamicUpdates(nowHero)
									r.HT = HT
								case 6:
									// 撞击
									// move
									// 随机攻击附近一个敌方单位
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, notice.ActionStart())
									l := len(r.MonitorCenter.MonitorLogs)
									// move
									move := r.MonitorCenter.BattleFiled.HeroMove(
										r.MonitorCenter.BattleFiled.Positions[nowHero.Pos],
										r.MonitorCenter.BattleFiled.Positions[pos],
										nowHero)
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, move)
									hs := r.FindHerosByDistanceWithOutSelf(nowHero, 1)
									//attack
									if len(hs) != 0 {
										t := mc.RandomNumber(len(hs))
										attack := attack.Attack{
											Name:     "巨兽冲击",
											Attacker: nowHero,
											Targets:  hs[t],
											Damage:   nowHero.AttackPoint + (nowHero.AttackPoint / 2),
											Mc:       r.MonitorCenter,
										}
										routines.Gates(attack, r.MonitorCenter)
									}

									no = r.LogToNotice(l)
									HT = r.SkillsDynamicUpdates(nowHero)
									r.HT = HT
								case 7:
									r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, notice.ActionStart())
									l := len(r.MonitorCenter.MonitorLogs)
									// healing
									randomHealing := mc.RandomNumber(9)
									healingMap := map[int]int{}
									healingMap[monitorfile.BubbleIdMap("加血")] = randomHealing
									attr := attribute.Attribute{
										Hero:    h,
										AttrMap: healingMap,
									}
									attr.Publish()
									no = r.LogToNotice(l)
									HT = r.SkillsDynamicUpdates(nowHero)
									r.HT = HT
								}

								// htstr update
								r.HeroStr = HeroStrInit(r.HT)
								r.SkillsStr = HeroSkillStrInit(r.HT)
								r.GameState.Version++
							}
						}
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 7
						jsale, _ := json.Marshal(no)
						gs.GameData = string(jsale)
						jgs, _ := json.Marshal(gs)
						s.Data = string(jgs)
						r.ChOut <- &s

					case 101:
						// catch ht hero str
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 101
						m := r.HeroStr
						jm, _ := json.Marshal(m)
						gs.GameData = string(jm)
						jgs, _ := json.Marshal(gs)
						s.Data = string(jgs)
						r.ChOut <- &s
					case 102:
						// catch ht hero str
						s, gs := ServerSession{}, GameSession{}
						s.Type = 3
						gs.GameDatatype = 102
						m := r.SkillsStr
						jm, _ := json.Marshal(m)
						gs.GameData = string(jm)
						jgs, _ := json.Marshal(gs)
						s.Data = string(jgs)
						r.ChOut <- &s
					}
				}
			}
			// 回归 action point
		}
		r.MonitorCenter.Time.Actions = r.MonitorCenter.Time.Actions[1:]
	}
	// round past check
	no := &notice.Notice{}
	r.MonitorCenter.MonitorLogs = append(r.MonitorCenter.MonitorLogs, notice.ActionStart())
	l := len(r.MonitorCenter.MonitorLogs)
	r.MonitorCenter.RoundPast()
	no = r.LogToNotice(l)
	if no != nil {
		s, gs := ServerSession{}, GameSession{}
		s.Type = 3
		gs.GameDatatype = 7
		jsale, _ := json.Marshal(no)
		gs.GameData = string(jsale)
		jgs, _ := json.Marshal(gs)
		s.Data = string(jgs)
		r.ChOut <- &s
	}
}

func HeroListToStrList(heros []*hero.Hero) []string {
	result := make([]string, 0)
	for i := 0; i < len(heros); i++ {
		herostr, _ := json.Marshal(heros[i])
		result = append(result, string(herostr))
	}
	return result
}

func HeroStrInit(ht *HeroTurn) string {

	// hero
	heroStr, _ := json.Marshal(ht.Hero)
	return string(heroStr)
}

func HeroSkillStrInit(ht *HeroTurn) string {
	// skills
	skillstrlist := make([]string, 0)
	for i := 0; i < len(ht.Skills); i++ {
		ss := SkillStr{
			Skill:        "",
			Selectable:   "",
			Available:    "",
			NotAvailable: "",
		}
		// 1 not available
		ss.NotAvailable = ListToStrList(ht.Skills[i].NotAvailable)
		// 2 available
		if ht.Skills[i].Available {
			ss.Available = "t"
		} else {
			ss.Available = "f"
		}
		// 3 skill
		skilljson, _ := json.Marshal(ht.Skills[i].Skill)
		ss.Skill = string(skilljson)
		// 4
		SAS := SelectableStr{
			Unit:   "",
			Pos:    "",
			Length: "",
		}
		// 4.1
		SAS.Length = strconv.Itoa(ht.Skills[i].Selectable.SelectPoint)
		// 4.2
		unitstrlist := make([]string, 0)
		for i := 0; i < len(ht.Skills[i].Selectable.Unit); i++ {
			unitstrlist = append(unitstrlist, strconv.Itoa(ht.Skills[i].Selectable.Unit[i].Tid))
		}
		SAS.Unit = ListToStrList(unitstrlist)
		// 4.3
		posstrlist := make([]string, 0)
		for i := 0; i < len(ht.Skills[i].Selectable.Pos); i++ {
			posstrlist = append(posstrlist, strconv.Itoa(ht.Skills[i].Selectable.Pos[i]))
		}
		SAS.Pos = ListToStrList(posstrlist)
		// json
		SASJSon, _ := json.Marshal(SAS)
		ss.Selectable = string(SASJSon)
		// end
		ssjson, _ := json.Marshal(ss)
		skillstrlist = append(skillstrlist, string(ssjson))
		//htstr.SkillsStr = append(htstr.SkillsStr, string(ssjson))
	}

	result := ListToStrList(skillstrlist)
	return result
}

func StrToStrList(str string) []string {
	result := make([]string, 0)
	p := ""
	for i := 0; i < len(str); i++ {
		if str[i] != '-' {
			p += string(str[i])
		} else {
			result = append(result, p)
			p = ""
		}
	}
	return result
}

func ListToStrList(list []string) string {
	result := ""
	for i := 0; i < len(list); i++ {
		result += list[i]
		if i != len(list)-1 {
			result += "-"
		}
	}
	return result
}
