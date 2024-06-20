package ele

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func EleHeroInit() *hero.Hero {
	Ele := hero.Hero{
		Id:          monitorfile.HeroNameToint("大象"),
		Health:      12,
		THealth:     12,
		Name:        "大象",
		AttackPoint: 4,
		Price:       450,
		ActionPoint: 1,
		GameTempo:   map[int]int{},
		RoundTempo:  map[int]int{},
		Speed:       12,
	}
	return &Ele
}

func EleMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func EleMonitorInit(mcc *mc.MonitorCenter) {
	cr := EleHeroInit()
	mcc.AddHeroInHeroMap(cr)
	cl := EleMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{cl})
}
