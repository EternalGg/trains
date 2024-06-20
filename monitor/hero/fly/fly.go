package fly

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func FlyHeroInit() *hero.Hero {
	Fly := hero.Hero{
		Id:          monitorfile.HeroNameToint("苍蝇"),
		Health:      2,
		THealth:     2,
		Name:        "苍蝇",
		AttackPoint: 1,
		Price:       250,
		ActionPoint: 1,
		GameTempo:   map[int]int{},
		RoundTempo:  map[int]int{},
		Speed:       8,
	}
	return &Fly
}

func FlyMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func FlyMonitorInit(mcc *mc.MonitorCenter) {
	cr := FlyHeroInit()
	mcc.AddHeroInHeroMap(cr)
	cl := FlyMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{cl})
}
