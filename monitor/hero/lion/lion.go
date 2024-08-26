package lion

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
	"train/monitor/skills"
)

func LionHeroInit() *hero.Hero {
	Lion := hero.Hero{
		Id:          monitorfile.HeroNameToint("狮子"),
		Health:      5,
		THealth:     5,
		Name:        "狮子",
		AttackPoint: 2,
		Price:       450,
		ActionPoint: 10,

		Speed:          6,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &Lion
}

func Beast(h *hero.Hero) *monitors.Monitor {
	B := skills.Beast()
	B.Owner = h
	return B
}
func BeastKing(h *hero.Hero) *monitors.Monitor {
	B := skills.BeastKing()
	B.Owner = h
	return B
}
func NocturnalAnimal(h *hero.Hero) *monitors.Monitor {
	NA := skills.NocturnalAnimal()
	NA.Owner = h
	return NA
}
func LionMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	Beast := Beast(h)
	NA := NocturnalAnimal(h)
	BeastKing := BeastKing(h)
	mcc.MonitorsActive([]*monitors.Monitor{Beast, NA, BeastKing})
}
