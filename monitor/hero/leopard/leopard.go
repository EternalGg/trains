package leopard

import (
	mc "train/monitor"
	"train/monitor/hero"
	"train/monitor/monitorfile"
	"train/monitor/monitors"
	"train/monitor/skills"
)

func LeopardHeroInit() *hero.Hero {
	leopard := hero.Hero{
		Id:             monitorfile.HeroNameToint("豹子"),
		Health:         6,
		THealth:        6,
		Name:           "豹子",
		AttackPoint:    2,
		Price:          250,
		ActionPoint:    2,
		AreadyChose:    false,
		Speed:          8,
		PositiveSkills: []int{1, 2, 3, 4},
	}
	return &leopard
}

func Skill1leopard(hero *hero.Hero) (result *monitors.Monitor) {
	beast := skills.Beast()
	beast.Owner = hero
	return beast
}
func Skill2leopard(hero *hero.Hero) (result *monitors.Monitor) {
	night := skills.NocturnalAnimal()
	night.Owner = hero
	return night
}

func LeopardMonitorInit(mcc *mc.MonitorCenter, h *hero.Hero) {
	//leopard := LeopardHeroInit()
	//mcc.AddHeroInHeroMap(leopard)
	//skill1
	//野兽buff
	//skill2
	//夜行buff
	skill1 := Skill1leopard(h)
	skill2 := Skill2leopard(h)
	mcc.MonitorsActive([]*monitors.Monitor{skill1, skill2})
}
