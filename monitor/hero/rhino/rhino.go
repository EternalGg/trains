package rhino

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func RhinoHeroInit() *hero.Hero {
	Rhino := hero.Hero{
		Id:          monitorfile.HeroNameToint("犀牛"),
		Health:      9,
		THealth:     9,
		Name:        "犀牛",
		AttackPoint: 3,
		ActionPoint: 1,
		Price:       350,
		GameTempo:   map[int]int{},
		RoundTempo:  map[int]int{},
		Speed:       12,
	}
	return &Rhino
}

//func RhinoMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
//	return
//}

func RhinoMonitorInit(mcc *mc.MonitorCenter) {
	cr := RhinoHeroInit()
	mcc.AddHeroInHeroMap(cr)
	//cl := RhinoMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
