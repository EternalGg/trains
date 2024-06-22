package jugg

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func JuggHeroInit() *hero.Hero {
	Jugg := hero.Hero{
		Id:          monitorfile.HeroNameToint("主宰"),
		Health:      16,
		THealth:     16,
		Name:        "主宰",
		AttackPoint: 6,
		Price:       800,
		ActionPoint: 1,
		GameTempo:   map[int]int{},
		RoundTempo:  map[int]int{},
		Speed:       12,
	}
	return &Jugg
}

func JuggMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func JuggMonitorInit(mcc *mc.MonitorCenter) {
	cr := JuggHeroInit()
	mcc.AddHeroInHeroMap(cr)
	//cl := JuggMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
