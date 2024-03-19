package monitorcenter

import (
	"fmt"
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
		GID                uint                //游戏房间Id
		HeroMap            map[uint]*Hero      //根据hero id 查找英雄
		MonitorMap         map[uint]*Monitor   //根据monitor id索引monitor
		TypeMappingMonitor map[uint][]*Monitor //根据行动的id 查询monitor数组
	}

	Monitor struct {
		MID uint // monitor id 根据id鉴别monitor
		Tid uint
		MonitorLicense
		ListenerList   map[uint]*Hero // 对应Hero id的task map 多对一中的多
		Froze          bool           // 是否暂停中
		SonMonitor     []*Monitor
		BrotherMonitor []*Monitor
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
func (mc *MonitorCenter) ListenAndFilter(heroId, listenType uint) {
	for _, monitor := range mc.MonitorMap {

		if monitor.Type == listenType {
			fmt.Println(monitor.MID, heroId, monitor.Subject)
			if heroId == monitor.Owner.Id && monitor.Subject == monitorfile.OwnerMap("自己") {
				//监听自己的监听者
				mc.Publish(monitor)
				return
			}
			if monitor.Subject == monitorfile.OwnerMap("己方单位不包含自己") && mc.HeroMap[heroId].Owner == monitor.Owner.Owner {
				// 监听队友的监听者
				mc.Publish(monitor)
				return
			}
			if monitor.Subject == monitorfile.OwnerMap("己方单位不包含自己") && mc.HeroMap[heroId].Owner != monitor.Owner.Owner && mc.HeroMap[heroId].Owner != monitorfile.OwnerMap("中立") {
				// 监听对手的监听者
				mc.Publish(monitor)
				return
			}
			if monitor.Subject == monitorfile.OwnerMap("己方单位不包含自己") && mc.HeroMap[heroId].Owner != monitor.Owner.Owner && mc.HeroMap[heroId].Owner != monitorfile.OwnerMap("地方单位") {
				// 监听中立单位的监听者
				mc.Publish(monitor)
				return
			}
		}
	}
}
func (mc *MonitorCenter) DeleteSonMonitor(m *Monitor) {
	for _, monitor := range m.SonMonitor {
		mc.MonitorMap[monitor.Tid] = nil
	}
	m.SonMonitor = nil
}
func (mc *MonitorCenter) DeleteBrotherMonitor(m *Monitor) {
	for _, monitor := range m.BrotherMonitor {
		mc.MonitorMap[monitor.Tid] = nil
	}
	m.BrotherMonitor = nil
}
func (mc *MonitorCenter) TotalDeleteMonitor(m *Monitor) {
	mc.DeleteSonMonitor(m)
	mc.DeleteBrotherMonitor(m)
	mc.MonitorMap[m.Tid] = nil
}
func (mc *MonitorCenter) Publish(m *Monitor) {
	switch m.MID {
	case monitorfile.MonitorIdMap("狂战士之血"):
		m.Owner.AttackPoint++
	case monitorfile.MonitorIdMap("死亡"):
		//亡语...
		mc.TotalDeleteMonitor(m)

	case monitorfile.MonitorIdMap("放逐"):
		mc.TotalDeleteMonitor(m)
	}
}

func (m *Monitor) SonMonitorAdd(sons []*Monitor) {
	for _, son := range sons {
		m.SonMonitor = append(m.SonMonitor, son)
	}
}
func (m *Monitor) BrotherMonitorAdd(Brother []*Monitor) {
	for _, brothers := range Brother {
		m.BrotherMonitor = append(m.BrotherMonitor, brothers)
	}
}
