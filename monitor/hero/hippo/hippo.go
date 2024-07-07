package hippo

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func HippoHeroInit() *hero.Hero {
	Hippo := hero.Hero{
		Id:          monitorfile.HeroNameToint("河马"),
		Health:      9,
		THealth:     9,
		Name:        "河马",
		AttackPoint: 3,
		Price:       350,
		ActionPoint: 1,

		Speed:          12,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &Hippo
}

func HippoMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func HippoMonitorInit(mcc *mc.MonitorCenter, hero *hero.Hero) {
	//hippo := HippoHeroInit()
	//mcc.AddHeroInHeroMap(hippo)
	//cl := HippoMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
