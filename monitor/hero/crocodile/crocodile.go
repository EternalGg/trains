package crocodile

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func CrocodileHeroInit() *hero.Hero {
	Crocodile := hero.Hero{
		Id:             monitorfile.HeroNameToint("鳄鱼"),
		Health:         0,
		THealth:        0,
		Name:           "鳄鱼",
		AttackPoint:    0,
		Price:          350,
		ActionPoint:    1,
		GameTempo:      map[int]int{},
		RoundTempo:     map[int]int{},
		Speed:          12,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &Crocodile
}

func CrocodileMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func CrocodileMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	//cr := CrocodileHeroInit()
	//mcc.AddHeroInHeroMap(cr)
	//cl := CrocodileMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
