package crazy

import (
	mc "train/monitor"
	"train/monitor/general"
	"train/monitor/monitorfile"
)

func CrazyHeroInit() *mc.Hero {
	crazy := mc.Hero{
		Owner:       0,
		Id:          1,
		Health:      5,
		Name:        "狂战士",
		AttackPoint: 1,
		GameTempo:   nil,
		RoundTempo:  nil,
	}
	return &crazy
}

func CrazyMonitorLicense(hero *mc.Hero) []*mc.Monitor {
	//狂战士之血技能监听自己，队友，对方单位
	result := []*mc.Monitor{}
	selfMonitor := mc.Monitor{
		MID: monitorfile.MonitorIdMap("狂战士之血"),
		Tid: 0,
		MonitorLicense: mc.MonitorLicense{
			Type:    monitorfile.MonitorIdMap("掉血"),
			Owner:   hero,
			Subject: monitorfile.OwnerMap("自己"),
		},
		ListenerList: nil,
		Froze:        false,
		Bubble:       map[uint]int{},
	}
	// 反弹bubble
	//selfMonitor.Bubble[monitorfile.BubbleIdMap("反弹")] = 1
	result = append(result, &selfMonitor)
	teamMonitor := mc.Monitor{
		MID: monitorfile.MonitorIdMap("狂战士之血"),
		Tid: 0,
		MonitorLicense: mc.MonitorLicense{
			Type:    monitorfile.MonitorIdMap("掉血"),
			Owner:   hero,
			Subject: monitorfile.OwnerMap("己方单位不包含自己"),
		},
		ListenerList: nil,
		Froze:        false,
	}
	result = append(result, &teamMonitor)
	enemyMonitor := mc.Monitor{
		MID: monitorfile.MonitorIdMap("狂战士之血"),
		Tid: 0,
		MonitorLicense: mc.MonitorLicense{
			Type:    monitorfile.MonitorIdMap("掉血"),
			Owner:   hero,
			Subject: monitorfile.OwnerMap("敌方单位"),
		},
		ListenerList: nil,
		Froze:        false,
	}
	result = append(result, &enemyMonitor)
	return result
}

func CrazyMonitorInit(mcc *mc.MonitorCenter) {
	cr := CrazyHeroInit()
	mcc.AddHeroInHeroMap(cr)
	cl := CrazyMonitorLicense(cr)
	mcc.AddMonitorList(cl)
	general.GeneralHeroMonitor(cr, mcc)
}
