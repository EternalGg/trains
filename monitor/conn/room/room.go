package room

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	mc "train/monitor"
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
		GameState     int //游戏状态 1选择卡牌 2早上 3早上routine 4结束 5中午 6下午routine 7晚上 8夜晚routine
		T             int // 时间

	}
	Gamer struct {
		ID          int
		WaitingTime int //等待时间
	}
	ServerSession struct {
		Token []byte // token验证
		Type  int    // 0未登陆 1登陆player信息 2大厅等待信息 3卡牌选择信息 4游戏内时间与State
		// 5 end 6 input channel
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
			fmt.Println("当前卡牌"+strconv.Itoa(id), len(cc.ChoseCards), string(CardChose.Data), gs.GameData, gs)
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
	//fmt.Println(r.MonitorCenter.Economy[r.Gamer.ID].ShowHero())
	// card chose 选择结束 进入routine
	r.GameState = 2
	//fmt.Println("game state" + strconv.Itoa(r.GameState))
	// 自由模式 游戏结束根据 游戏外quit
	for r.GameState != 4 {
		switch r.GameState {
		case 2:
			//早上 同时接收部署和购买 当结束进入routine
			select {
			case LandingOrBuying := <-r.Ch:
				// buying 1查询商店 2购买 b+id 3 出售 s+id
				// landing 1查询地图 2部署 l+id+位置
				// ending 结束回合
				fmt.Println(LandingOrBuying)
			}
		case 3:
		case 5:
		case 6:

		}
	}

	return
	//r.MonitorCenter.
	// 2.游戏循环到十回合

}
