package octopus

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func OctopusHeroInit() *hero.Hero {
	Octopus := hero.Hero{
		Id:          monitorfile.HeroNameToint("章鱼"),
		Health:      12,
		THealth:     12,
		Name:        "章鱼",
		AttackPoint: 5,
		Price:       600,
		ActionPoint: 1,

		Speed:          12,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &Octopus
}

func OctopusMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func OctopusMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	//octopus := OctopusHeroInit()
	//mcc.AddHeroInHeroMap(octopus)
	//cl := OctopusMonitorLicense(cr)

	mcc.MonitorsActive([]*monitors.Monitor{})
}
