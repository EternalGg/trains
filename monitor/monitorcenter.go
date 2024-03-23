package monitorcenter

import (
	"strconv"
	"train/monitor/monitorfile"
)

type (
	Hero struct {
		Owner       uint // 0 GamerA 1 GamerB 2 中立单位
		Id          uint
		Health      uint
		Name        string
		AttackPoint uint
	}

	MonitorCenter struct {
		GID         uint              //游戏房间Id
		HeroMap     map[uint]*Hero    //根据hero id 查找英雄
		MonitorMap  map[uint]*Monitor //根据monitor id索引monitor
		MonitorLogs []string          //记录monitor 所有发生的事情
	}

	Monitor struct {
		MID uint // monitor id 根据id鉴别monitor
		Tid uint
		MonitorLicense
		ListenerList map[uint]*Hero // 对应Hero id的task map 多对一中的多
		Froze        bool           // 是否暂停中
		Logs         []string
	}
	MonitorLicense struct {
		Type    uint  // 订阅条件type 当canChange后 登场....
		Owner   *Hero // hero
		Subject uint  // 可订阅的英雄类型：属于谁的规则 0 全部 1 友军 2 中立 3 敌人 ...
	}
	Listener interface {
		Filter()
	}
)

// MonitorCenterInit MonitorCenter Start
func MonitorCenterInit(gid uint) *MonitorCenter {
	return &MonitorCenter{
		GID:        gid,
		MonitorMap: map[uint]*Monitor{},
		HeroMap:    map[uint]*Hero{},
	}
}

func (mc *MonitorCenter) AddMonitor(monitor *Monitor) {
	// 将monitor增加到mc中
	for i := 0; i < len(mc.MonitorMap)+1; i++ {
		if mc.MonitorMap[uint(i)] == nil {
			mc.MonitorMap[uint(i)] = monitor
			monitor.Tid = uint(i)
			break
		}
	}
}
func (mc *MonitorCenter) AddMonitorList(monitor []*Monitor) {
	// 将monitor列表增加到mc中
	for _, m := range monitor {
		mc.AddMonitor(m)
	}
}
func (mc *MonitorCenter) AddHeroInHeroMap(hero *Hero) {
	//增加hero到mcHeroMap
	mc.HeroMap[hero.Id] = hero
}

func (mc *MonitorCenter) ListenAndFilter(heroId, listenType uint, targets interface{}) []string {
	result := make([]string, 0)
	last := len(mc.MonitorLogs)
	for _, monitor := range mc.MonitorMap {
		if monitor.Type == listenType {
			if monitor.Subject == monitorfile.OwnerMap("自己") && heroId == monitor.Owner.Id {
				//监听自己的监听者
				mc.Publish(monitor, targets)
			}
			if monitor.Subject == monitorfile.OwnerMap("己方单位不包含自己") && mc.HeroMap[heroId].Owner == monitor.Owner.Owner && heroId != monitor.Owner.Id {
				// 监听队友的监听者
				mc.Publish(monitor, targets)
			}
			if monitor.Subject == monitorfile.OwnerMap("敌方单位") && mc.HeroMap[heroId].Owner != monitor.Owner.Owner && mc.HeroMap[heroId].Owner != 2 {
				// 监听对手的监听者
				mc.Publish(monitor, targets)
			}
			if monitor.Subject == monitorfile.OwnerMap("中立单位") && mc.HeroMap[heroId].Owner == 2 {
				// 监听中立单位的监听者
				mc.Publish(monitor, targets)
			}
		}
	}
	result = mc.MonitorLogs[last:]
	return result
}

func (mc *MonitorCenter) TotalDeleteMonitor(m *Monitor) {
	//删除相关所有monitor
	for id, monitor := range mc.MonitorMap {
		if monitor.Owner.Id == m.Owner.Id {
			mc.MonitorMap[id] = nil
		}
	}
}
func (mc *MonitorCenter) Publish(m *Monitor, targets interface{}) {
	switch m.MID {
	case monitorfile.MonitorIdMap("狂战士之血"):
		mc.MonitorLogs = append(mc.MonitorLogs, "之前攻击力", strconv.Itoa(int(m.Owner.AttackPoint)), "触发了狂战士之血，狂战士攻击+1", "之后攻击力为", strconv.Itoa(int(m.Owner.AttackPoint+1)))
		//time
		m.Logs = append(m.Logs, "之前攻击力", strconv.Itoa(int(m.Owner.AttackPoint)), "触发了狂战士之血，狂战士攻击+1", "之后攻击力为", strconv.Itoa(int(m.Owner.AttackPoint+1)))
		m.Owner.AttackPoint++
	case monitorfile.MonitorIdMap("死亡"):
		//亡语...
		mc.TotalDeleteMonitor(m)

	case monitorfile.MonitorIdMap("放逐"):
		mc.TotalDeleteMonitor(m)
	}
}
