package ele

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
	"train/monitor/skills"
)

func EleHeroInit() *hero.Hero {
	Ele := hero.Hero{
		Id:             monitorfile.HeroNameToint("大象"),
		Health:         12,
		THealth:        12,
		Name:           "大象",
		AttackPoint:    3,
		Price:          450,
		ActionPoint:    1,
		AreadyChose:    false,
		Speed:          12,
		PositiveSkills: []int{1, 5, 3, 4},
	}
	return &Ele
}

// 巨大食草动物
func GH(hero *hero.Hero) (result *monitors.Monitor) {
	result = skills.GiantHerbivore()
	result.Owner = hero

	return
}

// 坚硬外皮
func HS(hero *hero.Hero) (result *monitors.Monitor) {
	result = skills.HardSkin()
	result.Owner = hero
	return
}

func EleMonitorInit(mcc *mc.MonitorCenter, hero *hero.Hero) {
	GH := GH(hero)
	HS := HS(hero)
	mcc.AddMonitorList([]*monitors.Monitor{GH, HS})
}
