package NatureGolem

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
)

func NatureGolemHeroInit() *hero.Hero {
	NatureGolem := hero.Hero{
		Id:          monitorfile.HeroNameToint("自然魔像"),
		Health:      7,
		THealth:     7,
		Name:        "自然魔像",
		AttackPoint: 3,
		Price:       450,
		ActionPoint: 1,

		Speed:          7,
		PositiveSkills: []int{1, 2, 3, 4, 7},
	}
	return &NatureGolem
}

// 光环
func NatureGolemMonitorLicense(hero *hero.Hero) (result *monitors.Monitor) {
	return
}

func NatureGolemMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	//nature := NatureGolemHeroInit()
	//mcc.AddHeroInHeroMap(nature)
	//cl := NatureGolemMonitorLicense(cr)
	mcc.MonitorsActive([]*monitors.Monitor{})
}
