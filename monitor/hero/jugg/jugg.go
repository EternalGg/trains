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
		AttackPoint: 8,
		Price:       800,
		ActionPoint: 1,

		Speed:          12,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &Jugg
}

func JuggMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func JuggMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	//jugg := JuggHeroInit()
	//mcc.AddHeroInHeroMap(jugg)
	//cl := JuggMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
