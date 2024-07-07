package monitorcenter

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"train/monitor/battlefiled"
	"train/monitor/conn/room/notice"
	"train/monitor/economy"
	"train/monitor/hero"
	"train/monitor/hero/attribute"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

type (
	MonitorCenter struct {
		GID            int                //游戏房间Id
		HeroMap        map[int]*hero.Hero //根据hero id 查找英雄
		HeroMonitorMap map[*hero.Hero]map[*monitors.Monitor]bool
		MonitorMap     map[int]*monitors.Monitor //根据monitor id索引monitor
		MonitorCount   map[int]int               //monitor类别的记录
		MonitorLogs    []*notice.ActionData      //记录monitor 所有发生的事情
		BattleFiled    *battlefiled.BattleFiled
		Time           MonitorTime
		Economy        map[int]*economy.Economy // map[userid]*economy
	}

	Worker interface {
		Checker()
		Calculator()
		Processer()
	}
	MonitorTime struct {
		Time           *Time                      //时间类
		TimeChange     map[*monitors.Monitor]uint //倒计时类
		DayChange      map[*monitors.Monitor]uint //天数
		NQuantumChange map[*monitors.Monitor]uint //早上0 下午1 晚上2
		NightChange    map[*monitors.Monitor]bool //白天黑夜 白天True 黑夜False
		HeroTurnEnd    map[*monitors.Monitor]bool //英雄回合结束类
		Actions        []*Action                  //行动
	}
)

// session class
type (
	PreAttackMonitor struct {
		Changes []monitors.MonitorSummary //攻击前会发生什么
	}
	BeforeAttackMonitor struct {
		Changes []monitors.MonitorSummary //攻击前会发生什么
	}
	SessionsAttack struct {
		Name string
		PreAttackMonitor
		BeforeAttackMonitor
	}

	MakeDamage struct {
		FinalDamage uint                      // 最终伤害
		Sessions    []monitors.MonitorSummary // 信息
	}
	Time struct {
		Day         uint               // 日期
		Round       uint               // 回合 一天共48个
		TimeQuantum uint               // 时间 早上，下午，晚上 0,1,2
		HeroTime    map[*hero.Hero]int // map[hero]speed 根据heroid映射速度
		IsNight     bool               // 是夜晚吗
	}
	// rountine中的Action们 包含英雄，单位，建筑的行动 定时的被动技能
	Action struct {
		HeroAction bool // 是否是英雄行动
		HID        int  // 应该行动的hero id
		MID        int  // monitor id
	}
)

// time 类 开始
func TimeInit() *Time {
	return &Time{
		Day:         0,
		Round:       1,
		TimeQuantum: 0,
		HeroTime:    map[*hero.Hero]int{},
		IsNight:     true,
	}
}

func (t *MonitorCenter) RoundPast() {
	// hero round
	heroList := []*hero.Hero{}

	for h, _ := range t.Time.Time.HeroTime {
		t.Time.Time.HeroTime[h] -= 1
		if t.Time.Time.HeroTime[h] == 0 {
			heroList = append(heroList, h)
		}
	}
	if len(heroList) != 0 {
		sort.Slice(heroList, func(i, j int) bool {
			return heroList[i].Speed < heroList[j].Speed
		})
	}
	for _, h := range heroList {
		a := Action{}
		a.HeroAction = true
		a.HID = h.Tid
		t.Time.Actions = append(t.Time.Actions, &a)
	}
	// 时间监听
	t.TimeListener(0)
	// 英雄速度归位
	for _, h := range heroList {
		t.Time.Time.HeroTime[h] = h.Speed
		fmt.Println(h.Name, h.Speed)
	}
	// time past
	t.Time.Time.Round++
	if t.Time.Time.Round%13 == 0 {
		t.TimeQuantumPast()
	}
}
func (t *MonitorCenter) TimeQuantumPast() {
	switch t.Time.Time.Round {
	case 26:
		//夜晚开始
		t.Time.Time.IsNight = true
		t.Time.Time.TimeQuantum++
		t.DayPast()
	case 13:
		//中午开始
		t.Time.Time.TimeQuantum++
	case 52:

		t.Time.Time.IsNight = false
		//早晨开始
		t.Time.Time.TimeQuantum = 0
		t.NightPast()
	}
}
func (t *MonitorCenter) DayPast() {
	t.TimeListener(2)
}
func (t *MonitorCenter) NightPast() {
	t.Time.Time.Day++
	t.Time.Time.Round = 0
	t.TimeListener(3)
}
func (t *MonitorCenter) PutHeroInHeroMap(hero *hero.Hero) {
	t.AddHeroInHeroMap(hero)
}
func (t *MonitorCenter) PutHeroInTimeMap(hero *hero.Hero) {
	t.Time.Time.HeroTime[hero] = (hero.Speed)
}

// time 类 结束

// Init MonitorCenter Start
func MonitorCenterInit(gid int) *MonitorCenter {
	result := &MonitorCenter{
		GID:            gid,
		HeroMap:        map[int]*hero.Hero{},
		MonitorMap:     map[int]*monitors.Monitor{},
		MonitorLogs:    make([]*notice.ActionData, 0),
		HeroMonitorMap: map[*hero.Hero]map[*monitors.Monitor]bool{},
		MonitorCount:   map[int]int{},
		Economy:        map[int]*economy.Economy{},
	}
	result.Time = MonitorTime{
		Time:           TimeInit(),
		TimeChange:     map[*monitors.Monitor]uint{},
		DayChange:      map[*monitors.Monitor]uint{},
		NQuantumChange: map[*monitors.Monitor]uint{},
		NightChange:    map[*monitors.Monitor]bool{},
		HeroTurnEnd:    map[*monitors.Monitor]bool{},
	}
	result.BattleFiled = battlefiled.NormalGame()

	result.Economy[gid] = economy.EconomyInit()
	return result
}

// monitor的增加和删除
func (mc *MonitorCenter) MonitorCountChange(key, value int) {
	mc.MonitorCount[key] += value
}
func (mc *MonitorCenter) AddMonitor(monitor *monitors.Monitor) {
	// 将monitor增加到mc中
	for i := 0; i < len(mc.MonitorMap)+1; i++ {
		if mc.MonitorMap[int(i)] == nil {
			mc.MonitorMap[int(i)] = monitor
			mc.MonitorCountChange(monitor.MID, 1)
			mc.HeroMonitorMap[monitor.Owner][monitor] = true
			monitor.Tid = (i)
			switch monitor.MID {
			case monitorfile.MonitorIdMap("夜行动物"):
				//  转换
				mc.TimeListenSign(monitor)
			case monitorfile.MonitorIdMap("巨型食草动物"):
				mc.TimeListenSign(monitor)
			default:
				attr := attribute.Attribute{
					Hero:    monitor.Owner,
					AttrMap: monitor.Bubble,
				}
				attr.Publish()
			}
			break
		}
	}
}
func (mc *MonitorCenter) DeleteMonitor(m *monitors.Monitor) {
	//删除相关所有monitor
	for id, monitor := range mc.MonitorMap {
		if monitor.Owner.Tid == m.Owner.Id && monitor.RelianceOwner {
			mc.MonitorMap[id] = nil
			mc.MonitorCountChange(monitor.MID, -1)
		}
	}
}
func (mc *MonitorCenter) AddMonitorList(monitor []*monitors.Monitor) {
	// 将monitor列表增加到mc中
	for _, m := range monitor {
		mc.AddMonitor(m)
	}
}
func (mc *MonitorCenter) AddHeroInHeroMap(hero *hero.Hero) {
	//增加hero到mcHeroMap
	hero.Tid = len(mc.HeroMap)
	mc.HeroMap[hero.Tid] = hero
}

// listen
func (mc *MonitorCenter) ListenAndFilter(heroId, listenType int) (result []*monitors.Monitor) {
	for _, monitor := range mc.MonitorMap {
		// ListenLicense
		for _, license := range monitor.ListenLicense {
			if license.ListenType == listenType {
				if license.Subject == monitorfile.OwnerMap("自己") && heroId == monitor.Owner.Tid {
					//监听自己的监听者
					result = append(result, monitor)
				}
				if license.Subject == monitorfile.OwnerMap("己方单位不包含自己") && mc.HeroMap[heroId].Owner == monitor.Owner.Owner && heroId != monitor.Owner.Tid {
					// 监听队友的监听者
					result = append(result, monitor)
				}
				if license.Subject == monitorfile.OwnerMap("敌方单位") && mc.HeroMap[heroId].Owner != monitor.Owner.Owner {
					// 监听对手的监听者
					result = append(result, monitor)
				}
				if license.Subject == monitorfile.OwnerMap("中立单位") && mc.HeroMap[heroId].Owner == 2 {
					// 监听中立单位的监听者
					result = append(result, monitor)
				}
			}
		}
	}
	return
}
func (mc *MonitorCenter) MonitorPublish(m *monitors.Monitor) {

	switch m.MID {
	case 27:
		if m.LifeTime == 0 {
			m.LifeTime = 1
			m.LifeTimeState = 2
			attr := attribute.Attribute{
				Hero:    m.Owner,
				AttrMap: m.Bubble,
			}
			attr.Publish()
			mc.MonitorLogs = append(mc.MonitorLogs, notice.BubbleResultMade(m.Name, m.Owner, true, attr.AttrMap))
		}
	default:
		// 默认增加属性
		attr := attribute.Attribute{
			Hero:    m.Owner,
			AttrMap: m.PublishState,
		}
		// 增加到bubble中
		for key, value := range m.PublishState {
			m.Bubble[key] += value
		}
		attr.Publish()
		// mc log data
		mc.MonitorLogs = append(mc.MonitorLogs, notice.BubbleResultMade(m.Name, m.Owner, true, attr.AttrMap))
	}
	return
}

// 无序版本
func (mc *MonitorCenter) MonitorsPublish(ms []*monitors.Monitor) {
	if len(ms) == 0 {
		return
	}
	for _, m := range ms {
		mc.MonitorPublish(m)
	}
}
func (mc *MonitorCenter) MonitorsActive(ms []*monitors.Monitor) {
	if len(ms) == 0 {
		return
	}
	for _, m := range ms {
		mc.AddMonitor(m)
	}

}
func (mc *MonitorCenter) MonitorsDead(ms []*monitors.Monitor) {

}

// monitor in time
// 根据monitor id注册timelisten
// 比如倒计时类的技能，持续时间等
func (mc *MonitorCenter) TimeListenSign(m *monitors.Monitor) {
	switch m.MID {
	case monitorfile.MonitorIdMap("夜行动物"):
		mc.Time.NightChange[m] = false
	case monitorfile.MonitorIdMap("巨型食草动物"):
		mc.Time.HeroTurnEnd[m] = true
	}
}
func (mc *MonitorCenter) TimeListener(t uint) {
	switch t {
	case 1:
		// 度过一天

	case 2:
		// 度过时间区间
		for monitor, u := range mc.Time.NightChange {
			switch monitor.MID {
			case monitorfile.MonitorIdMap("夜行动物"):
				attr := attribute.Attribute{
					Hero:    monitor.Owner,
					AttrMap: monitor.Bubble,
				}
				if u == mc.Time.Time.IsNight {
					attr.ToNegative()
					attr.Publish()
					b := notice.BubbleResultMade("夜行动物取消", attr.Hero, true, attr.AttrMap)
					//fmt.Println("夜行动物,end!")
					mc.MonitorLogs = append(mc.MonitorLogs, b)
				} else {
					attr.Publish()
					b := notice.BubbleResultMade("夜行动物启动", attr.Hero, true, attr.AttrMap)
					//fmt.Println("夜行动物,start!")
					mc.MonitorLogs = append(mc.MonitorLogs, b)
				}
			}
		}
	case 3:
		// 白天true 黑天false
		for monitor, u := range mc.Time.NightChange {
			switch monitor.MID {
			case monitorfile.MonitorIdMap("夜行动物"):
				attr := attribute.Attribute{
					Hero:    monitor.Owner,
					AttrMap: monitor.Bubble,
				}
				if u == mc.Time.Time.IsNight {
					// 白天
					attr.ToNegative()
					attr.Publish()
					b := notice.BubbleResultMade("夜行动物取消", attr.Hero, true, attr.AttrMap)

					mc.MonitorLogs = append(mc.MonitorLogs, b)
				} else {
					attr.Publish()
					b := notice.BubbleResultMade("夜行动物启动", attr.Hero, true, attr.AttrMap)
					mc.MonitorLogs = append(mc.MonitorLogs, b)
				}
			}
		}
	default:
		for monitor, _ := range mc.Time.TimeChange {
			switch monitor.MID {

			}
		}
	}
}
func (mc *MonitorCenter) TurnEndListen(h *hero.Hero) {
	for monitor, b := range mc.Time.HeroTurnEnd {
		if monitor.Owner == h {
			fmt.Println(b)
			switch monitor.MID {
			case monitorfile.MonitorIdMap("巨型食草动物"):
				if monitor.LifeTime == 1 {
					attr := attribute.Attribute{
						Hero:    monitor.Owner,
						AttrMap: monitor.Bubble,
					}
					attr.ToNegative()
					attr.Publish()
					monitor.LifeTime = 0
					fmt.Println("ele 巨型食草动物结束")
					mc.MonitorLogs = append(mc.MonitorLogs, notice.BubbleResultMade("巨型食草动物结束!", h, true, attr.AttrMap))
				}
			}
		}
	}
}

// time end

// bf start
func (mc *MonitorCenter) BattleFiledChange() {

}
func (mc *MonitorCenter) BattleFiledListen() {

}
func (mc *MonitorCenter) BattleFiledSign() {

}

// bf end

// RandomByTime 输入概率百分比，返回概率结果
func RandomByTime(p int) bool {
	randomTime := time.Unix(int64(rand.Intn(1000000000)), 0)
	if int64(p) > ((randomTime.Unix()) % 100) {
		return true
	} else {
		return false
	}
}
