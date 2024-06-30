package flamegolem

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func FlameGolemHeroInit() *hero.Hero {
	FlameGolem := hero.Hero{
		Id:             monitorfile.HeroNameToint("烈焰魔像"),
		Health:         7,
		THealth:        7,
		Name:           "烈焰魔像",
		AttackPoint:    3,
		Price:          450,
		ActionPoint:    1,
		GameTempo:      map[int]int{},
		RoundTempo:     map[int]int{},
		Speed:          10,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &FlameGolem
}

func FlameGolemMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func FlameGolemMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	//flame := FlameGolemHeroInit()
	//mcc.AddHeroInHeroMap(flame)
	//cl := FlameGolemMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
