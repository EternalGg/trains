package monitorcenter

import (
	"math/rand"
	"sort"
	"strconv"
	"time"
	"train/monitor/battlefiled"
	"train/monitor/economy"
	"train/monitor/hero"
	"train/monitor/hero/attribute"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

type (
	MonitorCenter struct {
		GID          int                       //游戏房间Id
		HeroMap      map[int]*hero.Hero        //根据hero id 查找英雄
		MonitorMap   map[int]*monitors.Monitor //根据monitor id索引monitor
		MonitorCount map[int]int               //monitor类别的记录
		MonitorLogs  []Logs                    //记录monitor 所有发生的事情
		BattleFiled  *battlefiled.BattleFiled
		Time         MonitorTime
		Economy      map[int]*economy.Economy // map[userid]*economy
	}
	Logs struct {
		MainEvent interface{}
		SubEvent  []interface{}
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
	AttackCalculate struct {
		AttackerName       string
		TargetName         string
		Name               string
		Id                 uint
		BaseDamage         int                       // 基础攻击力
		DamageAddition     int                       // 攻击加成
		CriticalHitRate    int                       // 暴击概率
		CriticalStrikeRate int                       // 暴击倍率加成 0为100%暴击率（无暴击）
		OtherDamage        int                       // 固定伤害加成
		IsCritical         bool                      // 是否暴击
		FinalDamage        int                       // 最终伤害
		Sessions           []monitors.MonitorSummary // 信息
		ErrorSession       []int                     // 错误信息

	}
	BeAttackCalculate struct {
		Name             string
		Id               int
		DogeRate         int                       // 闪避率
		DamageReduce     int                       // 伤害减免
		DamageReduceRate int                       // 伤害减免百分比
		IsDoge           bool                      // 是否闪避
		FinalDamage      int                       // 最终伤害
		Sessions         []monitors.MonitorSummary // 信息
		ErrorSession     []int                     // 错误信息
		FightBack        bool                      // 反击
		DamageDepthRate  int                       // 伤害加深率
		DamageDepth      int                       // 伤害加深数值
	}
	TakeDamage struct {
		Name     string
		Id       uint
		HitBack  uint                      // 反弹伤害
		Sessions []monitors.MonitorSummary // 信息
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
		Round:       0,
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
		a.HID = h.Id
		t.Time.Actions = append(t.Time.Actions, &a)
	}
	// 时间监听
	t.TimeListener(0)
	// 英雄速度归位
	for _, h := range heroList {
		t.Time.Time.HeroTime[h] = (h.Speed)
	}
	// time past
	t.Time.Time.Round++
	if t.Time.Time.Round%12 == 0 {
		t.TimeQuantumPast()
	}
}
func (t *MonitorCenter) TimeQuantumPast() {
	switch t.Time.Time.Round {
	case 24:
		//夜晚开始
		t.Time.Time.IsNight = false
		t.Time.Time.TimeQuantum++
	case 12:
		//中午开始
		t.Time.Time.TimeQuantum++
	case 48:
		t.DayPast()
		t.Time.Time.IsNight = true
		//早晨开始
		t.Time.Time.TimeQuantum = 0
	}
	t.TimeListener(2)
}
func (t *MonitorCenter) DayPast() {
	t.Time.Time.Day++
	t.Time.Time.Round = 0
	t.TimeListener(2)
}
func (t *MonitorCenter) NightPast() {
	t.Time.Time.Day++
	t.Time.Time.Round = 0
	t.TimeListener(3)
}
func (t *MonitorCenter) PutHeroInHeroMap(hero *hero.Hero) {
	t.Time.Time.HeroTime[hero] = (hero.Speed)
	t.AddHeroInHeroMap(hero)
}

// time 类 结束

// Init MonitorCenter Start
func MonitorCenterInit(gid int) *MonitorCenter {
	result := &MonitorCenter{
		GID:          gid,
		HeroMap:      map[int]*hero.Hero{},
		MonitorMap:   map[int]*monitors.Monitor{},
		MonitorLogs:  make([]Logs, 0),
		MonitorCount: map[int]int{},
		Economy:      map[int]*economy.Economy{},
	}
	result.Time = MonitorTime{
		Time:           TimeInit(),
		TimeChange:     map[*monitors.Monitor]uint{},
		DayChange:      map[*monitors.Monitor]uint{},
		NQuantumChange: map[*monitors.Monitor]uint{},
		NightChange:    map[*monitors.Monitor]bool{},
	}
	result.BattleFiled = battlefiled.NormalGame()

	result.Economy[gid] = economy.EconomyInit()
	return result
}

// monitor的增加和删除
func (mc *MonitorCenter) AddMonitor(monitor *monitors.Monitor) {
	// 将monitor增加到mc中
	for i := 0; i < len(mc.MonitorMap)+1; i++ {
		if mc.MonitorMap[int(i)] == nil {
			mc.MonitorMap[int(i)] = monitor
			mc.MonitorCount[monitor.MID]++
			monitor.Tid = (i)
			switch monitor.MID {
			case monitorfile.MonitorIdMap("夜行动物"):
				//  转换
				mc.TimeListenSign(monitor)
			default:
				attr := attribute.Attribute{
					Hero:    monitor.Owner,
					AttrMap: monitor.Bubble,
				}
				attr.Positive()
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
			mc.MonitorCount[monitor.MID]--
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
				if license.Subject == monitorfile.OwnerMap("敌方单位") && mc.HeroMap[heroId].Owner != monitor.Owner.Owner && mc.HeroMap[heroId].Owner != 2 {
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
	log := Logs{
		MainEvent: m,
		SubEvent:  make([]interface{}, 0),
	}
	switch m.MID {
	default:
		// 默认增加属性
		attr := attribute.Attribute{
			Hero:    m.Owner,
			AttrMap: m.PublishState,
		}
		attr.Positive()
		for i, i2 := range m.PublishState {
			str := m.Owner.Name + monitorfile.BubbleIdMapToStr(i) + "增加了" + strconv.Itoa(i2)
			log.SubEvent = append(log.SubEvent, str)
		}
	}
	mc.MonitorLogs = append(mc.MonitorLogs, log)
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
	}
}
func (mc *MonitorCenter) TimeListener(t uint) {
	switch t {
	case 1:
		// 度过一天

	case 2:
		// 度过时间区间

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
					attr.Positive()
				} else {
					attr.Negative()
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

// time end

// bf start
func (mc *MonitorCenter) BattleFiledChange() {

}
func (mc *MonitorCenter) BattleFiledListen() {

}
func (mc *MonitorCenter) BattleFiledSign() {

}
func (mc *MonitorCenter) HeroLanding(hid int, bfid int) {

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
