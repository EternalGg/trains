package crazy

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

// 狂战士init
func CrazyHeroInit() *hero.Hero {
	crazy := hero.Hero{
		Id:          monitorfile.HeroNameToint("狂战士"),
		Health:      7,
		THealth:     7,
		Name:        "狂战士",
		AttackPoint: 3,
		Price:       600,
		ActionPoint: 1,
		GameTempo:   map[int]int{},
		RoundTempo:  map[int]int{},
		Speed:       10,
	}
	return &crazy
}

func CrazyMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	//狂战士之血技能监听自己，队友，对方单位
	result = &monitors.Monitor{
		Name:          "狂战士之血",
		Owner:         hero,
		MID:           monitorfile.MonitorIdMap("狂战士之血"),
		Tid:           0,
		Froze:         false,
		Logs:          []string{},
		IsForever:     true,
		LifeTimeState: 0,
		LifeTime:      0,
		PublishState:  map[int]int{},
		ListenLicense: []monitors.MonitorLicense{},
	}
	// forever state
	result.PublishState[monitorfile.BubbleIdMap("攻击力永久加成")] = 1
	// listenlicense
	selfBlood := monitors.MonitorLicense{
		ListenType: monitorfile.MonitorIdMap("掉血"),
		Subject:    monitorfile.OwnerMap("自己"),
	}
	enmyBlood := monitors.MonitorLicense{
		ListenType: monitorfile.MonitorIdMap("掉血"),
		Subject:    monitorfile.OwnerMap("敌方单位"),
	}
	teamMateBlood := monitors.MonitorLicense{
		ListenType: monitorfile.MonitorIdMap("掉血"),
		Subject:    monitorfile.OwnerMap("己方单位不包含自己"),
	}
	result.ListenLicense = append(result.ListenLicense, selfBlood, enmyBlood, teamMateBlood)
	return
}

func CrazyMonitorInit(mcc *mc.MonitorCenter) {
	cr := CrazyHeroInit()
	mcc.AddHeroInHeroMap(cr)
	cl := CrazyMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{cl})
}
