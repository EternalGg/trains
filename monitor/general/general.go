package general

import (
	mc "train/monitor"
	"train/monitor/monitorfile"
)

// DeadMonitor 死亡monitor
func DeadMonitor(hero *mc.Hero) *mc.Monitor {
	return &mc.Monitor{
		MID: monitorfile.MonitorIdMap("死亡"),
		Tid: 0,
		MonitorLicense: mc.MonitorLicense{
			Type:    monitorfile.MonitorIdMap("死亡"),
			Owner:   hero,
			Subject: monitorfile.OwnerMap("自己"),
		},
		ListenerList: nil,
		Froze:        false,
	}
}

// 放逐monitor
func ExileMonitor(hero *mc.Hero) *mc.Monitor {
	return &mc.Monitor{
		MID: monitorfile.MonitorIdMap("放逐"),
		Tid: 0,
		MonitorLicense: mc.MonitorLicense{
			Type:    monitorfile.MonitorIdMap("放逐"),
			Owner:   hero,
			Subject: monitorfile.OwnerMap("自己"),
		},
		ListenerList: nil,
		Froze:        false,
	}
}

// 沉默
func SilentMonitor(hero *mc.Hero) *mc.Monitor {
	return &mc.Monitor{
		MID: monitorfile.MonitorIdMap("沉默"),
		Tid: 0,
		MonitorLicense: mc.MonitorLicense{
			Type:    monitorfile.MonitorIdMap("沉默"),
			Owner:   hero,
			Subject: monitorfile.OwnerMap("自己"),
		},
		ListenerList: nil,
		Froze:        false,
	}
}

// 生成一个默认英雄的monitor 含带最基础的死亡以及放逐monitor
func GeneralHeroMonitor(hero *mc.Hero, mcc *mc.MonitorCenter) (m1, m2 *mc.Monitor) {
	mcc.HeroMap[hero.Id] = hero
	dead := DeadMonitor(hero)
	exile := ExileMonitor(hero)
	//brother relationship
	mcc.AddMonitorList([]*mc.Monitor{dead, exile})
	return dead, exile
}
