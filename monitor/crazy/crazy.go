package crazy

import (
	mc "train/monitor"
	"train/monitor/general"
	"train/monitor/monitorfile"
)

func CrazyHeroInit() mc.Hero {
	crazy := mc.Hero{
		Id:          1,
		Health:      2,
		Name:        "狂暴战士",
		AttackPoint: 0,
	}

	return crazy
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
	}
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

func CrazyMonitorInit(hero *mc.Hero, mcc *mc.MonitorCenter) {
	mcc.AddHeroInHeroMap(hero)
	cl := CrazyMonitorLicense(hero)
	mcc.AddMonitorList(cl)
	general.GeneralHeroMonitor(hero, mcc)
}
