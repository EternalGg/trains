package rhino

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
	"train/monitor/skills"
)

func RhinoHeroInit() *hero.Hero {
	Rhino := hero.Hero{
		Id:             monitorfile.HeroNameToint("犀牛"),
		Health:         10,
		THealth:        10,
		Name:           "犀牛",
		AttackPoint:    2,
		ActionPoint:    1,
		Price:          350,
		Speed:          12,
		PositiveSkills: []int{1, 2, 3, 4, 6},
	}
	return &Rhino
}

func GH(hero *hero.Hero) (result *monitors.Monitor) {
	result = skills.GiantHerbivore()
	result.Owner = hero
	return
}
func HS(hero *hero.Hero) (result *monitors.Monitor) {
	result = skills.HardSkin()
	result.Owner = hero
	return
}

func RhinoMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	GH := GH(h)
	HS := HS(h)
	mcc.MonitorsActive([]*monitors.Monitor{GH, HS})
}
