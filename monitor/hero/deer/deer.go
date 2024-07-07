package deer

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func DeerHeroInit() *hero.Hero {
	Deer := hero.Hero{
		Id:             monitorfile.HeroNameToint("鹿"),
		Health:         3,
		THealth:        3,
		Name:           "鹿",
		AttackPoint:    1,
		ActionPoint:    1,
		Price:          150,
		Speed:          8,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &Deer
}

func DeerMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func DeerMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	//deer := DeerHeroInit()
	//mcc.AddHeroInHeroMap(deer)
	//cl := DeerMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
