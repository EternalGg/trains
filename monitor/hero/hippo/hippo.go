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
		GameTempo:   map[int]int{},
		RoundTempo:  map[int]int{},
		Speed:       12,
	}
	return &Hippo
}

func HippoMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func HippoMonitorInit(mcc *mc.MonitorCenter) {
	cr := HippoHeroInit()
	mcc.AddHeroInHeroMap(cr)
	//cl := HippoMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
