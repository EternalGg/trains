package soldiers

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func SoldiersHeroInit() *hero.Hero {
	Soldiers := hero.Hero{
		Id:          monitorfile.HeroNameToint("士兵"),
		Health:      4,
		THealth:     4,
		Name:        "士兵",
		AttackPoint: 2,
		Price:       150,
		ActionPoint: 1,
		GameTempo:   map[int]int{},
		RoundTempo:  map[int]int{},
		Speed:       12,
	}
	return &Soldiers
}

func SoldiersMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func SoldiersMonitorInit(mcc *mc.MonitorCenter) {
	soldiers := SoldiersHeroInit()
	mcc.AddHeroInHeroMap(soldiers)
	//cl := SoldiersMonitorLicense(cr)
	soldiers.PositiveSkills = []int{1, 2, 3}
	mcc.MonitorsActive([]*monitors.Monitor{})
}
