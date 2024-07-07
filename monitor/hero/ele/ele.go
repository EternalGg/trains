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
		Id:          monitorfile.HeroNameToint("大象"),
		Health:      12,
		THealth:     12,
		Name:        "大象",
		AttackPoint: 3,
		Price:       450,
		ActionPoint: 1,

		Speed:          12,
		PositiveSkills: []int{1, 5, 3, 4},
	}
	return &Ele
}

func GH(hero *hero.Hero) (result *monitors.Monitor) {
	result = skills.GiantHerbivore()
	result.Owner = hero
	selfBlood := monitors.MonitorLicense{
		ListenType: monitorfile.MonitorIdMap("掉血"),
		Subject:    monitorfile.OwnerMap("自己"),
	}
	result.ListenLicense = append(result.ListenLicense, selfBlood)
	return
}

func EleMonitorInit(mcc *mc.MonitorCenter, hero *hero.Hero) {
	GH := GH(hero)
	mcc.AddMonitor(GH)
}
