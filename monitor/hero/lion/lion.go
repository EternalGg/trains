package lion

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func LionHeroInit() *hero.Hero {
	Lion := hero.Hero{
		Id:          monitorfile.HeroNameToint("狮子"),
		Health:      5,
		THealth:     5,
		Name:        "狮子",
		AttackPoint: 2,
		Price:       250,
		ActionPoint: 1,

		GameTempo:  map[int]int{},
		RoundTempo: map[int]int{},
		Speed:      8,
	}
	return &Lion
}

func LionMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func LionMonitorInit(mcc *mc.MonitorCenter) {
	cr := LionHeroInit()
	mcc.AddHeroInHeroMap(cr)
	cl := LionMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{cl})
}
