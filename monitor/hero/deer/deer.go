package deer

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func DeerHeroInit() *hero.Hero {
	Deer := hero.Hero{
		Id:          monitorfile.HeroNameToint("鹿"),
		Health:      3,
		THealth:     3,
		Name:        "鹿",
		AttackPoint: 1,
		ActionPoint: 1,
		Price:       150,
		GameTempo:   map[int]int{},
		RoundTempo:  map[int]int{},
		Speed:       8,
	}
	return &Deer
}

func DeerMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func DeerMonitorInit(mcc *mc.MonitorCenter) {
	deer := DeerHeroInit()
	mcc.AddHeroInHeroMap(deer)
	//cl := DeerMonitorLicense(cr)
	deer.PositiveSkills = []int{1, 2, 3}
	mcc.MonitorsActive([]*monitors.Monitor{})
}
